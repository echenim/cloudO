package log

import (
	"bufio"
	"os"
	"sync"
)

type store struct {
	*os.File
	mu   sync.Mutex
	buf  *bufio.Writer
	size uint64
}
