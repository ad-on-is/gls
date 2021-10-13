# 🚦 gLS

## A fast, colorful and higly customizable `ls` written in Go

![Long](https://i.imgur.com/4dmiO3l.png)

---

## ✔️ Fast

### Provides fast output within just ms

## ✔️ Customizable

### Checks nearly every box, when it comes to customization via `json`

## ✔️ Git status

### Shows the status of untracked/modified/deleted files in Git

---

## Usage

```
usage: gls [<flags>] [<path>]

Flags:
--help Show context-sensitive help (also try --help-long and --help-man).
  -l, --long        Show long output
  -a, --all         Show hidden
  -t, --tree        Show tree view
  -n, --nest=2      Nest level for tree view
  -g, --group       Show group next to user
  -G, --git         Show Git status
  -o, --octal       Show octal permissions, ie 0755
  -d, --dirs-first  Show directories first

Args:
[<path>] The path or file to show
```

# Screenshots

## Simple

![Simple](https://i.imgur.com/dpI0v2w.png)

## Long

![Long](https://i.imgur.com/4dmiO3l.png)

## Tree

![Tree](https://i.imgur.com/RsTKQ9q.png)

## Git status and exlucde folders in tree-view

![Git and excluded](https://i.imgur.com/sylSGsm.png)

# Configuration [gls.json Example](https://github.com/ad-on-is/gls/blob/main/.config/gls.json)

Create your custom configuratoin file located at `~/.config/gls.json`.

```json
{
  "theme": "default",
  "displayInfos": true, // whether to show notification if <path> does not exist or <path> is empty
  "excludeDirs": ["node_modules", ".git"], // dirs to exclude (in tree view)
  "themes": {
    "default": {
      "dateFormat": "Mon 2006-01-02 15:04:05", // date format
      "folderSuffix": "", // indicate folders
      "folderPrefix": "",
      "exeSuffix": " ●", // indicate executables
      "exePrefix": "",
      "gitPrefix": {
        // indicate git status
        "m": "M",
        "d": "D",
        "u": "U"
      },

      /* 
      Some file and folder types have special icon colors. Specify how to style them
      */
      "specialColorizeDirs": false, // apply special color dir name
      "specialColorizeFiles": false, // ... to filename
      "specialColorizeDirIcons": true, // ... to file icon
      "specialColorizeFileIcons": true, // ... to dir icon
      "specialColorizeLinks": true, // ... to symlink
      "dimSpecialColorizeLinks": true, // ... make symlink special color slightli dimmer

      /*
        when using --git
      */
      "colorizeGitIcon": true, // colorize icon
      "colorizeGitName": false, // colorize name
      "colors": {
        "info": "gray", // info message if displayInfos set to true
        "octalPermissions": "gray",
        "permissions": {
          "none": "gray", // - (dash)
          "prefix": "gray", // L or d
          "r": "white",
          "w": "gray",
          "x": "red"
        },
        "git": {
          "m": "yellow",
          "d": "red",
          "u": "green"
        },
        "symlink": {
          "arrow": "gray",
          "link": "gray" // if specialColorizeLinks is set to false
        },
        "excluded": "gray", // excluded indicator
        "size": "gray",
        "tree": "gray", // tree branches
        "user": "white",
        "group": "white",
        "date": "gray",
        "dirName": "yellow",
        "fileName": "white",
        "dirIcon": "yellow",
        "fileIcon": "white",
        "dirLinkIcon": "white",
        "dirLinkName": "white",
        "hiddenDirIcon": "white",
        "hiddenDirName": "white",
        "hiddenFileIcon": "white",
        "hiddenFileName": "white",
        "exeIndicator": "red",
        "folderIndicator": "white"
      }
    }
  }
}
```
