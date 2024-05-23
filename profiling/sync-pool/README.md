# sync pool

```bash
$ GOWORK=off go test -v -bench=. -benchtime=2s -memprofile=prof.mem -cpuprofile=prof.cpu -gcflags -m=2
$ go tool pprof sync-pool.test prof.mem
$ go tool pprof sync-pool.test prof.cpu
$ GOWORK=off go test -run='^$' -bench=. -count=10 > sync-pool.txt
```
