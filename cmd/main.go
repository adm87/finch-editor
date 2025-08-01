package main

import (
	"github.com/adm87/finch-application/application"
	"github.com/adm87/finch-editor/editor"
)

func main() {
	cmd := application.NewApplicationCommand("finch-editor", editor.Application)
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
