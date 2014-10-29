# Atlas CLI

A Go-based implementation of the Atlas command line gem.  Here's the top-level help: 

```

$ atlas-cli help

NAME:
   atlas-cli - Atlas commandline API!

USAGE:
   atlas-cli [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
   login	Provide your login/API credentials
   whoami	Display your login/API credentials
   build	Build a project
   help, h	Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```

Here's the build.  The first argument is the project name (remember to include your username)

```
$ atlas-cli help build 
NAME:
   build - Build a project

USAGE:
   command build [command options] [arguments...]

OPTIONS:
   --pdf			 build a pdf
   --html			 build html format
   --epub			 build epub format
   --mobi			 build mobi format
   --branch, -b 'master'	branch to build

$ atlas-cli build odewahn/dds-field-guide --pdf --html
Working.........................
html => http://orm-atlas2-prod.s3.amazonaws.com/html/45c3725bb1dd3d4bcfa1a6d8e3145235.zip
pdf => http://orm-atlas2-prod.s3.amazonaws.com/pdf/418cbd43e30057818ff33aff1a8962c7.pdf

```