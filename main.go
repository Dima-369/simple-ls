package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/Gira-X/go-columnize"
	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

func splitIntoFilesAndDirs(dir []os.FileInfo, displayOnlyHidden bool) Extract {
	output := Extract{}

	for _, file := range dir {
		output.AppendIfMatching(file, displayOnlyHidden)
	}

	return output
}

func getCountString(e Extract) string {
	if e.HiddenFiles == 0 && e.HiddenDirs == 0 {
		return ""
	}

	fileString := ""
	dirsString := ""

	const pluralCount = 2

	if e.HiddenFiles == 1 {
		fileString = "1 hidden file"
	} else if e.HiddenFiles >= pluralCount {
		fileString = fmt.Sprintf("%v hidden files", e.HiddenFiles)
	}

	if e.HiddenDirs == 1 {
		dirsString = "1 hidden dir"
	} else if e.HiddenDirs >= pluralCount {
		dirsString = fmt.Sprintf("%v hidden dirs", e.HiddenDirs)
	}

	if fileString == "" {
		return " " + dirsString
	}

	if dirsString == "" {
		return " (" + fileString + ")"
	}

	return " (" + fileString + ", " + dirsString + ")"
}

func getTerminalWidthOr80() (int, error) {
	displayWidth, err := terminal.Width()
	if err != nil {
		return 0, err
	}

	if displayWidth == 0 {
		displayWidth = 80
	}

	return int(displayWidth), nil
}

func main() {
	displayOnlyHidden := false

	const hasCommandArgument = 2
	if len(os.Args) >= hasCommandArgument {
		displayOnlyHidden = os.Args[1] == "hidden"
	}

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(filepath.Clean(wd))
	if err != nil {
		panic(err)
	}

	fetchAll := -1
	dir, _ := f.Readdir(fetchAll)
	_ = f.Close()

	e := splitIntoFilesAndDirs(dir, displayOnlyHidden)

	fmt.Printf("%v%v\n\n", wd, getCountString(e))

	sort.Strings(e.Dirs)
	sort.Strings(e.Files)

	width, err := getTerminalWidthOr80()
	if err != nil {
		panic(err)
	}

	opts := columnize.DefaultOptions()
	opts.DisplayWidth = width

	if len(e.Dirs) >= 1 {
		// this is the ASCII code for a blue foreground
		print("\u001B[34m")

		out := columnize.Format(e.Dirs, opts)
		if out[len(out)-1:] != "\n" {
			// sometimes when there is only a single directory, there won't be a
			// trailing newline, so we add it here
			out += "\n"
		}

		println(out)
		// this is the ASCII code for a color reset
		print("\u001B[0m")
	}

	print(columnize.Format(e.Files, opts))
}
