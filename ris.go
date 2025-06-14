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
}

var opts Options

func main() {
	err := flagsfiller.Parse(&opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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

	for i, file := range files {
		seq := fmt.Sprintf("%04d", i + 1)
		filename := fmt.Sprintf("%s%s", seq, filepath.Ext(file))

		if opts.RenameTo != "" {
			filename = strings.Replace(filename, filepath.Ext(file), opts.RenameTo, 1)
		}

		newfile := filepath.Join(args[0], filename)
		err = os.Rename(files[i], newfile)
		if err != nil {
			panic(err)
		}

		println(files[i], "=>", newfile)
	}
}
