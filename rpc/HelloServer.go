package rpc

type HelloService struct{}

func (p *HelloService) Hello(request string, replay *string) error {
	*replay = "hello" + request
	return nil
}
