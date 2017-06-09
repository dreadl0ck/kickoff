# KICKOFF

     __   .__        __          _____  _____
    |  | _|__| ____ |  | _______/ ____\/ ____\
    |  |/ /  |/ ___\|  |/ /  _ \   __\\   __\
    |    <|  \  \___|    <  <_> )  |   |  |
    |__|_ \__|\___  >__|_ \____/|__|   |__|
         \/       \/     \/   Project Bootstrapping Tool

[![Go Report Card](https://goreportcard.com/badge/github.com/dreadl0ck/kickoff)](https://goreportcard.com/report/github.com/dreadl0ck/kickoff)
[![License](https://img.shields.io/aur/license/yaourt.svg)](https://raw.githubusercontent.com/dreadl0ck/kickoff/master/LICENSE)
[![Golang](https://img.shields.io/badge/Go-1.8-blue.svg)](https://golang.org)
![Linux](https://img.shields.io/badge/Supports-Linux-green.svg)
![macOS](https://img.shields.io/badge/Supports-macOS-green.svg)
![Windows](https://img.shields.io/badge/Supports-Windows-green.svg)

KICKOFF is a commandline tool for quick and easy project bootstrapping.

Have an amazing idea and want to dive straight into coding?

Setup your new project with one command!

## Install

```shell
go get -u github.com/dreadl0ck/kickoff
```

## Usage

```shell
kickoff [template] <projectname>
```

## Configuration

KICKOFF looks inside the *$HOME/.kickoff* directory for project templates when started.

Each subfolder and all its contents will serve as template and will simply be copied into a new directory with the desired project name.

If no template name was specified, the *default* folder will be used.

If you want to automatically initialize a git repository, just create an empty *.git* folder into your template.

### Configuration Examples

A simple template directory could look like this:

```text
user@host:~/.kickoff$ tree
.
├── default
│   ├── AUTHORS
│   ├── LICENSE
│   ├── Makefile
│   ├── README.md
│   └── TODO.md
├── go
│   ├── AUTHORS
│   ├── LICENSE
│   ├── Makefile
│   ├── README.md
│   ├── TODO.md
│   ├── main.go
│   └── zeus
│       ├── bench.sh
│       ├── build.sh
│       ├── clean.sh
│       ├── install.sh
│       ├── run.sh
│       └── test.sh
├── haskell
│   ├── AUTHORS
│   ├── LICENSE
│   ├── Makefile
│   ├── README.md
│   ├── TODO.md
│   └── main.hs
└── python
    ├── AUTHORS
    ├── LICENSE
    ├── Makefile
    ├── README.md
    ├── TODO.md
    └── main.py

5 directories, 29 files
```

What you see here are **4** template directories,
for the python, haskell and go programming languages and one for a default project.

## Usage Examples

Usage is pretty straight forward.
Here's a simple example for creating a new go project:

```shell
# bootstrap a new project named testProject
# use the go template
$ kickoff go testProject
...
```

## Build System

KICKOFF uses [ZEUS](github.com/dreadl0ck/zeus) as its build system.
But you can also use the go tools for compilation / installation directly.

## Project Stats

    -------------------------------------------------------------------------------
    Language                     files          blank        comment           code
    -------------------------------------------------------------------------------
    Go                               2             78             55            225
    Markdown                         2             31              0            135
    Bourne Shell                     4              8             12              9
    -------------------------------------------------------------------------------
    SUM:                             8            117             67            369
    -------------------------------------------------------------------------------

## License

```LICENSE
KICKOFF - Project Bootstrapping Tool
Copyright (c) 2017 Philipp Mieden <dreadl0ck [at] protonmail [dot] ch>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
```

## Contact

You have ideas, feedback, bugs, security issues, pull requests etc?

Contact me: dreadl0ck [at] protonmail [dot] ch

```PGP
-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: GnuPG v2

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
=sBwF
-----END PGP PUBLIC KEY BLOCK-----
```