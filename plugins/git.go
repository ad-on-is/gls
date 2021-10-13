package plugins

import (
	"os"
	"os/exec"
	"strings"

	"github.com/ad-on-is/gls/fileinfos"
)

func ApplyGitStatus(items *[]fileinfos.Item, path *string) {
	pathInfo, err := os.Lstat(*path)
	if err != nil {
		return
	}
	gitPath := *path
	if !pathInfo.IsDir() {
		ps := strings.Split(*path, "/")
		gitPath = strings.Join(ps[:len(ps)-1], "/")
	}
	c := exec.Command("git", "-C", gitPath, "status", "--porcelain", "-u")
	o, err := c.Output()
	if err != nil {
		return
	}
	s := string(o)
	statuses := strings.Split(s, "\n")
	traverse(items, &statuses, &gitPath)
}

func traverse(items *[]fileinfos.Item, statuses *[]string, gitPath *string) {
	if len(*items) == 0 {
		return
	}
	for i, item := range *items {
		if item.Excluded {
			continue
		}
		for _, status := range *statuses {
			if status == "" {
				continue
			}
			// needs some rework
			s, f := splitStatusAndFile(&status)
			if strings.Contains(f, clean(item.Root+item.Name, gitPath)) {
				(*items)[i].GitStatus = strings.ReplaceAll(s, "??", "U")
			}
			traverse((*items)[i].Children, statuses, gitPath)
		}
	}
}

func clean(s string, gp *string) string {
	so := strings.ReplaceAll(s, "./", "")
	if *gp != "." {
		so = strings.ReplaceAll(so, *gp, "")
	}

	return so
}

func splitStatusAndFile(status *string) (string, string) {
	split := strings.Split(strings.TrimLeft(*status, " "), " ")
	return split[0], split[1]
}
