package ui

import "embed"

// This is not a comment, but a directive so that go knows that it has to store
// the files from htlm and static in an embed.FS referenced by the global Files
//
//go:embed "html" "static"
var Files embed.FS
