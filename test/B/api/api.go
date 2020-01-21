package api

import (
	"github.com/OhYee/ait/server"
	"github.com/OhYee/ait/test/B/sum"
)

var register = server.MakeRegister("B")

type SumRequest struct {
	A int
	B int
}
type SumResponse struct {
	C int
}

var Sum = func() func(req SumRequest) (rep SumResponse, err error) {
	caller := register("sum", sum.Sum, SumRequest{}, SumResponse{})
	return func(req SumRequest) (rep SumResponse, err error) {
		err = caller(req, &rep)
		return
	}
}()
