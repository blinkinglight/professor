# professor - golang secure pprof
Safer interface to Golang net/http/pprof

```go
import "runtime" 

func main() {
	professor.SetToken("tokenstring") // default: securitytoken
	professor.Launch(":1234")

	runtime.Goexit()
}
```

```
curl http://localhost:1234/debug/pprof/goroutine?debug=1&token=tokenstring
```
