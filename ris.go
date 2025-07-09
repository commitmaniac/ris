// SPDX-License-Identifier: BSD-3-Clause
// SPDX-FileCopyrightText: Copyright (c) 2025 commitmaniac

package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/itzg/go-flagsfiller"
)

type Options struct {
	Index    int    `usage:"Start at provided index" default:"1"`
	KeepName bool   `usage:"Keep original filename"`
	Offset   int    `usage:"Offset sequence by a specific multiple" default:"1"`
	RenameTo string `usage:"Use preferred file extension"`
	Reverse  bool   `usage:"Reverse list of selected files"`
	Select   string `usage:"Select which files to rename"`
	Simulate bool   `usage:"Simulate operations with selected files"`
	Version  bool   `usage:"Print installed version"`
	Zfill    int    `usage:"Use preferred zfill" default:"4"`
}

var (
	opts Options
	Version string
)

func KeepFilename(base string, sequence string, ext string) string {
	sequence_prefix := fmt.Sprintf("%s_", sequence)
	if strings.HasPrefix(base, sequence_prefix) {
		base = strings.TrimPrefix(base, sequence_prefix)
	}

	return fmt.Sprintf("%s_%s%s", sequence, base, ext)
}

func RenameFiles(target string, newfile string) {
	_, err := os.Stat(newfile)
	if errors.Is(err, os.ErrNotExist) {
		err = os.Rename(target, newfile)
	}

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

	path_glob := filepath.Join(args[0], opts.Select)
	files, err := filepath.Glob(path_glob)
	if err != nil {
		panic(err)
	}

	if opts.Reverse {
		slices.Reverse(files)
	}

	if opts.Simulate {
		fmt.Println("WARNING: simulating, no files renamed")
	}

	for position, file := range files {
		file_ext := filepath.Ext(file)
		basename := strings.TrimSuffix(file, file_ext)
		int_sqnc := (position + opts.Index) * opts.Offset
		sequence := fmt.Sprintf("%0*d", opts.Zfill, int_sqnc)

		newname := fmt.Sprintf("%s%s", sequence, file_ext)
		if opts.KeepName {
			newname = KeepFilename(basename, sequence, file_ext)
		}

		if opts.RenameTo != "" {
			newname = strings.Replace(newname, file_ext, opts.RenameTo, 1)
		}

		newfile := filepath.Join(args[0], newname)
		if opts.Simulate {
			fmt.Println(file, "=>", newfile)
		} else {
			RenameFiles(file, newfile)
		}
	}
}
