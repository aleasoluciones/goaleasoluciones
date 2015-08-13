// Copyright 2014 Alea Soluciones SLL. All rights reserved.  Use of this
// source code is governed by a MIT-style license that can be found in the
// LICENSE file.

package main

import (
	"log"

	"github.com/aleasoluciones/goaleasoluciones/periodictask"
)

func main() {
	periodictask.PeriodicTask(func() { log.Println("TASK") }, "* * * * * *")
}
