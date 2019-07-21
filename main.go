package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	from, to      string
	offset, limit int64
)

func main() {
	flag.StringVar(&from, "from", "", "source file")
	flag.StringVar(&to, "to", "", "destination file")
	flag.Int64Var(&offset, "offset", 0, "offset in source file, bytes")
	flag.Int64Var(&limit, "limit", 0, "bytes to copy from source file to destination")
	flag.Parse()

	fmt.Println("Will copy", limit, "bytes from", from, "with offset", offset, "to", to)

	err := copyFile(from, to, offset, limit)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}
