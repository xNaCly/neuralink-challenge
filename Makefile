zstd-encoding:
	go build ./cmd/zstd.go
	./zstd
	du -h ./zstd_out
