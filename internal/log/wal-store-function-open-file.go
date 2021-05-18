package log

import (
	"os"
)

func openFile(name string) (file *os.File, size int64, err error) {
	f, err := os.OpenFile(
		name, os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0644,
	)
	if err != nil {
		return nil, 0, err
	}

	fi, err := f.Stat()
	if err != nil {
		return nil, 0, err
	}

	return f, fi.Size(), nil
}
