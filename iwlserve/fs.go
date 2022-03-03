package main

import "embed"

//go:embed static templates data
var embeddedFS embed.FS
