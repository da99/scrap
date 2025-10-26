
package files

import (
	"os"
	"os/exec"
	// "path/filepath"
	"bytes"
	"bufio"
	// "strings"
)

// Returns true if the argument exists on the filesystem.
func Is(file_path string) bool {
	_, err := os.Stat(file_path);
	return os.IsNotExist(err)
}

func Bytes_To_Lines(buf *bytes.Buffer) []string {
	scanner := bufio.NewScanner(buf);
	fin := []string{}
	for scanner.Scan() {
		fin = append(fin, scanner.Text())
	}
	return fin
}

// Returns a slice of file paths, w/o directories, 1 level deep.
func List_Shallow_Files(file_path string) ( []string, error ) {
	return Cmd_To_Lines("find", file_path, "-mindepth", "1", "-maxdepth", "1", "-type", "f")
}

// The same as List_Shallow_Files, but with ext argument: "*.go.html", "*.js", etc
func List_Shallow_Files_Ext(file_path string, ext string) ( []string, error ) {
	return Cmd_To_Lines("find", file_path, "-mindepth", "1", "-maxdepth", "1", "-type", "f", "-iname", ext)
}

// Returns a string slice of directory paths, 1 level deep, ignoring . directories.
func List_Shallow_Dirs(file_path string) ( []string, error ) {
	return Cmd_To_Lines("find", file_path, "-mindepth", "1", "-maxdepth", "1", "-type", "d", "-not", "-name", ".*")
}

// Returns a string slice of output from a command executed via exec.Command.
func Cmd_To_Lines(cmd string, args ...string) ([]string, error) {
	os_cmd := exec.Command(cmd, args...)
	var out bytes.Buffer
	os_cmd.Stdout = &out
	os_cmd.Stderr = os.Stderr;
	err := os_cmd.Run()
	if err != nil {
		return nil, err
	}
	return Bytes_To_Lines(&out), nil
}
