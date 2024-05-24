package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/klauspost/compress/zstd"
)

func main() {
	files, err := os.ReadDir("./data")
	if err != nil {
		panic(err)
	}
	start := time.Now()
	os.MkdirAll("./zstd_out", 0777)
	for _, f := range files {
		in, err := os.Open("./data/" + f.Name())
		if err != nil {
			panic(err)
		}
		out, err := os.Create("./zstd_out/" + f.Name())
		if err != nil {
			panic(err)
		}
		compress, err := zstd.NewWriter(out, zstd.WithEncoderLevel(zstd.SpeedBestCompression))
		if err != nil {
			panic(err)
		}
		_, err = io.Copy(compress, in)
		if err != nil {
			panic(err)
		}
		compress.Close()
	}
	since := time.Since(start)
	fmt.Println("Encoded", since, "avg time per file", since.Milliseconds()/int64(len(files)), "ms", "file count", len(files))
	fmt.Println("Decoded", time.Since(start))
}
