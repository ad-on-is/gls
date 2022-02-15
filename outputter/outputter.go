package outputter

import (
	"bytes"
	"fmt"
	"os"

	"github.com/ad-on-is/gls/fileinfos"
	"github.com/ad-on-is/gls/settings"
	"github.com/ad-on-is/gls/sorter"
	"github.com/ad-on-is/gls/textcol"
	"github.com/ad-on-is/gls/treeprint"
	"github.com/jedib0t/go-pretty/v6/table"
	"golang.org/x/term"
)

// test

func Long(items *[]fileinfos.Item, config *settings.Config, buf *bytes.Buffer) {
	t := table.NewWriter()
	t.SetStyle(table.Style{Options: table.Options{DrawBorder: false, SeparateColumns: true}, Box: table.BoxStyle{PaddingLeft: " ",
		PaddingRight: " "}})
	t.SetOutputMirror(os.Stdout)
	width, _, _ := term.GetSize(0)
	t.SetAllowedRowLength(width)

	theme := config.Themes[config.Theme]

	for _, item := range *sorter.Sort(items, config) {
		row := table.Row{}

		if config.ShowOctal {
			row = append(row, octal(&item, &theme))
		}

		row = append(row, permissions(&item, &theme))
		row = append(row, size(&item, &theme))
		row = append(row, user(&item, &theme))

		if config.ShowGroup {
			row = append(row, group(&item, &theme))
		}

		row = append(row, date(&item, &theme))
		if config.ShowGit {
			row = append(row, git(&item, &theme))
			// out += git(&item, &theme) + "\t"
		}

		icon, name, link, _ := name(&item, &theme, true)
		row = append(row, icon+" "+name+link)

		t.AppendRow(row)
	}
	t.Render()
}

func Short(items *[]fileinfos.Item, config *settings.Config) {

	textcol.Output = os.Stdout

	colStrings := []string{}
	theme := config.Themes[config.Theme]

	for _, item := range *sorter.Sort(items, config) {
		icon, name, _, _ := name(&item, &theme, false)
		colStrings = append(colStrings, icon+" "+name+" "+git(&item, &theme))

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
	for _, item := range *sorter.Sort(items, config) {
		icon, name, _, excl := name(&item, theme, false)

		if len(*item.Children) > 0 {
			n := (*tree).AddBranch(icon + name + " " + git(&item, theme) + excl)
			addTreeNodes(&n, item.Children, config, theme)
		} else {
			(*tree).AddNode(icon + " " + name + " " + git(&item, theme) + excl)
		}
	}
}
