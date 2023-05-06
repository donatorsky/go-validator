COUNT ?= "6"
BENCHTIME ?= "1s"
TIMEOUT ?= "20m"
FILE ?= "bench"
OLD ?= "old"
NEW ?= "new"

# Runs benchmark
bench:
	@go test -run=^# -bench=. -count="${COUNT}" -benchtime="${BENCHTIME}" -timeout="${BENCHTIME}" | tee "${FILE}.bench.log"

# Benchmarks statistics
bench-stats:
	@benchstat "${FILE}.bench.log"

# Compare benchmarks
bench-compare:
	@benchstat "${OLD}.bench.log" "${NEW}.bench.log"
