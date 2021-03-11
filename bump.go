package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/logrusorgru/aurora"
	log "github.com/sirupsen/logrus"
)

const (
	filename = "VERSION"
	format   = "%d.%d.%d%s"
)

// bump is a little helper to update a semantic VERSION file
func main() {
	var err error
	var f *os.File
	var current, hash string
	var major, minor, patch int
	var sections = map[string]*int{"major": &major, "M": &major, "minor": &minor, "m": &minor, "patch": &patch, "p": &patch}

	var force = flag.Bool("f", false, "force create "+filename)
	var section = flag.String("b", "patch", "which section to bump: major or M, minor or m, patch or p")
	var commit = flag.Bool("c", false, "git commit hash")
	var manual = flag.String("s", "", "set manual (overrides -b -c): M[.m[.p[-hash]]]")
	flag.Parse()

	bumpme, ok := sections[*section]
	if !ok {
		exitif(fmt.Errorf("invalid bump type"), *section)
	}

	if *force {
		f, err = os.Create(filename)
		exitif(err, "failed to create VERSION")
	} else {
		f, err = os.OpenFile(filename, os.O_RDWR, 0644)
		exitif(err, "consider using: bump -f")
	}

	defer func() {
		exitif(f.Close(), "closing VERSION")
	}()

	if *manual != "" {
		// read semantic from manual argument
		fmt.Sscanf(*manual, format, &major, &minor, &patch, &hash)
	} else {
		// read semantic from file
		fmt.Fscanf(f, format, &major, &minor, &patch, &hash)
		current = fmt.Sprintf(
			format,
			aurora.Red(major),
			aurora.BrightGreen(minor),
			aurora.Cyan(patch),
			aurora.Yellow(hash),
		)
		*bumpme++

		if *commit {
			// append commit hash
			hash = "-" + getCommitHash("HEAD")
		} else {
			hash = "" // can be set previously
		}
	}

	_, err = f.Seek(0, io.SeekStart)
	exitif(err, "io seek")

	_, err = fmt.Fprintf(f, format, major, minor, patch, hash)
	exitif(err, "io write")

	fmt.Printf(
		"version %s bumped to "+format+"\n",
		current,
		aurora.Red(major),
		aurora.BrightGreen(minor),
		aurora.Cyan(patch),
		aurora.Yellow(hash),
	)
}

func getCommitHash(path string) string {
	f, err := os.OpenFile(".git/"+path, os.O_RDONLY, 0644)
	exitif(err, "failed to extract commit hash")

	defer func() {
		exitif(f.Close(), fmt.Sprintf("closing %s", path))
	}()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	var row = scanner.Text()
	if row[0:5] == "ref: " {
		// ref: refs/heads/master
		return getCommitHash(row[5:]) // refs/heads/master
	}

	return row[:7] // 4fa39df or short of revision hash (4fa39dfe2e8be5838fc4251f6aada4caa59ea2bf)
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
