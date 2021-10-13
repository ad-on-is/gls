package outputter

import (
	"fmt"
	"os"
	"strings"

	"github.com/acarl005/textcol"
	"github.com/ad-on-is/gls/fileinfos"
	"github.com/ad-on-is/gls/settings"
	"github.com/ad-on-is/gls/sorter"
	"github.com/juju/ansiterm/tabwriter"
)

// test

func Long(items *[]fileinfos.Item, config *settings.Config) {

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 0, 2, ' ', 0)

	theme := config.Themes[config.Theme]

	fmt.Fprintln(w)
	for _, item := range sorter.Sort(items, config) {
		if !config.ShowAll && strings.HasPrefix(item.Name, ".") {
			continue
		}

		out := ""

		if config.ShowOctal {
			out += octal(&item, &theme) + "\t"
		}
		out += permissions(&item, &theme) + "\t"
		out += size(&item, &theme) + "\t"
		out += user(&item, &theme) + "\t"

		if config.ShowGroup {
			out += group(&item, &theme) + "\t"
		}

		out += date(&item, &theme) + "\t"

		if config.ShowGit {
			out += git(&item, &theme) + "\t"
		}

		out += name(&item, &theme, true)

		fmt.Fprintln(w, out)
	}
	w.Flush()
}

func Short(items *[]fileinfos.Item, config *settings.Config) {

	textcol.Output = os.Stdout

	colStrings := []string{}
	theme := config.Themes[config.Theme]

	for _, item := range sorter.Sort(items, config) {
		if !config.ShowAll && strings.HasPrefix(item.Name, ".") {
			continue
		}
		colStrings = append(colStrings, name(&item, &theme, false)+" "+git(&item, &theme))

	}
	textcol.PrintColumns(&colStrings, 2)
	// w.Flush()
}
