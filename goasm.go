package main

import (
//    "os"
    "fmt"
    "syscall/js"
)

// https://blog.owulveryck.info/2018/06/08/some-notes-about-the-upcoming-webassembly-support-in-go.html#exposing-a-function

func main() {
    fmt.Printf("hallo goasm\n")
/*
    var cb js.Callback
    cb = js.NewCallback(func(args []js.Value) {
        fmt.Println("button clicked")
        //cb.Release() // release the callback if the button will not be clicked again
    })
    js.Global().Get("document").Call("getElementById", "myButton").Call("addEventListener", "click", cb)
//    js.Global().Set("Sorted", "abcdefg")
*/
    initSortTable()
    st := func(i []js.Value) { SortTable(i[0].String()) }
    js.Global().Set("SortTable", js.NewCallback(st))

    initPickDate()
    pd := func(i []js.Value) { PickDate(i[0].String()) }
    js.Global().Set("PickDate", js.NewCallback(pd))

    initChart()
    js.Global().Set("DrawChart", js.NewCallback(
        func(i []js.Value) { DrawChart(i[0].String()) }))

    select {}
    fmt.Printf("bye goasm\n")
}


func SortTable(id string) {
    t := NewTable(id)
    t.Init()
}


func PickDate(id string) {
    d := NewDate(id)
    d.Init()
}


func DrawChart(id string) {
    c := NewChart(id)
    c.Init()
}
