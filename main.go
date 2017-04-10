// KICKOFF - Project Bootstrapping Tool
// Copyright (c) 2017 Philipp Mieden <dreadl0ck [at] protonmail [dot] ch>

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
	"errors"
	"os"
	"runtime"

	"github.com/Sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var (
	// Log instance
	Log = logrus.New()

	// path for config directory
	configDirPath string

	// ErrTooManyArguments means KICKOFF received more than 3 parameters
	ErrTooManyArguments = errors.New("too many arguments passed")

	// mapped template names to their path
	templates = make(map[string]string, 0)

	directorySeparator = "/"
)

const debug = false

func init() {
	// set up the prefixed formatter
	Log.Formatter = new(prefixed.TextFormatter)
}

func main() {

	var (
		cLog         = Log.WithField("prefix", "main")
		projectDir   string
		template     = "default"
		templatePath string
		ok           bool
	)

	if debug {
		Log.Level = logrus.DebugLevel
	}

	printASCII()

	if runtime.GOOS == "windows" {
		directorySeparator = "\\"
	}

	configDirPath = os.Getenv("HOME") + directorySeparator + ".kickoff"

	// parse config dir and setup templates
	readConfigDirectory()

	// handle args
	switch true {
	case len(os.Args) < 2:
		printHelp()
	case os.Args[1] == "-h" || os.Args[1] == "help" || os.Args[1] == "-help":
		printHelp()
	case len(os.Args) == 3: // template was specified
		template = os.Args[1]
		projectDir = os.Args[2]
	case len(os.Args) == 2: // use default template
		projectDir = os.Args[1]
	default:
		printHelp()
		cLog.Fatal(ErrTooManyArguments)
	}

	// check if the template exists
	if templatePath, ok = templates[template]; !ok {
		printTemplates()
		cLog.Fatal("unknown template: ", template)
	}

	cLog.Info("creating project: ", projectDir)
	copyTemplate(templatePath, projectDir)

	cLog.Info("Happy Coding!")
	return
}
