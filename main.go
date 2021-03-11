package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/logrusorgru/aurora"
	log "github.com/sirupsen/logrus"
	"github.com/thisisdevelopment/bump/pkg/version"
)

const (
	filename = "VERSION"
)

// bump is a little helper to update a semantic VERSION file
func main() {
	var err error
	var f *os.File
	var vNew *version.Version
	var vCurrent, hash string
	var section = flag.String("b", "patch", "which section to bump: major, minor or patch")
	var force = flag.Bool("f", false, "force create ./"+filename)
	var commit = flag.Bool("c", false, "append git commit hash")
	flag.Parse()

	if *force {
		f, err = os.Create(filename)
		exitif(err, "failed to create %s", filename)
	} else {
		f, err = os.OpenFile(filename, os.O_RDWR, 0644)
		exitif(err, "consider using: bump -f")
	}

	defer func() {
		exitif(f.Close(), "failed to close %s", filename)
	}()

	vCurrent = read(f)
	var vType = version.Type(*section)
	vNew, err = version.Change(vType, vCurrent)
	exitif(err, "failed to change version")

	if *commit {
		// append commit hash
		hash = "-" + getCommitHash("HEAD")
	} else {
		hash = "" // clear if set previously
	}

	_, err = f.Seek(0, io.SeekStart)
	exitif(err, "io seek failed")

	_, err = fmt.Fprintf(f, "%s%s", vNew.String(), hash)
	exitif(err, "io write failed")

	fmt.Printf(
		"version %s bumped to %s%s\n",
		aurora.Cyan(vCurrent),
		aurora.BrightGreen(vNew.String()),
		aurora.Yellow(hash),
	)
}

func getCommitHash(path string) string {
	f, err := os.OpenFile(".git/"+path, os.O_RDONLY, 0644)
	exitif(err, "failed to extract commit hash")

	defer func() {
		exitif(f.Close(), "failed to close %s", path)
	}()

	var row = read(f)
	if row[0:5] == "ref: " {
		// ref: refs/heads/master
		return getCommitHash(row[5:]) // refs/heads/master
	}

	return row[:7] // 4fa39df or short of revision hash (4fa39dfe2e8be5838fc4251f6aada4caa59ea2bf)
}

func read(f *os.File) string {
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	return scanner.Text()
}

func exitif(err error, format string, ctx ...interface{}) {
	if err != nil {
		if len(ctx) > 0 {
			format = aurora.Sprintf(format, ctx...)
		}
		log.Error(aurora.Sprintf("%v %s", aurora.BrightRed(err), aurora.Yellow(format)))
		os.Exit(-1)
	}
}
