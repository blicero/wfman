// /home/krylon/go/src/github.com/blicero/wfman/logdomain/logdomain.go
// -*- mode: go; coding: utf-8; -*-
// Created on 12. 10. 2026 by Benjamin Walkenhorst
// (c) 2026 Benjamin Walkenhorst
// Time-stamp: <2026-10-12 17:01:39 krylon>

// Package logdomain provides symbolic constants for logging.
package logdomain

//go:generate stringer -type=ID

// ID represents a component of the application that wants to write to the
// log.
type ID uint8

const (
	Database ID = iota
	Engine
	Main
)

// All returns all valid values for ID.
func All() []ID {
	return []ID{
		Database,
		Engine,
		Main,
	}
} // func All() []ID
