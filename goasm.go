package main

import (
//    "os"
    "fmt"
    "syscall/js"
    "fgdwcfgo/log"
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
    st := func(i []js.Value) { SortTable(i[0].String()) }
    js.Global().Set("SortTable", js.NewCallback(st))
    select {}
    fmt.Printf("bye goasm\n")
}


func SortTable(id string) {
    //println("table id is", id)
    elm := js.Global().Get("document").Call("getElementById", id)
    //udf := js.Undefined()
    //udf := js.Null()
    //fmt.Printf("elm=%#v, udf=%#v, %#v\n", elm, udf, elm == udf)
    if elm == js.Null() {
        log.Errorf("Can not find table with id=%s\n", id)
        return
    }
    
}
