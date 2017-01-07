# fasthttp-test
for people who want to do integration test valyala/fasthttp
### get
```
go get github.com/tspn/fasthttp-test

```
### example
```go
func Handler(ctx *fasthttp.RequestCtx){
	ctx.WriteString("example")
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
```
