package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/ad-on-is/gls/colorizer"
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
	tree      = kingpin.Flag("tree", "Show tree view").Short('t').Bool()
	nest      = kingpin.Flag("nest", "Nest level for tree view").Default("2").Short('n').Int()
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

	level := 1

	if *tree {
		level = *nest
	}

	items, err := fileinfos.GetItems(*path, all, &config.ExcludeDirs, level)
	if err != nil {
		if config.DisplayInfos {
			fmt.Println(colorizer.Parse("  "+*path+" not found", config.Themes[config.Theme].Colors.Info))
		}
		os.Exit(0)
	}
	if len(*items) == 0 {
		if config.DisplayInfos {
			fmt.Println(colorizer.Parse("  Folder is empty", config.Themes[config.Theme].Colors.Info))
		}
		os.Exit(0)
	}

	if config.ShowGit {
		plugins.ApplyGitStatus(items, path)
	}

	if *tree {
		outputter.Tree(items, &config)
		os.Exit(0)
	}

	if *long {
		buf := bytes.NewBuffer([]byte(""))
		outputter.Long(items, &config, buf)
		os.Exit(0)
	}
	outputter.Short(items, &config)

}
