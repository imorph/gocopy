package main

import (
	"io"
	"os"
	"strings"

	"github.com/cheggaaa/pb/v3"
	"github.com/pkg/errors"
)

// copyFile will try copy
func copyFile(from string, to string, offset int64, limit int64) error {
	fromFile, err := os.Open(from)
	if err != nil {
		return errors.Wrapf(err, "open file %s failed", from)
	}
	defer fromFile.Close()

	// get file info to determine file size
	ffi, err := fromFile.Stat()
	if err != nil {
		return errors.Wrapf(err, "can't get stats for file file %s", from)
	}
	if ffi.Size() < limit+offset {
		return errors.Errorf("Can not copy, %vb + %vb is bigger than %s size (which is %vb)", limit, offset, from, ffi.Size())
	}

	toFile, err := os.Create(to)
	if err != nil {
		return errors.Wrapf(err, "open file %s failed", to)
	}
	defer toFile.Close()

	var pBarLimit int64
	switch ffSize := ffi.Size(); {
	case ffSize > offset+limit:
		pBarLimit = limit
	case ffSize >= offset:
		pBarLimit = ffSize - offset
	default:
		pBarLimit = 0
	}
	// special case /dev files like /dev/zero /dev/urandom actual infinite size but zero size from stats
	if strings.HasPrefix(from, `/dev/`) {
		pBarLimit = limit
		// special case actual zero file size
		if from == "/dev/null" {
			pBarLimit = 0
		}
	}

	ffsReader := io.NewSectionReader(fromFile, offset, limit)
	pBar := pb.Full.Start64(pBarLimit)
	pBarReader := pBar.NewProxyReader(ffsReader)
	_, err = io.Copy(toFile, pBarReader)
	if err != nil {
		return errors.Wrapf(err, "can't copy from %s to %s", from, to)
	}
	pBar.Finish()

	if err := fromFile.Close(); err != nil {
		return errors.Wrapf(err, "close of file %s failed", from)
	}
	if err := toFile.Close(); err != nil {
		return errors.Wrapf(err, "close of file %s failed", to)
	}

	return nil
}
