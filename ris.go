// SPDX-License-Identifier: BSD-3-Clause
// SPDX-FileCopyrightText: Copyright (c) 2025 commitmaniac

package main

import (
	"fmt"
	"path/filepath"
	"os"
)

func main() {
	path := filepath.Join(os.Args[2], os.Args[1])
	files, err := filepath.Glob(path)
	if err != nil {
		panic(err)
	}

	for i, file := range files {
		seq := fmt.Sprintf("%04d", i + 1)
		filename := fmt.Sprintf("%s%s", seq, filepath.Ext(file))

		newfile := filepath.Join(os.Args[2], filename)
		err = os.Rename(files[i], newfile)
		if err != nil {
			panic(err)
		}

		println(files[i], "=>", newfile)
	}
}
