package outputter

import (
	"fmt"
	"os"

	"github.com/ad-on-is/gls/fileinfos"
	"github.com/ad-on-is/gls/settings"
	"github.com/ad-on-is/gls/sorter"
	"github.com/ad-on-is/gls/textcol"
	"github.com/ad-on-is/gls/treeprint"
	"github.com/juju/ansiterm/tabwriter"
)

// test

func Long(items *[]fileinfos.Item, config *settings.Config) {

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 0, 2, ' ', 0)

	theme := config.Themes[config.Theme]

	fmt.Fprintln(w)
	for _, item := range sorter.Sort(items, config) {

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

	fmt.Println()

	colStrings := []string{}
	theme := config.Themes[config.Theme]

	for _, item := range sorter.Sort(items, config) {
		colStrings = append(colStrings, name(&item, &theme, false)+" "+git(&item, &theme))

	}
	textcol.PrintColumns(&colStrings, 2)
}

func Tree(items *[]fileinfos.Item, config *settings.Config) {

	theme := config.Themes[config.Theme]
	tree := treeprint.New(theme.Colors.Tree)

	addTreeNodes(&tree, items, config, &theme)
	fmt.Println(tree.String())
	// w.Flush()
}

func addTreeNodes(tree *treeprint.Tree, items *[]fileinfos.Item, config *settings.Config, theme *settings.Theme) {
	for _, item := range sorter.Sort(items, config) {

		if len(*item.Children) > 0 {
			n := (*tree).AddBranch(name(&item, theme, true) + " " + git(&item, theme))
			addTreeNodes(&n, item.Children, config, theme)
		} else {
			(*tree).AddNode(name(&item, theme, true) + " " + git(&item, theme))
		}
	}
}
