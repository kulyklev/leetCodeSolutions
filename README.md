# leetCodeSolutions

Commands for testing and profiling:
- `go test -run none -bench . -benchmem -memprofile p.out`
- `go test -run none -bench . -benchmem -cpuprofile p.out`
- `go tool pprof integerToRoman.test p.out`