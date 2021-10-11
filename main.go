package main

import (
	"os"

	"github.com/ad-on-is/gls/fileinfos"
	"github.com/ad-on-is/gls/outputter"
	"github.com/ad-on-is/gls/plugins"
	"github.com/ad-on-is/gls/settings"
	"github.com/jwalton/gchalk"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	path      = kingpin.Arg("path", "The path or file to show").Default(".").String()
	long      = kingpin.Flag("long", "Show long output").Short('l').Bool()
	all       = kingpin.Flag("all", "Show hidden").Short('a').Bool()
	group     = kingpin.Flag("group", "Show group next to user").Short('g').Bool()
	git       = kingpin.Flag("git", "Show Git status").Short('G').Bool()
	octal     = kingpin.Flag("octal", "Show octal permissions, ie 0755").Short('o').Bool()
	dirsFirst = kingpin.Flag("dirs-first", "Show directories first").Short('d').Bool()
)

func main() {
	gchalk.SetLevel(gchalk.LevelAnsi16m)
	kingpin.Parse()

	config := settings.GetConfig()
	config.ShowAll = *all
	config.ShowGit = *git
	config.ShowGroup = *group
	config.ShowOctal = *octal
	config.DirsFirst = *dirsFirst

	// test

	items := fileinfos.GetItems(*path)

	if config.ShowGit {
		plugins.ApplyGitStatus(&items)
	}

	if *long {
		outputter.Long(&items, &config)
		os.Exit(0)
	}
	outputter.Short(&items, &config)

}
