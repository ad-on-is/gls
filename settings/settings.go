package settings

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Themes    map[string]Theme `json:"themes"`
	Theme     string           `json:"theme"`
	ShowAll   bool
	DirsFirst bool
}

type Theme struct {
	Colors Colors `json:"colors"`
}

type Colors struct {
	OctalPermissions string           `json:"octalPermissions"`
	Permissions      PermissionColors `json:"permissions"`
	Size             string           `json:"size"`
	User             string           `json:"user"`
	Group            string           `json:"group"`
	Date             string           `json:"date"`
}

type PermissionColors struct {
	None   string `json:"none"`
	Prefix string `json:"prefix"`
	R      string `json:"r"`
	W      string `json:"w"`
	X      string `json:"x"`
}

func GetConfig() Config {
	homedir, _ := os.UserHomeDir()
	settingsFile := homedir + "/.config/gls.json"
	jsonFile, _ := os.Open(settingsFile)
	// if err != nil {
	// 	createEmptySettings(settingsFile)
	// }

	jsonData, _ := ioutil.ReadAll(jsonFile)
	theme := make(map[string]Theme)

	theme["default"] = Theme{
		Colors: Colors{
			OctalPermissions: "gray",
			Permissions: PermissionColors{
				None:   "gray",
				Prefix: "gray",
				R:      "yellow",
				W:      "green",
				X:      "red",
			},
			Size:  "white",
			User:  "white",
			Group: "yellow",
			Date:  "gray",
		},
	}

	config := Config{
		ShowAll:   false,
		DirsFirst: false,
		Theme:     "default",
		Themes:    theme,
	}

	json.Unmarshal(jsonData, &config)
	return config
}

func Hex2Byte(s string) string {
	b, _ := hex.DecodeString(s)
	return fmt.Sprintf("%d", b)
}
