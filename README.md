This small project shows how terrible is OOP and OOP thinking for performance.
If you're in the high load business - use DOD.
Use structures of arrays instead of arrays of structures.

Run the code with `go test -bench=.`

Results for me look like this:

```
goos: linux
goarch: amd64
pkg: github.com/reflechant/cache-miss
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
BenchmarkCPAOOP-8                1326728               921.5 ns/op
BenchmarkCPAOOPPtr-8             1252496               990.1 ns/op
BenchmarkCPAOOPPtrRange-8        1244438               999.0 ns/op
BenchmarkCPAOOPIface-8           1000000              1349 ns/op
BenchmarkCPADOD-8               47681769                24.34 ns/op
BenchmarkCPADODTwoLoops-8       20639508                65.87 ns/op
PASS
ok      github.com/reflechant/cache-miss        12.613s
```