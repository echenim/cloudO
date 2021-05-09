package log

import (
	"bufio"
	"encoding/binary"
	"os"
)

var (
	enc = binary.BigEndian
)

const (
	lenWidth = 8
)

func newStore(f *os.File) (*store, error) {
	fi, err := os.Stat(f.Name())
	if err != nil {
		return nil, err
	}
	size := uint64(fi.Size())
	return &store{
		File: f,
		size: size,
		buf:  bufio.NewWriter(f),
	}, nil
}
