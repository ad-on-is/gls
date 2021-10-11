package fileinfos

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/ad-on-is/gls/iconizer"
)

type Item struct {
	Permissions  os.FileMode
	User         string
	Group        string
	Size         int64
	DateModified int64
	DateCreated  int64
	Name         string
	IsDir        bool
	IsLink       bool
	Link         string
	Extension    string
}

func (i *Item) Icon() *iconizer.Icon_Info {
	if i.IsDir {
		icon, exists := iconizer.Icon_Set["dir-"+strings.ToLower(strings.ReplaceAll(i.Name, ".", ""))]
		if exists {
			return icon
		}
		return iconizer.Icon_Def["dir"]
	}
	// first check if icon for filename is present
	icon, exists := iconizer.Icon_FileName[strings.ToLower(i.Name)]
	if exists {
		return icon
	}
	// now check if icon for extension is present
	icon, exists = iconizer.Icon_Ext[i.Extension]
	if exists {
		return icon
	}
	// default icon
	return iconizer.Icon_Def["file"]
}

func (i *Item) HumanDate(layout string) string {
	t := time.Unix(i.DateModified, 0)
	return t.Format(layout)
}

func (i *Item) HumanPermissions() string {
	return i.Permissions.String()
}

func (i *Item) OctalPermissions() string {
	s := fmt.Sprintf("%04o", i.Permissions)
	return s[len(s)-4:]
}

func (i *Item) HumanSize() string {
	const unit = 1000
	if i.Size == 0 {
		return "-"
	}
	if i.Size < unit {
		return fmt.Sprintf("%dB", i.Size)
	}
	div, exp := int64(unit), 0
	for n := i.Size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%c",
		float64(i.Size)/float64(div), "KMGTPE"[exp])
}

func GetItems(path string) []Item {
	items := []Item{}
	pathInfo, _ := os.Lstat(path)

	files := []fs.FileInfo{}

	if pathInfo.IsDir() {
		if !strings.HasSuffix(path, "/") {
			path = path + "/"
		}
		files, _ = ioutil.ReadDir(path)

	} else {
		files = append(files, pathInfo)
	}

	for _, file := range files {
		item := Item{}
		stat := file.Sys().(*syscall.Stat_t)
		uname, gname := getUserGroupNames(stat)
		item.Name = file.Name()
		item.User = uname
		item.Group = gname
		item.Size = file.Size()
		item.Permissions = file.Mode()
		item.IsDir = file.IsDir()
		item.DateModified = stat.Mtim.Sec
		item.Extension = strings.ReplaceAll(filepath.Ext(path+file.Name()), ".", "")

		if file.Mode()&os.ModeSymlink != 0 {
			item.IsLink = true
			item.Link, _ = os.Readlink(path + file.Name())

		}

		items = append(items, item)

	}
	return items
}

func getUserGroupNames(stat *syscall.Stat_t) (uname string, gname string) {
	uid := stat.Uid
	gid := stat.Gid
	u := strconv.FormatUint(uint64(uid), 10)
	g := strconv.FormatUint(uint64(gid), 10)
	usr, _ := user.LookupId(u)
	group, _ := user.LookupGroupId(g)
	return usr.Username, group.Name
}
