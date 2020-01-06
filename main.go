/*
Copyright © 2020 Ian Bibby ian.bibby@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"log"
	"path/filepath"

	"github.com/ianbibby/task/cmd"
	"github.com/ianbibby/task/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	a := &db.App{}

	p := filepath.Join(home, "tasks.db")
	err = a.Init(p)
	if err != nil {
		log.Fatal(err)
	}
	defer a.Close()

	cmd.App = a
	cmd.Execute()
}
