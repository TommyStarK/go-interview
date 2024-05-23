# memory optimization

```bash
$ GOWORK=off go test -v -bench=. -benchtime=2s -memprofile=prof.mem -cpuprofile=prof.cpu -gcflags -m=2
$ go tool pprof memory-optimization.test prof.mem
$ go tool pprof memory-optimization.test prof.cpu
$ GOWORK=off go test -run='^$' -bench=. -count=10 > opti-mem.txt
```
