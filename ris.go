// SPDX-License-Identifier: BSD-3-Clause
// SPDX-FileCopyrightText: Copyright (c) 2025 commitmaniac

package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"os"
	"strings"

	"github.com/itzg/go-flagsfiller"
)

type Options struct {
	Select   string `usage:"Select which files to rename"`
	RenameTo string `usage:"Use preferred extension"`
	Index    int    `usage:"Start at provided index" default:"1"`
	Version  bool   `usage:"Print installed version"`
	KeepName bool   `usage:"Keep original filename"`
	Zfill    int    `usage:"Specify preferred zfill" default:"4"`
	Offset   int    `usage:"Offset sequence by a specific multiple" default:"1"`
	Simulate bool   `usage:"Don't actually rename selected files"`
}

var (
	opts Options
	Version string
)

func RenameFiles(target string, newfile string) {
	err := os.Rename(target, newfile)
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] PATH\n", os.Args[0])
		flag.PrintDefaults()
	}

	err := flagsfiller.Parse(&opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if opts.Version {
		fmt.Println(os.Args[0], Version)
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(0)
	}

	path := filepath.Join(args[0], opts.Select)
	files, err := filepath.Glob(path)
	if err != nil {
		panic(err)
	}

	if opts.Simulate {
		fmt.Println("WARNING: simulating, no files renamed")
	}

	for position, file := range files {
		basename := strings.TrimSuffix(file, filepath.Ext(file))
		sequence := fmt.Sprintf("%0*d", opts.Zfill, (position + opts.Index) * opts.Offset)

		newname := fmt.Sprintf("%s%s", sequence, filepath.Ext(file))
		if opts.KeepName {
			newname = fmt.Sprintf("%s_%s%s", sequence, basename, filepath.Ext(file))
		}

		if opts.RenameTo != "" {
			newname = strings.Replace(newname, filepath.Ext(file), opts.RenameTo, 1)
		}

		newfile := filepath.Join(args[0], newname)
		if opts.Simulate {
			fmt.Println(file, "=>", newfile)
		} else {
			RenameFiles(file, newfile)
		}
	}
}
