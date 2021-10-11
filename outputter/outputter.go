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

// test

func Long(items *[]fileinfos.Item, config *settings.Config) {

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintln(w)
	for _, item := range sorter.Sort(items, config) {
		if !config.ShowAll && strings.HasPrefix(item.Name, ".") {
			continue
		}

		theme := config.Themes[config.Theme]
		out := ""

		if config.ShowOctal {
			out += colorizer.Parse(item.OctalPermissions(), theme.Colors.OctalPermissions) + "\t"
		}
		out += colorizer.Permissions(item.HumanPermissions(), theme.Colors.Permissions) + "\t" + colorizer.Parse(item.HumanSize(), theme.Colors.Size) + "\t" + colorizer.Parse(item.User, theme.Colors.User) + "\t"
		if config.ShowGroup {
			out += colorizer.Parse(item.Group, theme.Colors.Group) + "\t"
		}

		out += colorizer.Parse(item.HumanDate(theme.DateFormat), theme.Colors.Date) + "\t"

		if config.ShowGit {
			out += colorizer.Parse(gitPrefix(item.GitStatus, &theme), gitColor(item.GitStatus, &theme))
			out += "\t"

		}

		out += itemName(&item, &theme, true)

		fmt.Fprintln(w,
			out)
	}
	w.Flush()
}

func gitColor(status string, theme *settings.Theme) string {
	switch status {
	case "M":
		return theme.Colors.Git.M
	case "D":
		return theme.Colors.Git.D
	case "U":
		return theme.Colors.Git.U
	}
	return "white"
}

func gitPrefix(status string, theme *settings.Theme) string {
	switch status {
	case "M":
		return theme.GitPrefix.M
	case "D":
		return theme.GitPrefix.D
	case "U":
		return theme.GitPrefix.U
	}
	return ""
}

func itemName(item *fileinfos.Item, theme *settings.Theme, showSymlink bool) string {
	special, icn := item.Icon()
	specialColor := colorizer.GetIconColor(*icn)
	// specialColor := [3]uint8{255, 136, 0}
	// specialColor := color.RGB(sc[0], sc[1], sc[2])
	// dat := specialColor.Sprintf("%s", "testasf")
	// ret += dat

	icnOut := ""
	nameOut := ""
	linkOut := ""

	if item.IsDir {
		if theme.ColorizeGitIcon {
			icnOut = colorizer.Parse(icn.GetGlyph(), gitColor(item.GitStatus, theme))
		} else {
			if theme.SpecialColorizeDirIcons && special != "" {
				icnOut = colorizer.RGB(icn.GetGlyph(), specialColor)
			} else {
				icnOut = colorizer.Parse(icn.GetGlyph(), theme.Colors.DirIcon)
			}
		}

		if theme.ColorizeGitName {
			nameOut = colorizer.Parse(item.Name, gitColor(item.GitStatus, theme))
		} else {
			if theme.SpecialColorizeDirs && special != "" {
				nameOut = colorizer.RGB(item.Name, specialColor)
			} else {
				nameOut = colorizer.Parse(item.Name, theme.Colors.DirName)
			}
		}

		nameOut = colorizer.Parse(theme.FolderPrefix, theme.Colors.FolderIndicator) + nameOut + colorizer.Parse(theme.FolderSuffix, theme.Colors.FolderIndicator)
	}

	if !item.IsDir {

		if theme.ColorizeGitIcon {
			icnOut = colorizer.Parse(icn.GetGlyph(), gitColor(item.GitStatus, theme))
		} else {
			if theme.SpecialColorizeFileIcons && special != "" {
				icnOut = colorizer.RGB(icn.GetGlyph(), specialColor)
			} else {
				icnOut = colorizer.Parse(icn.GetGlyph(), theme.Colors.FileIcon)
			}
		}

		if theme.ColorizeGitName {
			nameOut = colorizer.Parse(item.Name, gitColor(item.GitStatus, theme))
		} else {
			if theme.SpecialColorizeFiles && special != "" {
				nameOut = colorizer.RGB(item.Name, specialColor)
			} else {
				nameOut = colorizer.Parse(item.Name, theme.Colors.FileName)
			}
		}

		if item.IsExecutable {
			nameOut = colorizer.Parse(theme.ExePrefix, theme.Colors.ExeIndicator) + nameOut + colorizer.Parse(theme.ExeSuffix, theme.Colors.ExeIndicator)
		}

	}

	if item.IsLink && showSymlink {
		linkOut = colorizer.Parse(" => ", theme.Colors.Symlink.Arrow)
		if theme.SpecialColorizeLinks && special != "" {
			linkOut += colorizer.RGB(item.Link, specialColor, theme.DimSpecialColorizeLinks)
		} else {
			linkOut += colorizer.Parse(item.Link, theme.Colors.Symlink.Link)
		}

	}
	return icnOut + "  " + nameOut + linkOut

}

func Short(items *[]fileinfos.Item, config *settings.Config) {

	textcol.Output = os.Stdout

	colStrings := []string{}
	theme := config.Themes[config.Theme]

	for _, item := range sorter.Sort(items, config) {
		if !config.ShowAll && strings.HasPrefix(item.Name, ".") {
			continue
		}
		colStrings = append(colStrings, itemName(&item, &theme, false)+" "+colorizer.Parse(gitPrefix(item.GitStatus, &theme), gitColor(item.GitStatus, &theme)))

	}
	textcol.PrintColumns(&colStrings, 2)
	// w.Flush()
}
