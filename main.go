package main

import (
	"os"

	"github.com/ad-on-is/gls/fileinfos"
	"github.com/ad-on-is/gls/outputter"
	"github.com/ad-on-is/gls/settings"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	path      = kingpin.Arg("path", "The path or file to show").Default(".").String()
	long      = kingpin.Flag("long", "Show long output").Short('l').Bool()
	all       = kingpin.Flag("all", "Show hidden").Short('a').Bool()
	dirsFirst = kingpin.Flag("dirs-first", "Show directories first").Short('d').Bool()
)

func main() {

	kingpin.Parse()

	config := settings.GetConfig()
	config.ShowAll = *all
	config.DirsFirst = *dirsFirst

	items := fileinfos.GetItems(*path)
	if *long {
		outputter.Long(&items, &config)
		os.Exit(0)
	}
	outputter.Short(&items, &config)

}
