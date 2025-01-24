package static

import "embed"

//go:embed index.html
var StaticIndex embed.FS
