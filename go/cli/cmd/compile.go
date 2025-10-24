/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"html/template"
	"github.com/spf13/cobra"
	"os"
	"errors"
	"encoding/json"
)

func GetConfigFile() (string, error) {
	if _, err := os.Stat("config.json"); !os.IsNotExist(err) {
		return "config.json", nil
	}
	if _, err := os.Stat("config/main.json"); !os.IsNotExist(err) {
		return "config/main.json", nil
	}
	return "", errors.New("Config file not found.")
}

func GetConfig() (map[string]interface{}, error) {
	fin := make(map[string]interface{})

	config_file, config_err := GetConfigFile()
	if config_err != nil {
		return fin, nil
	}
	contents, read_err := os.ReadFile(config_file)

	if read_err != nil {
		return fin, read_err;
	}

	j_err := json.Unmarshal([]byte(contents), &fin)
	if j_err != nil {
		return fin, nil
	}

	return fin, nil
}

func CompileFile(fp string) error {
	fmt.Println("Compiling: " + fp)
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		fmt.Println("errorr")
		return err
	}
	return tmpl.Execute(os.Stdout, "http://www.lewrockwell.com/>a?")
}

// compileCmd represents the compile command
var compileCmd = &cobra.Command{
	Use:   "compile [dir]",
	Short: "Compile files from `ls [dir]`",
	Args : cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return LsFiles(args[0], CompileFile)
	},
}

func init() {
	rootCmd.AddCommand(compileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
