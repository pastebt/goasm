package main

import (
    "sort"
    "syscall/js"
    "fgdwcfgo/log"
)

func insertFirst(elm, chd js.Value) {
    /*
    fc := elm.Get("firstChild")
    if fc.Type == js.Null {
    } else {
        elm.Call("")
    }
    */
}

func initSortTable() {
    doc := js.Global().Get("document")
    ss := doc.Get("styleSheets")
    log.Debugf("ss.Length = %v", ss.Length())
    for i := 0; i < ss.Length(); i++ {
        log.Debugf("title %v = %v", i, ss.Index(i).Get("title"))
    }

    st := doc.Call("createElement", "style")
    st.Set("innerHTML",
           `table.sorted th.header {
                background-image: url(/css/images/bg.gif);
                background-repeat: no-repeat;
                #background-position: center right;
                cursor: pointer;
                background-color: red;
                background-position: center left;
            }
            table.sorted th.headerUp {
                background-image: url(/css/images/asc.gif);
                background-color: #8dbdd8;
            }
            table.sorted th.headerDn {
                background-image: url(/css/images/desc.gif);
                background-color: #8dbdd8;
            }
            `)
    hd := doc.Call("getElementsByTagName", "head")
    if hd.Length() < 1 {
        log.Errorf("Can not find <head> in html ")
        return
    }
    hd.Index(0).Call("insertBefore", st, hd.Index(0).Get("firstChild"))

    /*
    st := doc.Call("querySelector", "table.sorted")
    if st.Type() == js.TypeNull {
        log.Errorf("can not find class table.sorted")
    }
    hd := doc.Call("querySelector", "table.sorted thead tr .header")
    if hd.Type() == js.TypeNull {
        log.Errorf("can not find class header")
    }
    up := doc.Call("querySelector", "table.sorted thead tr .headerUp")
    if up.Type() == js.TypeNull {
        log.Errorf("can not find class headerUp")
    }
    dn := doc.Call("querySelector", "table.sorted thead tr .headerDn")
    if dn.Type() == js.TypeNull {
        log.Errorf("can not find class headerDn")
    } else {
        log.Debugf("dn=%#v, dn.Type=%v", dn, dn.Type())
    }
    */
}


type Row struct {
    tr  js.Value
    tds []js.Value
}


type Table struct {
    id   string
    elm  js.Value       // table js element
    body js.Value       // table body element, use to redraw table
    cols []js.Value     // table headers list
    rows []Row          // table body tr list

}

// c: column
// t: type
// o: order
func less(rows []Row, col, typ, ord int) func(i, j int) bool {
    return func(i, j int) bool {
    I := rows[i].tds[col].Get("innerText")
    J := rows[j].tds[col].Get("innerText")
    // TODO, typ and ord
    return I.String() < J.String()
}
}

func NewTable(id string) *Table {
    return &Table{id: id,
                  cols: make([]js.Value, 0, 5),
                  rows: make([]Row, 0, 10),
                  }
}


func (t *Table)Init() {
    t.elm = js.Global().Get("document").Call("getElementById", t.id)
    //if t.elm == js.Null() {
    if t.elm.Type() == js.TypeNull {
        log.Errorf("Can not find table with id=%s", t.id)
        return
    }
    // contains(class)
    t.elm.Get("classList").Call("add", "sorted")
    t.get_head()
    t.get_body()
}


// return func will copy col_id to local, which is necessory
func (t *Table)do_sort(col_id int) func([]js.Value) {
    return func([]js.Value) {
    sort.Slice(t.rows, less(t.rows, col_id, 0, 0))
    log.Debugf("do_sort table %s, column %d", t.id, col_id)
    for _, row := range t.rows {
        for _, td := range row.tds {
            s := td.Get("innerText").String()
            log.Debugf("s = %#v", s)
        }
        t.body.Call("removeChild", row.tr)
        t.body.Call("appendChild", row.tr)
    }

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
    tbd := t.elm.Call("getElementsByTagName", "tbody")
    if tbd.Length() < 1 {
        log.Errorf("Can not find tbody of table %s", t.id)
        return
    }
    t.body = tbd.Index(0)
    trs := t.body.Call("getElementsByTagName", "tr")
    for i := 0; i < trs.Length(); i++ {
        t.rows = append(t.rows, t.get_row(trs.Index(i)))
    }
}


func (t *Table)get_row(tr js.Value) Row {
    row := Row{tr: tr, tds: make([]js.Value, 0, 5)}
    tds := tr.Call("getElementsByTagName", "td")
    for i := 0; i < tds.Length(); i++ {
        row.tds = append(row.tds, tds.Index(i))
    }
    return row
}
