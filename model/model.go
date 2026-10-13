// /home/krylon/go/src/github.com/blicero/wfman/model/model.go
// -*- mode: go; coding: utf-8; -*-
// Created on 13. 10. 2026 by Benjamin Walkenhorst
// (c) 2026 Benjamin Walkenhorst
// Time-stamp: <2026-10-13 20:14:16 krylon>

// Package model provides common data types used in the application.
package model

// ID identifies an object uniquely.
type ID int64

// Step is the common interface for various types of ... well, steps that
// make up a workflow.
type Step interface {
	ID() ID
	Description() string
	Execute(args ...string) (ID, error)
}

// Workflow is a sequence or tree-like set of Steps.
type Workflow struct {
	ID    ID
	Start ID
	Steps map[ID]Step
}
