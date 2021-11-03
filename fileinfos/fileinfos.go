package fileinfos

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/ad-on-is/gls/iconizer"
)

type Item struct {
	Permissions   os.FileMode
	Root          string
	User          string
	Group         string
	Size          int64
	DateModified  int64
	DateCreated   int64
	Name          string
	IsDir         bool
	IsLink        bool
	IsHidden      bool
	IsExecutable  bool
	Link          string
	LinkRoot      string
	LinkName      string
	Extension     string
	LinkExtension string
	GitStatus     string
	Children      *[]Item
	Excluded      bool
}

func (i *Item) Icon() (string, *iconizer.Icon_Info) {

	if i.IsDir {
		icon, exists := iconizer.Icon_Set["dir-"+strings.ToLower(strings.ReplaceAll(i.Name, ".", ""))]
		if exists {
			return "dir", icon
		}
		if i.IsLink {
			icon, exists = iconizer.Icon_Set["dir-"+strings.ToLower(strings.ReplaceAll(i.LinkName, ".", ""))]
			if exists {
				return "dir", icon
			}
			return "", iconizer.Icon_Def["hiddendir"]
		}

		return "", iconizer.Icon_Def["diropen"]
	} else {
		icon, exists := iconizer.Icon_FileName[strings.ToLower(i.Name)]
		if exists {
			return "file", icon
		}
		if i.IsLink {
			icon, exists = iconizer.Icon_FileName[strings.ToLower(i.LinkName)]
			if exists {
				return "file", icon
			}
		}
		icon, exists = iconizer.Icon_Ext[i.Extension]
		if exists {
			return "ext", icon
		}
		if i.IsLink {
			icon, exists = iconizer.Icon_Ext[i.LinkExtension]
			if exists {
				return "ext", icon
			}
			return "", iconizer.Icon_Def["hiddenfile"]
		}

		return "", iconizer.Icon_Def["file"]
	}

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
		return "  -  "
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

func GetItems(path string, all *bool, excludeDirs *[]string, maxLevel int) (*[]Item, error) {

	rootFile, err := os.Lstat(path)
	level := 1

	if err != nil {
		return nil, errors.New("path not found")
	}

	if !rootFile.IsDir() {
		items := []Item{}
		item := getItem(&rootFile, path)

		if (rootFile).Mode()&os.ModeSymlink != 0 {
			if strings.Contains(path, "/") {
				ss := strings.Split(path, "/")
				item.Root = strings.Join(ss[:len(ss)-1], "/") + "/"
			} else {
				wd, _ := os.Getwd()
				item.Root = wd + "/"
			}

			getLinkItem(item)
			if !item.IsDir {
				items = append(items, *item)
				return &items, nil
			}
			return traverse(item.Root+item.Link, all, excludeDirs, maxLevel, &level), nil
		}
		items = append(items, *item)
		return &items, nil
	}

	return traverse(path, all, excludeDirs, maxLevel, &level), nil
}

func traverse(path string, all *bool, excludeDirs *[]string, maxLevel int, level *int) *[]Item {

	cl := *level + 1
	items := []Item{}
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}
	f, err := os.Open(path)

	if err != nil {
		return &items
	}

	files, _ := f.Readdirnames(-1)

	for _, dn := range files {

		file, err := os.Lstat(path + "/" + dn)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if !*all && strings.HasPrefix(file.Name(), ".") {
			continue
		}

		item := getItem(&file, path)
		if (file).Mode()&os.ModeSymlink != 0 {
			getLinkItem(item)
		}
		sort.Strings(*excludeDirs)
		exclude := false
		i := sort.SearchStrings(*excludeDirs, file.Name())
		if i < len(*excludeDirs) && (*excludeDirs)[i] == file.Name() {
			exclude = true
		}
		if !exclude {
			if item.IsDir && cl <= maxLevel {
				tPath := item.Root + item.Name
				if item.IsLink {
					tPath = item.Root + item.Link
				}
				item.Children = traverse(tPath, all, excludeDirs, maxLevel, &cl)
			}
		} else {
			item.Excluded = true
		}
		items = append(items, *item)
	}

	return &items
}

func getItem(file *fs.FileInfo, path string) *Item {
	item := Item{}
	item.Root = path
	stat := (*file).Sys().(*syscall.Stat_t)
	uname, gname := getUserGroupNames(stat)
	item.Name = (*file).Name()
	item.IsHidden = isHidden(item.Name)
	item.User = uname
	item.Group = gname
	item.Size = (*file).Size()
	item.Permissions = (*file).Mode()
	item.IsDir = (*file).IsDir()
	item.Children = &[]Item{}
	item.DateModified = stat.Mtim.Sec
	item.Extension = strings.ReplaceAll(filepath.Ext(path+(*file).Name()), ".", "")
	item.IsExecutable = (*file).Mode()&0111 == 0111 && !(*file).IsDir()

	return &item
}

func getLinkItem(item *Item) {
	itemPath := item.Root + "/" + item.Name
	itemPath = strings.ReplaceAll(itemPath, "//", "/")
	item.IsLink = true
	lnk, err := os.Readlink(itemPath)
	if err == nil {
		item.Link = lnk
	} else {
		item.Link = ""
	}

	// points somewhere else
	if strings.HasPrefix(item.Link, "/") || strings.HasPrefix(item.Link, "..") {
		ss := strings.Split(item.Link, "/")
		item.Root = strings.Join(ss[:len(ss)-1], "/") + "/"
	}

	linkPath := item.Root + "/" + item.Link
	linkPath = strings.ReplaceAll(linkPath, "//", "/")

	linkinfo, err := os.Lstat(linkPath)
	if err == nil {
		item.IsDir = linkinfo.IsDir()
		ss := strings.Split(item.Link, "/")
		item.LinkName = ss[len(ss)-1]
		item.LinkExtension = strings.ReplaceAll(filepath.Ext(item.Root+item.LinkName), ".", "")
		item.IsHidden = isHidden(item.LinkName)
	}

}

func isHidden(s string) bool {
	return strings.HasPrefix(s, ".") && !strings.HasPrefix(s, "..") && !strings.HasPrefix(s, "./")
}

func getUserGroupNames(stat *syscall.Stat_t) (uname string, gname string) {
	uid := stat.Uid
	gid := stat.Gid
	u := strconv.FormatUint(uint64(uid), 10)
	g := strconv.FormatUint(uint64(gid), 10)

	uout := u
	gout := g
	usr, err := user.LookupId(u)
	if err == nil {
		uout = usr.Username
	}

	group, err := user.LookupGroupId(g)

	if err == nil {
		gout = group.Name
	}

	return uout, gout
}
