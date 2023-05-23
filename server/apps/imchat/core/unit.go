package core

import (
	"chat/apps/imchat"
	"chat/apps/imchat/message"
	"chat/apps/imchat/pipline/output"
	"chat/apps/imchat/pipline/process"
	"chat/toybox/components/config"
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
	input <-chan []byte
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
	for name, output := range output.Outputs {
		out.Output = append(out.Output, output())
		fmt.Printf("msg pipline load msg out %v\n", name)
	}
	return src, &out
}

func (cs *ChatHubServer) startProcessor(dst chan<- message.MessageBox) (chan<- message.MessageBox, []*processorUnit) {
	processorList := config.Cfg.GetStringSlice("chat_process")
	if len(processorList) <= 0 {
		panic("no process is configured")
	}
	duts := make([]*processorUnit, 0, len(processorList))
	var src chan message.MessageBox
	for _, item := range processorList {
		if val, ok := process.Process[item]; ok {
			src = make(chan message.MessageBox, 100)
			duts = append(duts, &processorUnit{
				src:     src,
				dst:     dst,
				Process: val(),
			})
			dst = src
		}
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
	for _, output := range outUnit.Output {
		go func(output imchat.MessageOutPut) {
			for {
				select {
				case <-cs.ctx.Done():
					return
				case val := <-outUnit.src:
					output.Write(val)
				}
			}
		}(output)
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
		msg, err := message.NewMsgBoxWithByte(val)
		if err != nil {
			logger.Log.Errorf("input read bytes err: %v", err)
			continue
		}
		inUnit.dst <- msg
	}
}
