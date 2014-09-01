pm-tools - Package Manager in Go 
================================

## Admin Application
BSD-License applies, see LICENSE for more information.

### Purpose
Change settings, CRUD-toolkit for packages for GoPack.

### Overview
- [Documentation](http://godoc.org/github.com/broeman/pm-tools)

### Current Version 0.1a
Placeholders for init (initialize db) and installed (shows all packages installed)

### Roadmap to 0.1
- CRUD: Create, Read, Update, Delete packages

### Installation
GoPack is the library used, pm-get is useful for testing at this point:

```
$ go get github.com/broeman/gopack
$ go get github.com/broeman/pm-tools
$ go get github.com/broeman/pm-get
```

Make sure that you have set $GOPATH and source it in your PATH (e.g. ~/.bashrc):
```
export GOPATH="$HOME/go"
export PATH=$PATH:$GOPATH/bin
```

### About
Go Packge is written by [Jesper Brodersen](http://jesperbrodersen.dk)
