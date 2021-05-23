package log

//index defines our index file, which comprises a persisted file and a memory- mapped file.
//The size tells us the size of the index and where to write the next entry appended to the index.

import (
	"os"

	"github.com/tysontate/gommap"
)

var (
	offWidth uint64 = 4
	posWidth uint64 = 8
	//entWidth to jump straight to the position of an entry
	//given its offset since the position in the file is offset * entWidth
	entWidth = offWidth + posWidth
)

//newIndex(*os.File) creates an index for the given file. We create the index and save the current
//size of the file so we can track the amount of data in the index file as we add index entries.
//We grow the file to the max index size before memory-mapping the file and then return the created index
// to the caller.
func newIndex(f *os.File, c Config) (*index, error) {
	idx := &index{file: f}

	fi, err := os.Stat(f.Name())
	if err != nil {
		return nil, err
	}

	idx.size = uint64(fi.Size())
	if err = os.Truncate(
		f.Name(), int64(c.Segment.MaxIndexByte),
	); err != nil {
		return nil, err
	}

	if idx.nmap, err = gommap.Map(
		idx.file.Fd(),
		gommap.PROT_READ|gommap.PROT_WRITE,
		gommap.MAP_SHARED,
	); err != nil {
		return nil, err
	}

	return idx, nil
}

func (i *index) Close() error {
	if err := i.nmap.Sync(gommap.MS_ASYNC); err != nil {
		return err
	}
	if err := i.file.Sync(); err != nil {
		return err
	}

	if err := i.file.Truncate(int64(i.size)); err != nil {
		return err
	}

	return i.file.Close()
}
