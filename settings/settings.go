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
	ShowGroup bool
	ShowOctal bool
	DirsFirst bool
}

type Theme struct {
	DateFormat               string `json:"dateFormat"`
	FolderSuffix             string `json:"folderSuffix"`
	FolderPrefix             string `json:"folderPrefix"`
	ExeSuffix                string `json:"exeSuffix"`
	ExePrefix                string `json:"exePrefix"`
	SpecialColorizeLinks     bool   `json:"specialColorizeLinks"`
	DimSpecialColorizeLinks  bool   `json:"dimSpecialColorizeLinks"`
	SpecialColorizeDirs      bool   `json:"specialColorizeDirs"`
	SpecialColorizeFiles     bool   `json:"specialColorizeFiles"`
	SpecialColorizeDirIcons  bool   `json:"specialColorizeDirIcons"`
	SpecialColorizeFileIcons bool   `json:"specialColorizeFileIcons"`
	Colors                   Colors `json:"colors"`
}

type Colors struct {
	OctalPermissions string           `json:"octalPermissions"`
	Permissions      PermissionColors `json:"permissions"`
	Symlink          SymlinkColors    `json:"symlink"`
	Size             string           `json:"size"`
	User             string           `json:"user"`
	Group            string           `json:"group"`
	Date             string           `json:"date"`
	DirName          string           `json:"dirName"`
	FileName         string           `json:"fileName"`
	DirIcon          string           `json:"dirIcon"`
	FileIcon         string           `json:"fileIcon"`
	DirLinkIcon      string           `json:"dirLinkIcon"`
	DirLinkName      string           `json:"dirLinkName"`
	ExeIndicator     string           `json:"exeIndicator"`
	FolderIndicator  string           `json:"folderIndicator"`
}

type PermissionColors struct {
	None   string `json:"none"`
	Prefix string `json:"prefix"`
	R      string `json:"r"`
	W      string `json:"w"`
	X      string `json:"x"`
}

type SymlinkColors struct {
	Arrow string `json:"arrow"`
	Link  string `json:"link"`
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
		FolderSuffix:             "/",
		FolderPrefix:             "",
		ExeSuffix:                "*",
		ExePrefix:                "",
		SpecialColorizeLinks:     true,
		DimSpecialColorizeLinks:  true,
		SpecialColorizeDirs:      true,
		SpecialColorizeFiles:     true,
		SpecialColorizeDirIcons:  true,
		SpecialColorizeFileIcons: true,
		Colors: Colors{
			OctalPermissions: "gray",
			Permissions: PermissionColors{
				None:   "gray",
				Prefix: "gray",
				R:      "yellow",
				W:      "green",
				X:      "red",
			},
			Symlink: SymlinkColors{
				Arrow: "gray",
				Link:  "blue",
			},
			Size:            "white",
			User:            "white",
			Group:           "yellow",
			Date:            "gray",
			DirName:         "white",
			FileName:        "white",
			DirIcon:         "white",
			FileIcon:        "white",
			DirLinkIcon:     "white",
			DirLinkName:     "white",
			ExeIndicator:    "red",
			FolderIndicator: "white",
		},
	}

	config := Config{
		ShowAll:   false,
		DirsFirst: false,
		Theme:     "default",
		Themes:    theme,
	}

	// var ts Config

	json.Unmarshal(jsonData, &config)
	return config
}

func Hex2Byte(s string) string {
	b, _ := hex.DecodeString(s)
	return fmt.Sprintf("%d", b)
}
