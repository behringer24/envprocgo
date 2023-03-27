// Copyright © 2023 Andreas Behringer <abe@activecube.de>
//
// Licensed under the GNU GENERAL PUBLIC LICENSE, Version 3.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://fsf.org/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/behringer24/argumentative"
)

const (
	title       = "envproc"
	description = "Config file preprocessor, inject environment variables into static config files"
	version     = "v1.0.4"
)

var (
	placeholderIndicator *string
	force                *bool
	inFileName           *string
	outFileName          *string

	infile  *os.File
	outfile *os.File
)

func parseArgs() {
	flags := &argumentative.Flags{}
	placeholderIndicator = flags.Flags().AddString("char", "c", false, "$", "Another description")
	force = flags.Flags().AddBool("force", "f", "Show version information")
	inFileName = flags.Flags().AddPositional("infile", true, "", "File to read from")
	outFileName = flags.Flags().AddPositional("outfile", false, "", "File to write to")

	showHelp := flags.Flags().AddBool("help", "h", "Show this help text")
	showVer := flags.Flags().AddBool("version", "v", "Show version information")

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

func parseRow(row string) (string, error) {
	r, _ := regexp.Compile(fmt.Sprintf("\\%s\\{env\\:([a-zA-Z0-9]+)\\}", *placeholderIndicator))
	matches := r.FindAllStringSubmatch(row, -1)
	for _, match := range matches {
		env, exists := os.LookupEnv(match[1])

		if !*force && !exists {
			return row, fmt.Errorf("found placeholder %s with name %s but no matching environment variable", match[0], match[1])
		}

		row = strings.ReplaceAll(row, match[0], env)
	}
	return row, nil
}

func parser() int {
	var err error

	// open infile config template
	infile, err = os.OpenFile(*inFileName, os.O_RDONLY, 0755)
	if err != nil {
		fmt.Printf("Cannot open %s for reading. %s\n", *inFileName, err)
		return 1
	}
	defer infile.Close()

	// open outfile or stdout
	if *outFileName != "" {
		if outfile, err = os.OpenFile(*outFileName, os.O_RDWR|os.O_CREATE, 0755); err != nil {
			fmt.Printf("Cannot open %s for writing. %s\n", *outFileName, err)
			return 1
		}
	} else {
		outfile = os.Stdout
	}
	defer outfile.Close()

	// Scan each line in file
	scanner := bufio.NewScanner(infile)
	for scanner.Scan() {
		row := scanner.Text()
		newrow, err := parseRow(row)
		if err != nil {
			fmt.Println(err)
			return 1
		}
		outfile.WriteString(newrow + "\n")
	}

	// handle scanner errors
	if err = scanner.Err(); err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}

func main() {
	parseArgs()
	os.Exit(parser())
}
