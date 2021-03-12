package xgit

import (
	"os"

	"github.com/thisisdevelopment/bump/pkg/xerr"
	"github.com/thisisdevelopment/bump/pkg/xio"
)

// GetCommitHash tries to extract the current commit hash pointed to by ./.git/HEAD
func GetCommitHash(path string) string {
	f, err := os.OpenFile(".git/"+path, os.O_RDONLY, 0644)
	xerr.Exitif(err, "failed to extract commit hash")

	defer func() {
		xerr.Exitif(f.Close(), "failed to close %s", path)
	}()

	var row = xio.ReadFirstRow(f)
	if row[0:5] == "ref: " {
		// ref: refs/heads/master
		return GetCommitHash(row[5:]) // follow to eg refs/heads/master
	}

	return row[:7] // 4fa39df or short of revision hash (4fa39dfe2e8be5838fc4251f6aada4caa59ea2bf)
}
