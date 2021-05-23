package log

import (
	"os"

	"github.com/tysontate/gommap"
)

type index struct {
	file *os.File
	nmap gommap.MMap
	size uint64
}
