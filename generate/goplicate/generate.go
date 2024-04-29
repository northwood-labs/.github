// Copyright 2024, Northwood Labs
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

var err error

func main() {
	varMap := map[string]bool{
		"GO":     os.Getenv("GO") == "true",
		"TF_MOD": os.Getenv("TF_MOD") == "true",
	}

	writeFileFromTemplate(
		varMap,
		getAbs("./goplicate.tmpl.yaml"),
		getAbs("../../updates/.goplicate.yaml"),
	)
}

func newTemplate(filename string) *template.Template {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		panic(err)
	}

	return tmpl
}

func getAbs(path string) string {
	s, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return s
}

func writeFileFromTemplate(varMap map[string]bool, templatePath, writePath string) {
	fmt.Println("Generating " + writePath)

	tmpl := newTemplate(templatePath)

	var f *os.File

	f, err = os.Create(writePath) // #nosec G304 -- lint:allow_possible_insecure
	if err != nil {
		panic(err)
	}

	defer func() {
		e := f.Close()
		if e != nil {
			panic(e)
		}
	}()

	err = tmpl.Execute(f, varMap)
	if err != nil {
		panic(err)
	}
}
