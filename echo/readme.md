# Echo Service
This service reponses everything that client passes to server.

## Service Setup
### API
Define `echo.api` first as follows:
```
type (
	Request {
		Name string `path:"name,options=[you,me]"` // parameters are auto validated
	}

	Response {
		Message string `json:"message"`
	}
)

service echo-api {
	@handler EchoHandler
	get /echo/:name(Request) returns (Response)
}
```

### Code Generation
Generate the service code using [goctl](https://github.com/zeromicro/go-zero/tree/master/tools/goctl)
```
goctl api go -api echo.api -dir .
```
The generated project will be like:
```
.
├── echo.api
├── echo.go
├── etc
│   └── echo-api.yaml
└── internal
    ├── config
    │   └── config.go
    ├── handler
    │   ├── echohandler.go
    │   └── routes.go
    ├── logic
    │   └── echologic.go
    ├── svc
    │   └── servicecontext.go
    └── types
        └── types.go
```

### Biz Logic Implementation
Implement the biz logic in `internal/logic/echologic.go`, the real place where develops need to focus on.
```golang
func (l *EchoLogic) Echo(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info(req)
	ret := &types.Response{
		Message: req.Name,
	}
	return ret, nil
}
```

### Run
Run the service in the current dir:
```sh
go run echo.go -f etc/echo-api.yaml
```

Make the request use `curl`:
```sh
curl -i http://localhost:8888/echo/me
curl -i http://localhost:8888/echo/you

# validation failed case
curl -i http://localhost:8888/echo/foo
```

## Reference
- https://github.com/zeromicro/go-zero
