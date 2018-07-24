package fasthttp_test

import (
	"fmt"
	"net"
	"testing"

	"github.com/buaazp/fasthttprouter"
	"github.com/parnurzeal/gorequest"
	"github.com/valyala/fasthttp"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

func StartServerOnPort(t *testing.T, method, path string, port int, handler func(ctx *fasthttp.RequestCtx, router *fasthttprouter.Router), requestBody interface{}) (gorequest.Response, string, []error) {
	ln, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	defer ln.Close()

	if err != nil {
		t.Fatalf("cannot start tcp server on port %d: %s", port, err)
	}

	router.Handle(method, path, handler)
	go fasthttp.Serve(ln, router.Handler)

	agent := gorequest.New()
	fullPath := fmt.Sprintf("http://localhost:%d%s", port, path)
	var (
		resp gorequest.Response
		body string
		errs []error
	)

	switch method {
	case GET:
		resp, body, errs = agent.Get(fullPath).End()
	case POST:
		resp, body, errs = agent.Post(fullPath).Send(requestBody).End()
	case PUT:
		resp, body, errs = agent.Put(fullPath).Send(requestBody).End()
	case DELETE:
		resp, body, errs = agent.Delete(fullPath).End()
	}

	return resp, body, errs
}
