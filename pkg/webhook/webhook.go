// Package webhook is a stand-in module for the automation playground.
// Imports norman so go mod tidy keeps the require entry.
package webhook

import "github.com/tomleb/frameworks-automation-norman/pkg/norman"

const Version = "0.8"

func Hello() string { return "hello from webhook " + Version + " using " + norman.Hello() }
