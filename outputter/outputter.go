package outputter

import (
	"fmt"
	"os"
	"strings"

	"github.com/acarl005/textcol"
	"github.com/ad-on-is/gls/colorizer"
	"github.com/ad-on-is/gls/fileinfos"
	"github.com/ad-on-is/gls/settings"
	"github.com/ad-on-is/gls/sorter"
	"github.com/juju/ansiterm/tabwriter"
)

func Long(items *[]fileinfos.Item, config *settings.Config) {

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w)
	for _, item := range sorter.Sort(items, config) {
		if !config.ShowAll && strings.HasPrefix(item.Name, ".") {
			continue
		}
		iColor := colorizer.GetIconColor(*item.Icon())
		theme := config.Themes[config.Theme]
		fmt.Fprintln(w,
			colorizer.Parse(item.OctalPermissions(), theme.Colors.OctalPermissions)+"\t"+colorizer.Permissions(item.HumanPermissions(), theme.Colors.Permissions)+"\t"+colorizer.Parse(item.HumanSize(), theme.Colors.Size)+"\t"+colorizer.Parse(item.User, theme.Colors.User)+"\t"+colorizer.Parse(item.Group, theme.Colors.Group)+"\t"+colorizer.Parse(item.HumanDate("2006-01-02 15:04:05"), theme.Colors.Date)+"\t"+colorizer.Icon(item.Icon())+"  "+colorizer.Name(item, iColor))
	}
	w.Flush()
}

func Short(items *[]fileinfos.Item, config *settings.Config) {

	textcol.Output = os.Stdout

	colStrings := []string{}

	for _, item := range sorter.Sort(items, config) {
		if !config.ShowAll && strings.HasPrefix(item.Name, ".") {
			continue
		}
		iColor := colorizer.GetIconColor(*item.Icon())
		colStrings = append(colStrings, colorizer.Icon(item.Icon())+" "+colorizer.Name(item, iColor))

	}
	textcol.PrintColumns(&colStrings, 2)
	// w.Flush()
}
