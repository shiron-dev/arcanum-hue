package vscode

import _ "embed"

//go:embed README.md.go.tpl
var README string

//go:embed package.json.go.tpl
var PackageJSON string

//go:embed .vscodeignore
var VSCodeIgnore string

//go:embed CHANGELOG.md
var CHANGELOG string

//go:embed launch.json
var LaunchJSON string
