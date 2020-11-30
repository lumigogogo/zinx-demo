package main

import (
	"github.com/lumigogogo/zinx/ziface"
	"github.com/lumigogogo/zinx/znet"
	// "github.com/aceld/zinx/ziface"
	// "github.com/aceld/zinx/znet"
)

type PingRouter struct {
	znet.Router
	// znet.BaseRouter
}

func (p *PingRouter) PreHandle(request ziface.IRequest) {

}

func (p *PingRouter) Handle(request ziface.IRequest) {
	// request.GetConnection().SendBuffMsg(1, []byte("ping...ping...ping"))
	request.GetConnection().SendMsg(1, []byte("ping...ping...ping"))
}

func (p *PingRouter) AfterHandle(request ziface.IRequest) {

}

type PangRouter struct {
	znet.Router
	// znet.BaseRouter
}

func (pr *PangRouter) PreHandle(request ziface.IRequest) {

}

func (pr *PangRouter) Handle(request ziface.IRequest) {
	// request.GetConnection().SendBuffMsg(1, []byte("pang...pang...pang"))
	request.GetConnection().SendMsg(1, []byte("pang...pang...pang"))
}

func (pr *PangRouter) AfterHandle(request ziface.IRequest) {

}

func main() {

	s := znet.NewServer("zinx-v0.1", "0.0.0.0", 9999)
	// s := znet.NewServer()
	s.AddRouter(1, &PingRouter{})
	s.AddRouter(2, &PangRouter{})

	s.Serve()
}
