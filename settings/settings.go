package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Themes       map[string]Theme `json:"themes"`
	Theme        string           `json:"theme"`
	DisplayInfos bool             `json:"displayInfos"`
	ExcludeDirs  []string         `json:"excludeDirs"`
	ShowAll      bool
	ShowGroup    bool
	ShowOctal    bool
	ShowGit      bool
	DirsFirst    bool
}

type Theme struct {
	DateFormat               string    `json:"dateFormat"`
	FolderSuffix             string    `json:"folderSuffix"`
	FolderPrefix             string    `json:"folderPrefix"`
	ExeSuffix                string    `json:"exeSuffix"`
	ExePrefix                string    `json:"exePrefix"`
	SpecialColorizeLinks     bool      `json:"specialColorizeLinks"`
	DimSpecialColorizeLinks  bool      `json:"dimSpecialColorizeLinks"`
	SpecialColorizeDirs      bool      `json:"specialColorizeDirs"`
	SpecialColorizeFiles     bool      `json:"specialColorizeFiles"`
	SpecialColorizeDirIcons  bool      `json:"specialColorizeDirIcons"`
	SpecialColorizeFileIcons bool      `json:"specialColorizeFileIcons"`
	ColorizeGitIcon          bool      `json:"colorizeGitIcon"`
	ColorizeGitName          bool      `json:"colorizeGitName"`
	Colors                   Colors    `json:"colors"`
	GitPrefix                GitPrefix `json:"gitPrefix"`
}

type Colors struct {
	OctalPermissions string           `json:"octalPermissions"`
	Permissions      PermissionColors `json:"permissions"`
	Symlink          SymlinkColors    `json:"symlink"`
	Git              GitColors        `json:"git"`
	Excluded         string           `json:"excluded"`
	Size             string           `json:"size"`
	Tree             string           `json:"tree"`
	Info             string           `json:"info"`
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

type GitColors struct {
	M string `json:"m"`
	D string `json:"d"`
	U string `json:"u"`
}

type GitPrefix struct {
	M string `json:"m"`
	D string `json:"d"`
	U string `json:"u"`
}

type SymlinkColors struct {
	Arrow string `json:"arrow"`
	Link  string `json:"link"`
}

func GetConfig() Config {

	theme := make(map[string]Theme)

	theme["default"] = Theme{
		DateFormat:               "Mon 2006-01-02 15:04:05",
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
		ColorizeGitIcon:          true,
		ColorizeGitName:          true,
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
			Git: GitColors{
				M: "yellow",
				D: "red",
				U: "green",
			},
			Size:            "gray",
			Excluded:        "gray",
			Tree:            "gray",
			Info:            "gray",
			User:            "blue",
			Group:           "magenta",
			Date:            "gray",
			DirName:         "yellow",
			DirIcon:         "yellow",
			FileName:        "white",
			FileIcon:        "white",
			DirLinkIcon:     "green",
			DirLinkName:     "green",
			ExeIndicator:    "red",
			FolderIndicator: "white",
		},
	}

	config := Config{
		DisplayInfos: true,
		ShowAll:      false,
		DirsFirst:    false,
		Theme:        "default",
		Themes:       theme,
		ExcludeDirs:  []string{},
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		return config
	}
	settingsFile := homedir + "/.config/gls.json"
	jsonFile, err := os.Open(settingsFile)
	if err != nil {
		return config
	}

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err == nil {
		json.Unmarshal(jsonData, &config)
	}
	fmt.Println(config)
	return config
}
