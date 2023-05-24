package core

import (
	"chat/apps/imchat"
	"chat/apps/imchat/message"
	"chat/toybox/components/logger"
	"fmt"
	"sync"
)

type outputUnit struct {
	src    <-chan message.MessageBox
	Output []imchat.MessageOutPut
}

type processorUnit struct {
	src     <-chan message.MessageBox
	dst     chan<- message.MessageBox
	Process imchat.MessageProcess
}

type inputUnit struct {
	dst   chan<- message.MessageBox
	input <-chan message.MessageBox
}

func (cs *ChatHubServer) runPipline() {
	next, out := cs.startOutput()
	next, psut := cs.startProcessor(next)
	in := cs.startInput(next)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		cs.runOutput(out)
	}()

	wg.Add(1)
	go func() {
		wg.Done()
		cs.runProcessor(psut)
	}()

	wg.Add(1)
	go func() {
		wg.Done()
		cs.runInput(in)
	}()

	wg.Wait()
	fmt.Println("msg process pipline close.")
}

func (cs *ChatHubServer) startOutput() (chan<- message.MessageBox, *outputUnit) {
	src := make(chan message.MessageBox, 100)
	out := outputUnit{
		src: src,
	}
	for _, output := range cs.msgoutput {
		out.Output = append(out.Output, output)
		fmt.Printf("msg pipline load msg out %v\n", output.Name())
	}
	return src, &out
}

func (cs *ChatHubServer) startProcessor(dst chan<- message.MessageBox) (chan<- message.MessageBox, []*processorUnit) {
	duts := make([]*processorUnit, 0, len(cs.process))
	var src chan message.MessageBox
	for _, item := range cs.process {
		src = make(chan message.MessageBox, 100)
		duts = append(duts, &processorUnit{
			src:     src,
			dst:     dst,
			Process: item,
		})
		dst = src
	}
	return src, duts
}

func (cs *ChatHubServer) startInput(dst chan<- message.MessageBox) *inputUnit {
	return &inputUnit{
		dst:   dst,
		input: cs.broadcastInput,
	}
}

func (cs *ChatHubServer) runOutput(outUnit *outputUnit) {
	for msg := range outUnit.src {
		select {
		case <-cs.ctx.Done():
			return
		default:
			for _, out := range outUnit.Output {
				out.Write(msg)
			}
		}
	}
}

func (cs *ChatHubServer) runProcessor(proUnit []*processorUnit) {
	var wg sync.WaitGroup
	for _, pt := range proUnit {
		wg.Add(1)
		go func(pt *processorUnit) {
			for {
				select {
				case <-cs.ctx.Done():
					wg.Done()
					return
				case val := <-pt.src:
					res, err := pt.Process.Process(val)
					if err != nil {
						logger.Log.Errorf("process Exceptions: %v", err)
						continue
					}
					pt.dst <- res
				}
			}
		}(pt)
	}
	wg.Wait()
}

func (cs *ChatHubServer) runInput(inUnit *inputUnit) {
	for val := range inUnit.input {
		inUnit.dst <- val
	}
}
