package fasthttp_test_test

import (
	"github.com/valyala/fasthttp"
	"testing"
	"github.com/tspn/fasthttp-test"
)

func Handler(ctx *fasthttp.RequestCtx){
	ctx.WriteString("example")
}

func API(ctx *fasthttp.RequestCtx){
	email := ctx.FormValue("email")
	ctx.Write(email)
}

func TestSimpleGETRequest(t *testing.T){
	resp, body, errs := fasthttp_test.StartServerOnPort(t, fasthttp_test.GET, "/example", 9000, Handler, nil)

	if errs != nil {
		for _, err := range errs {
			t.Error(err)
		}
	}

	if resp == nil {
		t.Error("Response is nil")
	}

	if body != "example" {
		t.Error("expected example but got", body)
	}
}

func TestSimpleAPIRequest(t *testing.T){
	form := `email=example@mail.com`

	resp, body, errs := fasthttp_test.StartServerOnPort(t, fasthttp_test.POST, "/api", 9000, API, form)

	if errs != nil {
		for _, err := range errs {
			t.Error(err)
		}
	}

	if resp == nil {
		t.Error("Response is nil")
	}

	if body != "example@mail.com" {
		t.Error("expected example@mail.com but got", body)
	}
}