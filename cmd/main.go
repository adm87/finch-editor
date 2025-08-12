package main

import (
	finch "github.com/adm87/finch-application/application"
	editor "github.com/adm87/finch-editor/application"
)

func main() {
	cmd := finch.NewApplicationCommand("finch-editor", editor.Application)
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
