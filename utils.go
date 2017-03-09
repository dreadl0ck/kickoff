// KICKOFF - Project Bootstrapping Tool
// Copyright (c) 2017 Philipp Mieden <dreadl0ck@protonmail.ch>

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

/*
 *	Utils
 */

func printHelp() {
	fmt.Println("usage: kickoff [template] <projectName>")
}

func printTemplates() {

	var maxLength int

	// determine maxLength for padding
	for name := range templates {
		if len(name) > maxLength {
			maxLength = len(name)
		}
	}

	for name, path := range templates {
		fmt.Println(pad(name, maxLength+1), "~>", path)
	}
}

func copyTemplate(templatePath, projectDir string) {

	var cLog = Log.WithField("prefix", "copyTemplate")

	// create project directory
	err := os.Mkdir(projectDir, 0700)
	if err != nil {
		cLog.WithError(err).Fatal("failed to create project directory")
	}

	err = filepath.Walk(templatePath, func(path string, info os.FileInfo, err error) error {

		if path != templatePath { // ignore self

			if strings.HasSuffix(path, ".git") {

				// change directory
				err := os.Chdir(projectDir)
				if err != nil {
					cLog.WithError(err).Fatal("failed to change directory to projectDir")
				}

				// initialize git
				gitInit()

				// change back one level up
				err = os.Chdir("..")
				if err != nil {
					cLog.WithError(err).Fatal("failed to change directory")
				}

				return nil
			}

			// assemble relative path
			relativePath := projectDir + strings.TrimPrefix(path, templatePath)

			if info.IsDir() {
				createDirectory(relativePath)
			} else {
				copyFile(path, relativePath)
			}
		}
		return nil
	})
	if err != nil {
		Log.WithError(err).Fatal("failed to walk " + configDirPath)
	}
}

// parse config directory and init templateMap
func readConfigDirectory() {

	err := filepath.Walk(configDirPath, func(path string, info os.FileInfo, err error) error {

		// ignore self and default directory
		if path != configDirPath {
			if info.IsDir() {
				// add to templates
				templates[filepath.Base(path)] = path
				return filepath.SkipDir
			}
		}
		return nil
	})
	if err != nil {
		Log.WithError(err).Fatal("failed to walk " + configDirPath)
	}
}

// copy a file from source to destination
func copyFile(source, destination string) {

	var cLog = Log.WithField("prefix", "copyFile")
	cLog.Info(source, " ~> ", destination)

	// read file contents
	contents, err := ioutil.ReadFile(source)
	if err != nil {
		cLog.WithError(err).Fatal("failed to read: ", source)
	}

	// create new file
	f, err := os.Create(destination)
	if err != nil {
		cLog.WithError(err).Fatal("failed to create file: ", destination)
	}
	defer f.Close()

	// write contents to file
	_, err = f.Write(contents)
	if err != nil {
		cLog.WithError(err).Fatal("failed to write contents to file: ", destination)
	}
}

// initialize an empty git repository
func gitInit() {

	var cLog = Log.WithField("prefix", "gitInit")

	out, err := exec.Command("git", "init").CombinedOutput()
	if err != nil {
		cLog.WithError(err).Fatal("git init failed. output: ", string(out))
	}

	cLog.Info("initialized git repository")
}

// create an empty file, fatals if something goes wrong
func createFile(name string) {

	var cLog = Log.WithField("prefix", "createFile")

	cLog.Info("creating file: ", name)
	f, err := os.Create(name)
	if err != nil {
		cLog.WithError(err).Fatal("failed to create file: ", name)
	}
	err = f.Close()
	if err != nil {
		cLog.WithError(err).Fatal("failed to close file handle: ", name)
	}
}

// create an empty directory, fatals if something goes wrong
func createDirectory(path string) {

	var cLog = Log.WithField("prefix", "createDirectory")

	cLog.Info("creating directory: ", path)

	err := os.Mkdir(path, 0700)
	if err != nil {
		cLog.WithError(err).Fatal("failed to create directory: ", path)
	}
}

// print directory as a tree structure
func printTree(dir string) {

	var (
		cLog            = Log.WithField("prefix", "printTree")
		files           int
		dirs            int
		directoryPhrase = "directories, "
		filePhrase      = "files"
		offset          = 4
	)

	fmt.Println(dir)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		// skip git dir and self
		if !strings.Contains(path, ".git") && path != dir {

			path = strings.TrimPrefix(path, dir+"/")
			level := len(strings.Split(path, directorySeparator)) - 1

			if info.IsDir() {
				dirs++
				fmt.Println("└── " + path)
			} else {
				files++
				fmt.Println(fmt.Sprintf("%-"+strconv.Itoa(level*offset)+"s", "") + "├── " + path)
			}
		}
		return nil
	})
	if err != nil {
		cLog.WithError(err).Fatal("failed to walk directory")
	}

	if dirs == 1 {
		directoryPhrase = "directory, "
	}

	if files == 1 {
		filePhrase = "file"
	}

	println()
	fmt.Println(dirs, directoryPhrase, files, filePhrase)
	println()
}

// pad the input string up to the given number of space characters
func pad(in string, length int) string {
	if len(in) < length {
		return fmt.Sprintf("%-"+strconv.Itoa(length)+"s", in)
	}
	return in
}

func printASCII() {
	ascii := `
     __   .__        __          _____  _____
    |  | _|__| ____ |  | _______/ ____\/ ____\
    |  |/ /  |/ ___\|  |/ /  _ \   __\\   __\
    |    <|  \  \___|    <  <_> )  |   |  |
    |__|_ \__|\___  >__|_ \____/|__|   |__|
         \/       \/     \/   Project Bootstrapping Tool
	`
	fmt.Println(ascii)
}
