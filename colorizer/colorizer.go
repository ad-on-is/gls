package colorizer

import (
	"strings"

	"github.com/ad-on-is/gls/fileinfos"
	"github.com/ad-on-is/gls/iconizer"
	"github.com/ad-on-is/gls/settings"
	"github.com/jwalton/gchalk"
)

func Parse(value string, col string) string {
	c := strings.Split(col, " ")
	return gchalk.StyleMust(c[:]...)(value)
}

func Permissions(value string, permissions settings.PermissionColors) string {
	c := strings.Split(permissions.R, " ")
	value = strings.ReplaceAll(value, "r", gchalk.StyleMust(c[:]...)("r"))
	c = strings.Split(permissions.W, " ")
	value = strings.ReplaceAll(value, "w", gchalk.StyleMust(c[:]...)("w"))
	c = strings.Split(permissions.X, " ")
	value = strings.ReplaceAll(value, "x", gchalk.StyleMust(c[:]...)("x"))
	c = strings.Split(permissions.None, " ")
	value = strings.ReplaceAll(value, "-", gchalk.StyleMust(c[:]...)("-"))
	c = strings.Split(permissions.Prefix, " ")
	value = strings.ReplaceAll(value, "d", gchalk.StyleMust(c[:]...)("d"))
	value = strings.ReplaceAll(value, "L", gchalk.StyleMust(c[:]...)("L"))
	return value
}

func Name(item fileinfos.Item, iColor [3]uint8) string {
	// value := gchalk.RGB(iColor[0], iColor[1], iColor[2])(item.Name)
	value := gchalk.White(item.Name)
	// if item.IsLink {
	// 	value += gchalk.Gray(" => ") + gchalk.Gray(item.Link)
	// }
	return value
}

func Link(value string) string {
	return gchalk.White(value)
}

func Icon(icn *iconizer.Icon_Info) string {
	return gchalk.RGB(icn.GetColor()[0], icn.GetColor()[1], icn.GetColor()[2])(icn.GetGlyph())
}

func GetIconColor(icn iconizer.Icon_Info) [3]uint8 {
	return icn.GetColor()
}
