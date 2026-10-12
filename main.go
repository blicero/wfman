// /home/krylon/go/src/github.com/blicero/wfman/main.go
// -*- mode: go; coding: utf-8; -*-
// Created on 12. 10. 2026 by Benjamin Walkenhorst
// (c) 2026 Benjamin Walkenhorst
// Time-stamp: <2026-10-12 17:03:33 krylon>

package main

import (
	"fmt"

	"github.com/blicero/wfman/common"
)

func main() {
	fmt.Printf("%s %s\n",
		common.AppName,
		common.Version)

	fmt.Println("Nothing to see here (yet), move along.")
}
