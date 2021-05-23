package log

//index defines our index file, which comprises a persisted file and a memory- mapped file.
//The size tells us the size of the index and where to write the next entry appended to the index.

import (
	"io"
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

//Close function makes sure the memory-mapped file has synced its data to the persisted
// file and that the persisted file has flushed its contents to stable storage.
//Then it truncates the persisted file to the amount of data that’s actually in it and closes the file.
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

//ead(int64) takes in an offset and returns the associated record’s position in the store.
//The given offset is relative to the segment’s base offset; 0 is always the offset of the
//index’s first entry, 1 is the second entry, and so on.
func (i *index) Read(in int64) (out uint32, pos uint64, err error) {
	if i.size == 0 {
		return 0, 0, io.EOF
	}

	if in == -1 {
		out = uint32((i.size / entWidth) - uint64(1))
	} else {
		out = uint32(in)
	}

	pos = uint64(out) * entWidth
	if i.size < pos+entWidth {
		return 0, 0, io.EOF
	}
	out = enc.Uint32(i.nmap[pos : pos+offWidth])
	pos = enc.Uint64(i.nmap[pos+offWidth : pos+entWidth])
	return out, pos, nil
}
