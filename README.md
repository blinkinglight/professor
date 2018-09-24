# professor
Safer interface to Golang net/http/pprof

```go
	professor.SetToken("tokenstring") // default: securitytoken
	professor.Launch(":1234")
```

```
curl http://localhost:1234/debug/pprof/goroutine?debug=1&token=randomas
```