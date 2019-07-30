package rpc

import "net/rpc"

const (
	Add = "Add"
	Mul = "Mul"
	Div = "Div"
)

type Args struct {
	A, B int
}

type CalculateInterface = interface {
	Add(args Args, replay *int) error
	Mul(args Args, replay *int) error
	Div(args Args, replay *int) error
}

func RegisterService(srvs []string, calculater CalculateInterface) {

	for _, srv := range srvs {
		_ = rpc.RegisterName(srv, calculater)
	}
}
