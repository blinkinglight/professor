# professor - golang secure pprof
Sometimes you need to expose things to internet, so some basic protection to debug/pprof would be useful.

```go
import "runtime" 

func main() {
	professor.SetToken("tokenstring") // default: securitytoken
	// or basic auth
	// professor.SetBasicAuth("user", "pass")
	professor.Launch(":1234")

	runtime.Goexit()
}
```

```
curl http://localhost:1234/debug/pprof/goroutine?debug=1&token=tokenstring
```
