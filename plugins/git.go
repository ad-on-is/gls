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
	i := 0
	for _, item := range *items {

		for _, status := range statuses {
			if status == "" {
				continue
			}
			// needs some rework
			s, f := parseStatus(&status)
			if strings.Contains(f, clean(item.Name)) {
				(*items)[i].GitStatus = strings.ReplaceAll(s, "??", "U")
			}
		}
		i++
	}
}

func clean(s string) string {
	return strings.ReplaceAll(s, "./", "")
}

func parseStatus(status *string) (string, string) {
	split := strings.Split(strings.TrimLeft(*status, " "), " ")
	return split[0], split[1]
}
