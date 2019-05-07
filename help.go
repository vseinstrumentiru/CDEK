package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func PrintUsage() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "Usage: \t")
	fmt.Fprintln(w, "help \t show usage")
	fmt.Fprintln(w, "version \t show package version")
	fmt.Fprintln(w, "pvzlist \t get pvzlist from cdek")
	w.Flush()
}
