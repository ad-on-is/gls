package plugins

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/ad-on-is/gls/fileinfos"
)

func ApplyGitStatus(items *[]fileinfos.Item) {

	c := exec.Command("git", "status", "--porcelain", "-u")
	o, _ := c.Output()
	s := string(o)
	statuses := strings.Split(s, "\n")
	i := 0
	for _, item := range *items {

		for _, status := range statuses {
			if status == "" {
				continue
			}
			s, f := parseStatus(&status)
			fmt.Println(f)
			fmt.Println(clean(item.Root + item.Name))
			if strings.HasPrefix(f, clean(item.Name)) || strings.HasPrefix(f, clean(item.Root+item.Name)) {
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
