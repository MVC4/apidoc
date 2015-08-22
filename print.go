// Copyright 2015 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"runtime"

	"github.com/issue9/term/colors"
)

const (
	out          = colors.Stdout
	titleColor   = colors.Green
	contentColor = colors.Default
	errorColor   = colors.Red
	warnColor    = colors.Cyan
)

func printError(msg ...interface{}) {
	colors.Println(out, errorColor, colors.Default, msg...)
}

func printSyntaxErrors(errs []error) {
	for _, v := range errs {
		colors.Println(out, warnColor, colors.Default, v)
	}
}

func printLangs() {
	colors.Println(out, titleColor, colors.Default, "目前支持以下类型的代码解析:")
	for k, v := range langs {
		colors.Print(out, titleColor, colors.Default, k, ":")
		colors.Println(out, contentColor, colors.Default, v.exts)
	}
}

func printVersion() {
	colors.Print(out, titleColor, colors.Default, "apidoc: ")
	colors.Println(out, contentColor, colors.Default, version)

	colors.Print(out, titleColor, colors.Default, "Go: ")
	goVersion := runtime.Version() + " " + runtime.GOOS + "/" + runtime.GOARCH
	colors.Println(out, contentColor, colors.Default, goVersion)
}