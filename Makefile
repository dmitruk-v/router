test:
	go test -v -run=^Test github.com/dmitruk-v/router
bench:
	go test -v -bench . -benchmem -run=^Benchmark github.com/dmitruk-v/router
