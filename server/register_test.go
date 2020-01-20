package server

import (
	"fmt"
	"math/rand"
	"testing"
)

// var closeServerA = make(chan bool, 1)
// var closeServerB = make(chan bool, 1)

// func serverA() {

// }

var register = MakeRegister("test_server")

func sum(a int, b int) (c int) {
	c = a + b
	return
}

type SumRequest struct {
	A int
	B int
}
type SumReponse struct {
	C int
}

func Sum(req SumRequest) (rep SumReponse) {
	register("sum", sum, req, &rep)
	return
}

func div(a int, b int) (c int, d int) {
	c = a / b
	d = a % b
	return
}

type DivRequest struct {
	A int
	B int
}
type DivReponse struct {
	C int
	D int
}

func Div(req DivRequest) (rep DivReponse) {
	register("div", div, req, &rep)
	return
}

func TestRegister(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("Test Sum %d", i), func(t *testing.T) {
			a := rand.Int()
			b := rand.Int()
			var rep SumReponse
			if CallFunction(sum, SumRequest{a, b}, &rep); a+b != rep.C {
				t.Errorf("%d + %d = %d got %d", a, b, a+b, rep.C)
			}
		})
	}

	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("Test Div %d", i), func(t *testing.T) {
			a := rand.Int()
			b := rand.Int()
			var rep DivReponse
			if CallFunction(div, DivRequest{a, b}, &rep); a/b != rep.C || a%b != rep.D {
				t.Errorf("%d / %d = %d ... %d got %d %d", a, b, a/b, a%b, rep.C, rep.D)
			}
		})
	}
}

var serverClose = make(chan bool, 1)

func serverA() {
	
}
