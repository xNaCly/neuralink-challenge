all: zstd-encoding tcmpr-encoding

zstd-encoding:
	go build ./cmd/zstd.go
	./zstd
	du -h ./zstd_out

tcmpr-encoding:
	go build ./cmd/rle.go
	./rle
	du -h ./rle_out
