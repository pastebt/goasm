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


// return func will copy col_id to local, which is necessory
func do_sort(tb_id string, col_id int) func([]js.Value) {
    return func([]js.Value) {
    log.Debugf("do_sort table %s, column %d", tb_id, col_id)
    }
}


func SortTable(id string) {
    //println("table id is", id)
    elm := js.Global().Get("document").Call("getElementById", id)
    //udf := js.Undefined()
    //udf := js.Null()
    //fmt.Printf("elm=%#v, udf=%#v, %#v\n", elm, udf, elm == udf)
    if elm == js.Null() {
        log.Errorf("Can not find table with id=%s", id)
        return
    }
    thd := elm.Call("getElementsByTagName", "thead")
    if thd.Length() < 1 {
        log.Errorf("Can not find thead of table %s", id)
        return
    }
    ths := thd.Index(0).Call("getElementsByTagName", "th")
    log.Debugf("found %d th(s) in table %s", ths.Length(), id)
    for i := 0; i < ths.Length(); i++ {
        cb := js.NewCallback(do_sort(id, i))
        ths.Index(i).Call("addEventListener", "click", cb)
    }
}
