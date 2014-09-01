// Copyright 2014 Jesper Brodersen. All rights reserved.
// This code is BSD-licensed, see LICENSE file.

// Implementation of CLI command structure
package cmd

import (
	"fmt"
	"github.com/broeman/gopack/db" // using DB struct for placeholder REMOVE
	//"github.com/broeman/gopack/pack" // using Package struct
	"github.com/codegangsta/cli"
)

var Init = cli.Command{
	Name:        "init",
	Usage:       "Placeholder: Initialization",
	Description: "Initialize the Database",
	Action:      runPlaceholder,
	Flags:       []cli.Flag{},
}

var Installed = cli.Command{
	Name:        "installed",
	Usage:       "Placeholder: Shows installed packages",
	Description: `Shows all currently installed packages`,
	Action:      runInstalled,
	Flags:       []cli.Flag{},
}

// Init DB placeholder, should be a setup function REMOVE
func runPlaceholder(ctx *cli.Context) {
	db.InitDB()
}

// Placeholder to see if things works REMOVE
func runInstalled(*cli.Context) {
	rows, err := db.QueryAllPackages()
	if err != nil {
		fmt.Println("Query Error: : %v\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var versionregex string
		var installed bool
		rows.Scan(&id, &name, &versionregex, &installed)
		fmt.Println(id, name, installed)
	}
	rows.Close()
}
