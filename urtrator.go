// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016, Stanslav N. a.k.a pztrn (or p0z1tr0n)
// All rights reserved.
//
// Licensed under Terms and Conditions of GNU General Public License
// version 3 or any higher.
// ToDo: put full text of license here.
package main

import (
    // local
    "github.com/pztrn/urtrator/context"
    "github.com/pztrn/urtrator/ui"

    // stdlib
    "fmt"
)

func main() {
    fmt.Println("This is URTrator, version 0.1")

    ctx := context.New()
    ctx.Initialize()

    ui := ui.NewMainWindow(ctx)
    ui.Initialize()
}
