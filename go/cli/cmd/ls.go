/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	// "os"
	"os/exec"
	"bytes"
	"bufio"
)

type FileHandler func(string) error

func print_it(str string) error {
	_, err := fmt.Printf("File: %v\n", str)
	return err
}

func each_line(buf *bytes.Buffer, fh FileHandler) error {
	scanner := bufio.NewScanner(buf);
	for scanner.Scan() {
		err := fh(scanner.Text())
		if err != nil {
			return err
		}
	}

	return scanner.Err()
}

func LsFiles(target string, fh FileHandler) error {
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
	return each_line(&out, fh)
}


// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use  : "ls [dir]",
	Args : cobra.ExactArgs(1),
	Short: "Lists files with extension: .go.html",
	RunE : func(cmd *cobra.Command, args []string) error {
		return LsFiles(args[0], print_it)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
