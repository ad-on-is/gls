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
	"golang.org/x/crypto/ssh/terminal"
)

// test

func Long(items *[]fileinfos.Item, config *settings.Config, buf *bytes.Buffer) {
	t := table.NewWriter()
	t.SetStyle(table.Style{Options: table.Options{DrawBorder: false, SeparateColumns: true}, Box: table.BoxStyle{PaddingLeft: " ",
		PaddingRight: " "}})
	t.SetOutputMirror(os.Stdout)
	width, _, _ := terminal.GetSize(0)
	t.SetAllowedRowLength(width)

	// w := new(tabwriter.Writer)
	// w.Init(os.Stdout, 0, 0, 2, ' ', 0)

	theme := config.Themes[config.Theme]

	// fmt.Fprintln(w)
	for _, item := range *sorter.Sort(items, config) {
		//
		// out := ""

		if config.ShowOctal {
			// rows = append(rows, octal(&item, &theme))
			// out += octal(&item, &theme) + "\t"
		}
		// rows = append(rows, permissions(&item, &theme))
		// out += permissions(&item, &theme) + "\t"
		// out += size(&item, &theme) + "\t"
		// out += user(&item, &theme) + "\t"

		if config.ShowGroup {
			// out += group(&item, &theme) + "\t"
		}

		// out += date(&item, &theme) + "\t"

		if config.ShowGit {
			// out += git(&item, &theme) + "\t"
		}

		// out += name(&item, &theme, true)

		// fmt.Fprintln(w, out)
		icon, name, link, excl := name(&item, &theme, true)
		t.AppendRow(table.Row{octal(&item, &theme), permissions(&item, &theme), size(&item, &theme), user(&item, &theme), group(&item, &theme), date(&item, &theme), icon + " " + name + excl + link})
	}
	t.Render()
	// w.Flush()
}

func Short(items *[]fileinfos.Item, config *settings.Config) {

	textcol.Output = os.Stdout

	fmt.Println()

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
		icon, name, _, _ := name(&item, theme, false)

		if len(*item.Children) > 0 {
			n := (*tree).AddBranch(icon + name + " " + git(&item, theme))
			addTreeNodes(&n, item.Children, config, theme)
		} else {
			(*tree).AddNode(icon + " " + name + " " + git(&item, theme))
		}
	}
}
