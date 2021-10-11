package sorter

import (
	"github.com/ad-on-is/gls/fileinfos"
	"github.com/ad-on-is/gls/settings"
)

func Sort(items *[]fileinfos.Item, config *settings.Config) []fileinfos.Item {

	if config.DirsFirst {
		return dirsFirst(items)
	}

	return *items

}

func dirsFirst(items *[]fileinfos.Item) []fileinfos.Item {
	sorted := []fileinfos.Item{}
	for _, item := range *items {
		if item.IsDir {
			sorted = append(sorted, item)
		}

	}
	for _, item := range *items {
		if !item.IsDir {
			sorted = append(sorted, item)
		}

	}
	return sorted
}
