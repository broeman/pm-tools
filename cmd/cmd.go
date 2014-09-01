// Copyright 2014 Jesper Brodersen. All rights reserved.
// This code is BSD-licensed, see LICENSE file.

// Implementation of CLI command structure
package cmd

import (
	"github.com/broeman/gopack/db" // using DB struct for placeholder REMOVE
	//"github.com/broeman/gopack/pack" // using Package struct
	"github.com/codegangsta/cli"
)

var Init = cli.Command{
	Name:        "init",
	Usage:       "Placeholder initialization",
	Description: "Placeholdering",
	Action:      runPlaceholder,
	Flags:       []cli.Flag{},
}

// Init DB placeholder, should be a setup function REMOVE
func runPlaceholder(ctx *cli.Context) {
	db.InitDB()
}
