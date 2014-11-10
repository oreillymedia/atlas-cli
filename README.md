# O'Reilly CLI

A Go-based oreilly API client.  Here's the top-level help: 

```
$ oreilly help

NAME:
   oreilly - OºReilly command line API

USAGE:
   oreilly [global options] command [command options] [arguments...]

VERSION:
   0.0.4-alpha

COMMANDS:
   login	Set your login/API credentials
   whoami	Display your login/API credentials
   atlas	Work with Atlas projects
   product	Product metadata and ownership server (requires admin access)
   sites	Publish an Atlas project to sites.oreilly.com
   help, h	Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version


```

## Installation

Put the `oreilly` binary somewhere on your path, and BOOM!  You're ready to go.

## Usage 

With this command, you can do things like this:


* `oreilly atlas build odewahn/dds-field-guide --html`

* `oreilly sites publish odewahn/dds-field-guide --public`

* `oreilly sites open odewahn/dds-field-guide --public`

* `oreilly product find "go programming"`

* `oreilly product grant 9781491913871.VIDEO rune@runemadsen.com`


## Building

To build, do this:

```
go build -o oreilly *.go
```

