/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"path"
	"github.com/spf13/cobra"
	"strings"
	"os/exec"
	"bytes"
	"bufio"
	"path/filepath"
	"da99/cli/files"
	"errors"
)

type FileHandler func(string) error
const PARTIAL_PATTERN = ".partial.go.html"

func print_it(str string) error {
	_, err := fmt.Printf("File: %v\n", str)
	return err
}

func each_line(matches []string, fh FileHandler) error {
	for _, v := range matches {
		err := fh(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func each_line_os(buf *bytes.Buffer, fh FileHandler) error {
	scanner := bufio.NewScanner(buf);
	for scanner.Scan() {
		new_path := scanner.Text()
		if strings.Contains(new_path, PARTIAL_PATTERN) {
			continue;
		}
		err := fh(new_path)
		if err != nil {
			return err
		}
	}

	return scanner.Err()
}


func CompileDir(target string, fh FileHandler) error {
	stuff, err := os.ReadDir(target)
	if err != nil { return err }
	for _, entry := range stuff {
		if entry.IsDir() && files.Is(path.Join(target, "index.go.html")) {
			fh(entry.Name())
		}
	}

	return nil
}

func LsFilesOS(target string, fh FileHandler) error {
	os_cmd := exec.Command("find", target, "-name", "*.go.html")
	var out bytes.Buffer
	var stderr bytes.Buffer
	os_cmd.Stdout = &out
	os_cmd.Stderr = &stderr
	err := os_cmd.Run()
	if err != nil {
		fmt.Println(stderr.String())
		return err
	}
	return each_line_os(&out, fh)
}

func LsFiles(target string) ([]string, error) {
	return filepath.Glob(filepath.Join(target, "/**/*.go.html"))
}


// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use  : "ls [files|dir] [target]",
	Args : cobra.ExactArgs(2),
	Short: "Lists files with extension: .go.html",
	RunE : func(cmd *cobra.Command, args []string) error {
		switch args[0] {
		case "dirs":
			dirs, err := files.List_Shallow_Dirs(args[1])
			if err != nil { return err; }
			for _, v := range dirs { fmt.Println(v) }
		case "files":
			files, err := files.List_Shallow_Files_Ext(args[1], "*.go.html")
			if err != nil { return err; }
			for _, v := range files { fmt.Println(v) }
		default:
			return errors.New("Invalid option: ls '" + args[0] + "' '" + args[1] + "'")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
