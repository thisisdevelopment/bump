package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
	log "github.com/sirupsen/logrus"
	"github.com/thisisdevelopment/bump/pkg/version"
	"github.com/thisisdevelopment/bump/pkg/xerr"
	"github.com/thisisdevelopment/bump/pkg/xgit"
	"github.com/thisisdevelopment/bump/pkg/xio"
)

const (
	// path to the semantic version file
	path = "VERSION"
)

var (
	// GitCommit holds short commit hash of source tree
	GitCommit string
	// GitBranch holds current branch name the code is built off
	GitBranch string
	// GitState shows whether there are uncommitted changes
	GitState string
	// BuildDate holds RFC3339 formatted UTC date (build time)
	BuildDate string
	// Version holds contents of ./VERSION file, if exists, or the value passed via the -version option
	Version string
)

// bump is a little helper to update a semantic VERSION file
func main() {
	var err error
	var f *os.File
	var vNew *version.Version
	var vCurrent, hash string
	var section = flag.String("b", "patch", "which section to bump: major, minor or patch")
	var force = flag.Bool("f", false, "force create ./"+path)
	var commit = flag.Bool("c", false, "append git commit hash")
	var info = flag.Bool("i", false, "prints the current version")
	var verbose = flag.Bool("v", false, "print verbose")
	flag.Parse()

	if *verbose {
		buildInfo()
	}

	f, err = xio.GetFile(path, *force)
	xerr.Exitif(err, "xio.GetFile")

	defer func() {
		xerr.Exitif(f.Close(), "failed to close %s", path)
	}()

	vCurrent = xio.ReadFirstRow(f)
	if vCurrent == "" {
		vCurrent = "0.0.0"
	}

	if *info {
		if *verbose {
			log.Printf("semantic of %s is %s\n", path, aurora.Cyan(vCurrent))
		} else {
			fmt.Println(vCurrent)
		}
		os.Exit(0)
	}

	vNew, err = version.Change(version.Type(*section), vCurrent)
	xerr.Exitif(err, "failed to change version")

	if *commit {
		// append commit hash
		hash = "-" + xgit.GetCommitHash("HEAD")
	}

	err = xio.ReplaceContent(f, fmt.Sprintf("%s%s", vNew.String(), hash))
	xerr.Exitif(err, "failed to replace content")

	if *verbose {
		log.Printf(
			"version %s bumped to %s%s\n",
			aurora.Cyan(vCurrent),
			aurora.BrightGreen(vNew.String()),
			aurora.Yellow(hash),
		)
	} else {
		fmt.Println(vNew.String() + hash)
	}

}

func buildInfo() {
	if Version != "" {
		fmt.Printf(`%s(
	Version: %s
	Commit: %s
	Branch: %s
	Status: %s
	BuildDate: %s)`+"\n\n",
			aurora.Cyan("ThisIsSemanticBump"),
			aurora.Yellow(Version),
			aurora.Yellow(GitCommit),
			GitBranch,
			GitState,
			aurora.Yellow(BuildDate),
		)
	}
}
