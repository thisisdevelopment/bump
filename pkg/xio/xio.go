package xio

import (
	"bufio"
	"io"
	"os"

	"github.com/pkg/errors"
)

// GetFile opens or creates the file at path
func GetFile(path string, create bool) (f *os.File, err error) {

	if create {
		f, err = os.Create(path)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create %s", path)
		}
	}

	f, err = os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open %s", path)
	}

	return f, nil
}

// ReadFirstRow reads contents till first break newline is encountered
func ReadFirstRow(f *os.File) string {
	var scanner = bufio.NewScanner(f)
	scanner.Scan()
	return scanner.Text()
}

// ReplaceContent truncates the file and writes content at offset 0
func ReplaceContent(f *os.File, content string) (err error) {

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return errors.Wrap(err, "io seek")
	}

	err = f.Truncate(0)
	if err != nil {
		return errors.Wrap(err, "io truncate")
	}

	_, err = f.WriteString(content)
	if err != nil {
		return errors.Wrap(err, "io write")
	}

	err = f.Sync()
	if err != nil {
		return errors.Wrap(err, "io sync")
	}

	return nil
}
