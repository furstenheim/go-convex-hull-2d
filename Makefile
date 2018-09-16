bench:
	go test -bench=.* -test.benchtime=3s
bench-graph:
	mkdir -p benchmarks/$$(git rev-parse HEAD)
	go test -run=XXX -bench=.* -cpuprofile cpu.prof
	go tool pprof -svg cpu.prof > benchmarks/$$(git rev-parse HEAD)/cpu.svg
	rm cpu.prof
