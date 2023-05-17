package toybox

import "sync"

type ToyBox struct {
	cfgPath   string
	cfgSource string //local || http

	readyComponents []string
	components      []Component
	servers         []Server
}

type Option func(tb *ToyBox)

func New(opts ...Option) *ToyBox {
	tb := &ToyBox{
		cfgPath: "config/",
	}
	for _, opt := range opts {
		opt(tb)
	}
	return tb
}
func (tb *ToyBox) Run() {
	var wg sync.WaitGroup
	for _, component := range tb.components {
		wg.Add(1)
		go func(component Component) {
			defer wg.Done()
			component.Init()
		}(component)
	}
	wg.Wait()

}

// 获取本地配置
func WithCfgPath(path string) Option {
	return func(tb *ToyBox) {
		tb.cfgSource = "local"
		tb.cfgPath = path
	}
}

// 通过GET获取远程接口的配置
func WithCfgURL(url string) Option {
	return func(tb *ToyBox) {
		tb.cfgSource = "url"
		tb.cfgPath = url
	}
}

func EnableComponents(comps ...Component) Option {
	return func(tb *ToyBox) {
		tb.components = comps
	}
}

// func EnableServers(svcs ...Server) Option {
// 	return func(tb *ToyBox) {
// 		tb.servers = svcs
// 	}
// }
