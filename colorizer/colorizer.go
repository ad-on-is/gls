package colorizer

import (
	"strings"

	"github.com/ad-on-is/gls/iconizer"
	"github.com/ad-on-is/gls/settings"

	"github.com/jwalton/gchalk"
)

func Parse(value string, col string, dim ...bool) string {
	if col == "" {
		return gchalk.White(value)
	}
	c := strings.Split(col, " ")
	if len(dim) > 0 && dim[0] {
		return gchalk.WithDim().StyleMust(c[:]...)(value)
	}
	return gchalk.StyleMust(c[:]...)(value)
}

func RGB(value string, col [3]uint8, dim ...bool) string {
	if len(dim) > 0 && dim[0] {
		return gchalk.WithDim().RGB(col[0], col[1], col[2])(value)
	}
	return gchalk.RGB(col[0], col[1], col[2])(value)
}

func Permissions(value string, permissions settings.PermissionColors) string {

	r := Parse("r", permissions.R)
	w := Parse("w", permissions.W)
	x := Parse("x", permissions.X)
	n := Parse("-", permissions.R)
	d := Parse("d", permissions.R)
	l := Parse("L", permissions.R)

	value = strings.ReplaceAll(value, "r", r)
	value = strings.ReplaceAll(value, "w", w)
	value = strings.ReplaceAll(value, "x", x)
	value = strings.ReplaceAll(value, "d", d)
	value = strings.ReplaceAll(value, "-", n)
	value = strings.ReplaceAll(value, "L", l)
	return value

}

func Icon(icn *iconizer.Icon_Info) string {
	return gchalk.RGB(icn.GetColor()[0], icn.GetColor()[1], icn.GetColor()[2])(icn.GetGlyph())
}

func GetIconColor(icn iconizer.Icon_Info) [3]uint8 {
	return icn.GetColor()
}
