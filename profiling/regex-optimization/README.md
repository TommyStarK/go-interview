# regex optimization

```bash
$ GOWORK=off go test -v -bench=. -benchtime=2s -memprofile=prof.mem -cpuprofile=prof.cpu
$ GOWORK=off go test -run='^$' -bench=. -count=10 > opti-regex.txt
```
