package main

import (
    "fmt"
    "github.com/rivo/tview"
)

func main() {
    fmt.Print("lets get started chating")
    box := tview.NewBox().SetBorder(true).SetTitle("Welcome to Sodhi!")
    if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
        panic(err)
    }

}
