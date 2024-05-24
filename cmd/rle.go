package main

import (
	"fmt"
	tcmpr "github.com/xnacly/tcmpr/v1"
	"os"
	"time"
)

func main() {
	files, err := os.ReadDir("./data")
	if err != nil {
		panic(err)
	}
	start := time.Now()
	os.MkdirAll("./rle_out", 0777)
	for _, f := range files {
		in, err := os.Open("./data/" + f.Name())
		if err != nil {
			panic(err)
		}
		out, err := os.Create("./rle_out/" + f.Name())
		if err != nil {
			panic(err)
		}
		err = tcmpr.Compress(in, out)
		if err != nil {
			panic(err)
		}
	}
	since := time.Since(start)
	fmt.Println("Encoded", since, "avg time per file", since.Milliseconds()/int64(len(files)), "ms", "file count", len(files))
	fmt.Println("Decoded", time.Since(start))
}
