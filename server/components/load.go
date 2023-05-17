package components

// import (
// 	"chat/toybox"
// 	"fmt"
// 	"sync"
// )

// var readyComponents = make(map[string]toybox.Component)

// func LoadComponents(componentName ...string) error {
// 	if len(componentName) <= 0 {
// 		return fmt.Errorf("a dependent component must be entered %v", componentName)
// 	}
// 	wg := sync.WaitGroup{}
// 	for _, name := range componentName {
// 		wg.Add(1)
// 		go func(name string) {
// 			defer wg.Done()
// 			component, ok := Components[name]
// 			if !ok {
// 				fmt.Printf("unknown components %v\n", name)
// 			}
// 			cp := component()

// 			// 初始化组件
// 			err := cp.Init()
// 			if err != nil {
// 				fmt.Printf("initialization component %v error: %v\n", name, err)
// 				return
// 			}

// 			if _, ok := readyComponents[name]; !ok {
// 				readyComponents[name] = cp
// 			}
// 		}(name)
// 	}
// 	wg.Wait()
// 	return nil
// }
