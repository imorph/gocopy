package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
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

	if from == "" {
		flag.PrintDefaults()
		log.Fatalln("Sorry, but you did not enter file to copy from. Exiting.")
	}

	if to == "" {
		flag.PrintDefaults()
		log.Fatalln("Sorry, but you did not enter file to copy to. Exiting.")
	}

	fmt.Println("Will try to copy", limit, "bytes from", from, "with offset", offset, "to", to)

	err := copyFile(from, to, offset, limit)
	if err != nil {
		switch err := errors.Cause(err).(type) {
		case *os.PathError:
			log.Println("File with path:", err.Path, "not found")
			os.Exit(1)
		default:
			log.Fatalln(err)
		}
	}
}
