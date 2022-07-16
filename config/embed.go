package config

import "embed"

//go:embed app.ini app.release.ini
var Files embed.FS
