package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pmezard/go-difflib/difflib"
)

var (
	// current and target are basically
	// the versions you want to compare
	current, target, directory string
)

func init() {
	flag.StringVar(&current, "current", "", "name and version of the package")
	flag.StringVar(&target, "target", "", "name and targeted version to compare")
	flag.StringVar(&directory, "directory", "./rdeps", "location of rdeps directory (https://qa-reports.gentoo.org/output/genrdeps/)")
}

func main() {
	flag.Parse()
	if len(current) == 0 || len(target) == 0 {
		fmt.Println("usage: ddiff -current sys-apps/dbus-1.12.18 -target sys-apps/dbus-1.12.20")
		os.Exit(1)
	}

	var (
		dCurrent = make([]string, 0, 0)
		dTarget  = make([]string, 0, 0)
	)

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		fmt.Printf("-directory %s is not a valid location\n", directory)
		os.Exit(1)
	}

	if err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("unable to open file: %w", err)
		}

		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			txt := scanner.Text()
			if strings.Contains(txt, current) {
				// the path looks like this: rdeps/rindex/sys-libs/liburing
				// only the sys-libs/liburing is relevant (name of the package)
				// so we splitN: [rdeps rindex sys-libs/liburing]
				entry := strings.SplitN(path, "/", 3)
				dCurrent = append(dCurrent, entry[2])
			}
			if strings.Contains(txt, target) {
				entry := strings.SplitN(path, "/", 3)
				dTarget = append(dTarget, entry[2])
			}
		}
		if err := scanner.Err(); err != nil {
			return fmt.Errorf("unable to scan file: %w", err)
		}
		return nil
	}); err != nil {
		fmt.Printf("unable to walk correctly into %s: %v\n", directory, err)
		os.Exit(1)
	}

	diff := difflib.UnifiedDiff{
		A:        dCurrent,
		B:        dTarget,
		FromFile: "Current",
		ToFile:   "Target",
	}

	ddiff, _ := difflib.GetUnifiedDiffString(diff)
	fmt.Println(ddiff)
}
