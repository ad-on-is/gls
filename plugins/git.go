package plugins

import (
	"os/exec"
	"strings"

	"github.com/ad-on-is/gls/fileinfos"
)

func ApplyGitStatus(items *[]fileinfos.Item) {
	m := getByStatus("m")
	d := getByStatus("d")
	o := getByStatus("o")

	statuses := append(append(m, d...), o...)

	i := 0
	for _, item := range *items {

		for _, status := range statuses {
			if strings.HasPrefix(status, item.Name) || strings.HasPrefix(status, item.Root+item.Name) {
				item.GitStatus = parseStatus(&status)
				(*items)[i].GitStatus = parseStatus(&status)
			}
		}
		i++
	}
}

func parseStatus(output *string) string {
	ss := strings.Split(*output, ":")
	return ss[len(ss)-1]
}

func getByStatus(status string) []string {
	c := exec.Command("git", "ls-files", "-"+status)
	o, _ := c.Output()
	s := string(o)
	s = strings.ReplaceAll(s, "\n", ":"+strings.ToUpper(strings.ReplaceAll(status, "o", "u"))+"\n")
	lines := strings.Split(s, "\n")
	return lines
}
