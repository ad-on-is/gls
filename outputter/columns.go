package outputter

import (
	"github.com/ad-on-is/gls/colorizer"
	"github.com/ad-on-is/gls/fileinfos"
	"github.com/ad-on-is/gls/settings"
)

func octal(item *fileinfos.Item, theme *settings.Theme) string {
	return colorizer.Parse(item.OctalPermissions(), theme.Colors.OctalPermissions)
}

func permissions(item *fileinfos.Item, theme *settings.Theme) string {
	return colorizer.Permissions(item.HumanPermissions(), theme.Colors.Permissions)
}

func size(item *fileinfos.Item, theme *settings.Theme) string {
	return colorizer.Parse(item.HumanSize(), theme.Colors.Size)
}

func user(item *fileinfos.Item, theme *settings.Theme) string {
	return colorizer.Parse(item.User, theme.Colors.User)
}

func group(item *fileinfos.Item, theme *settings.Theme) string {
	return colorizer.Parse(item.Group, theme.Colors.Group)
}

func date(item *fileinfos.Item, theme *settings.Theme) string {
	return colorizer.Parse(item.HumanDate(theme.DateFormat), theme.Colors.Date)
}

func git(item *fileinfos.Item, theme *settings.Theme) string {
	return colorizer.Parse(gitPrefix(item.GitStatus, theme), gitColor(item.GitStatus, theme))
}

func name(item *fileinfos.Item, theme *settings.Theme, showSymlink bool) string {
	special, icn := item.Icon()
	specialColor := colorizer.GetIconColor(*icn)

	icnOut := ""
	nameOut := ""
	linkOut := ""

	if item.IsDir {
		if theme.ColorizeGitIcon && item.GitStatus != "" {
			icnOut = colorizer.Parse(icn.GetGlyph(), gitColor(item.GitStatus, theme))
		} else {
			if theme.SpecialColorizeDirIcons && special != "" {
				icnOut = colorizer.RGB(icn.GetGlyph(), specialColor)
			} else {
				icnOut = colorizer.Parse(icn.GetGlyph(), theme.Colors.DirIcon)
			}
		}

		if theme.ColorizeGitName && item.GitStatus != "" {
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

		if theme.ColorizeGitIcon && item.GitStatus != "" {
			icnOut = colorizer.Parse(icn.GetGlyph(), gitColor(item.GitStatus, theme))
		} else {
			if theme.SpecialColorizeFileIcons && special != "" {
				icnOut = colorizer.RGB(icn.GetGlyph(), specialColor)
			} else {
				icnOut = colorizer.Parse(icn.GetGlyph(), theme.Colors.FileIcon)
			}
		}

		if theme.ColorizeGitName && item.GitStatus != "" {
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
