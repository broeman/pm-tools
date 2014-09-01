// Copyright 2014 Jesper Brodersen. All rights reserved.
// This code is BSD-licensed, see LICENSE file.

// pm-tools is a package manger written in Go Language
// this is the admin application for running CRUD-operations on packages,
// and settings management.
package main

import (
	"github.com/broeman/pm-tools/cmd" // using CLI command args
	"github.com/codegangsta/cli"
	"os"
	"runtime"
)

const MAN_VER = "0.1 Alpha"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	app := cli.NewApp()
	app.Name = "pm-tools"
	app.Usage = "Package Manager in Go: Toolkit"
	app.Version = MAN_VER
	app.Commands = []cli.Command{
		cmd.Init,      // initialize database if not exist, placeholder for setup
		cmd.Installed, // shows the installed packages, placeholder for testing
		//  cmd.Setup,	// settings for the user, init db
		// 	cmd.AddPackage,    // adding a package to the database
		// 	cmd.ShowPackage,   // showing a package from the database
		// 	cmd.EditPackage,   // editing a package from the database
		// 	cmd.RemovePackage, // removing a package from the database
	}
	app.Run(os.Args)

}
