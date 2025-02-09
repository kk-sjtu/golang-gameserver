package define

type HandlerParam struct {
	HandlerKey string
	Data       interface{}
}

type Handler func(data interface{})

type Player struct {
	UId            uint64
	FriendList     []uint64 // 朋友
	HandlerParamCh chan HandlerParam
	handlers       map[string]Handler
}

func (p *Player) HandlerRegister() {
	// 注册处理器
}

func (p *Player) Run() {
	for {
		select {
		case handlerParam := <-p.HandlerParamCh:
			if fn, ok := p.handlers[handlerParam.HandlerKey]; ok {
				fn(handlerParam.Data) // 传下来什么命令，就去找
			}
		}
	}
}
