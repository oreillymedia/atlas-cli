# Atlas CLI

A Go-based Atlas API client for building projects and publishing them as a web site.

Here's the top-level help: 

```
NAME:
   atlas - OºReilly Atlas command line tool

USAGE:
   atlas [global options] command [command options] [arguments...]

VERSION:
   0.0.8-alpha

COMMANDS:
   login	Set your login/API credentials
   whoami	Display your login/API credentials
   info		Display info about your Atlas project based on the git config file
   build	Build a project
   open		Open a site
   publish	Publish a site
   help, h	Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```

## Installation

* Download the latest release
* Put it on your path
* Do `chmod + x` on it

Someday I'll make a packager.


## Usage 

With this command, you can do things like this:


* `atlas build -p odewahn/dds-field-guide --html`

* `atlas publish -p odewahn/dds-field-guide --public`

* `atlas open -p odewahn/dds-field-guide --public`

If you omit the "-p" flag, the CLI will see if it can find a project name from the remotes defined in your git config.  So, if you're in your project's home directory, you can just do this:

* `atlas build --html`

* `atlas publish --public`

* `atlas open --public`


## Development

To build, do this:

```
go build -o atlas *.go
```

