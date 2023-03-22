/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"

	"github.com/behringer24/argumentative"
)

const (
	title       = "envproc"
	description = "Config file preprocessor, inject environment variables into static config files"
	version     = "v0.0.1"
)

var (
	placeholderIndicator *string
	showHelp             *bool
	showVer              *bool
	force                *bool
	inFileName           *string
	outFileName          *string
)

func parseArgs() {
	flags := &argumentative.Flags{}
	placeholderIndicator = flags.Flags().AddString("char", "c", false, "%", "Another description")
	showHelp = flags.Flags().AddBool("help", "h", "Show this help text")
	showVer = flags.Flags().AddBool("version", "v", "Show version information")
	force = flags.Flags().AddBool("force", "f", "Show version information")
	inFileName = flags.Flags().AddPositional("infile", true, "", "File to read from")
	outFileName = flags.Flags().AddPositional("outfile", false, "", "File to write to")

	err := flags.Parse(os.Args)
	if *showHelp {
		flags.Usage(title, description, nil)
		os.Exit(0)
	} else if *showVer {
		fmt.Println(title, "version", version)
		os.Exit(0)
	} else if err != nil {
		flags.Usage(title, description, err)
		os.Exit(1)
	}
}

func main() {
	parseArgs()

	fmt.Println("\nResult:", *force, *placeholderIndicator, *inFileName, *outFileName, *showHelp, *showVer)
}
