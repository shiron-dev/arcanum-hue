//go:build tools
// +build tools

package main

import (
	_ "github.com/cweill/gotests"
	_ "go.uber.org/mock/mockgen"
)
