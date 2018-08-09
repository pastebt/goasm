package main

import (
    "syscall/js"
    "fgdwcfgo/log"
)


type Row []js.Value


type Table struct {
    id   string
    elm  js.Value       // table js element
    cols []js.Value     // table headers list
    rows []Row          // table body tr list
}


func NewTable(id string) *Table {
    return &Table{id: id,
                  cols: make([]js.Value, 0, 5),
                  }
}


func (t *Table)Init() {
    t.elm = js.Global().Get("document").Call("getElementById", t.id)
    //if t.elm == js.Null() {
    if t.elm.Type() == js.TypeNull {
        log.Errorf("Can not find table with id=%s", t.id)
        return
    }
    t.get_head()
    t.get_body()
}


// return func will copy col_id to local, which is necessory
func (t *Table)do_sort(col_id int) func([]js.Value) {
    return func([]js.Value) {
    log.Debugf("do_sort table %s, column %d", t.id, col_id)
    }
}


func (t *Table)get_head() {
    thd := t.elm.Call("getElementsByTagName", "thead")
    if thd.Length() < 1 {
        log.Errorf("Can not find thead of table %s", t.id)
        return
    }
    ths := thd.Index(0).Call("getElementsByTagName", "th")
    log.Debugf("found %d th(s) in table %s", ths.Length(), t.id)
    for i := 0; i < ths.Length(); i++ {
        cb := js.NewCallback(t.do_sort(i))
        ths.Index(i).Call("addEventListener", "click", cb)
        t.cols = append(t.cols, ths.Index(i))
    }
}


func (t *Table)get_body() {
    thd := t.elm.Call("getElementsByTagName", "tbody")
    if thd.Length() < 1 {
        log.Errorf("Can not find tbody of table %s", t.id)
        return
    }

}
