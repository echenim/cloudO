package log

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	write = []byte("hello world")
	width = uint64(len(write)) + lenWidth
)

func TestStoreAppendRead(t *testing.T) {
	f, err := ioutil.TempFile("", "store appedn_read_test")
	require.NoError(t, err)
	defer os.Remove(f.Name())
	s, err := newStore(f)
	require.NoError(t, err)
	testAppend(t, s)
	testRead(t, s)
	testReadAt(t, s)

	s, err = newStore(f)
	require.NoError(t, err)
	testRead(t, s)

}



func testAppend(t *testing.T, s *store) {

}

func testRead(t *testing.T, s *store) {

}

func testReadAt(t *testing.T, s *store) {

}
