package main

import (
    "sort"
    "strings"
    "strconv"
    "net/http"
    "io/ioutil"
    "syscall/js"

    "fgdwcfgo/log"
)


func initSortTable() {
    doc := js.Global().Get("document")
    st := doc.Call("createElement", "style")
    st.Set("innerHTML",
           `table.sorted th {
                background-color: #d0dddd;
            }
            table.sorted th.header {
                #background-image: url(/css/images/bg.gif);
                background-image: url(data:image/gif;base64,R0lGODlhFQAJAIAAACMtMP///yH5BAEAAAEALAAAAAAVAAkAAAIXjI+AywnaYnhUMoqt3gZXPmVg94yJVQAAOw==);
                background-repeat: no-repeat;
                background-position: center right;
                cursor: pointer;
                #background-color: #d0dddd;
                #background-position: center left;
            }
            table.sorted th.headerUp {
                #background-image: url(/css/images/asc.gif);
                background-image: url(data:image/gif;base64,R0lGODlhFQAEAIAAACMtMP///yH5BAEAAAEALAAAAAAVAAQAAAINjB+gC+jP2ptn0WskLQA7);
                background-color: #8dbdd8;
            }
            table.sorted th.headerDn {
                #background-image: url(/css/images/desc.gif);
                background-image: url(data:image/gif;base64,R0lGODlhFQAEAIAAACMtMP///yH5BAEAAAEALAAAAAAVAAQAAAINjI8Bya2wnINUMopZAQA7);
                background-color: #8dbdd8;
            }
            table.sorted tbody tr:nth-child(even) {
                background-color: #F0F0F6;
            }
            table.sorted tbody tr:hover {
                background: #e6eeee;
                color: red;
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
    body js.Value       // table body element, use to get data / re-draw table
    cols []js.Value     // table headers list
    rows []Row          // table body tr list

}

var sort_type = map[string]int{"string": 0,     // default, fixed to 0
                               "number": 1,
                               "currency": 2}


func less_num(Is, Js string) bool {
    Ii, Ie := strconv.ParseFloat(Is, 64)
    Ji, Je := strconv.ParseFloat(Js, 64)
    if Ie != nil || Je != nil {
        log.Errorf("less Is=%v, Js=%v, Ie=%v, Je=%v", Is, Js, Ie, Je)
        return Is < Js
    }
    return Ii < Ji
}


// c: column
// t: type: number, currency, string(default)
// o: order
func less(rows []Row, col, typ, ord int) func(i, j int) bool {
    return func(i, j int) bool {
    I := rows[i].tds[col].Get("innerText")
    J := rows[j].tds[col].Get("innerText")
    // TODO, typ and ord
    b, Is, Js := false, I.String(), J.String()
    switch typ {
    case 1:
        b = less_num(Is, Js)
    case 2:
        b = less_num(strings.Replace(Is, ",", "", -1),
                     strings.Replace(Js, ",", "", -1))
    default:
        b = Is < Js
    }
    if ord < 0 { return !b }
    return b
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
    isc := t.get_head()
    log.Debugf("isc=%v", isc)
    t.get_body()
    if isc >= 0 { t.do_sort(isc) }
}


// visit url, get data, generate tbody innerHTML
func (t *Table)fetch_data(url string) (string, error) {
    log.Debugf("start fetch_data %v", url)
    resp, err := http.Get(url)
    log.Debugf("http.Get(%v) return %v, %v", url, resp, err)
    if err != nil { return "", err }
    defer resp.Body.Close()
    dat, err := ioutil.ReadAll(resp.Body)
    if err != nil { return "", err }
    return string(dat), nil
}


// return func will copy col_id to local, which is necessory
func (t *Table)click_col(col_id int) func([]js.Value) {
    return func([]js.Value) {
    cls := t.cols[col_id].Get("classList")
    if cls.Call("contains", "headerUp").Bool() {
        log.Debugf("remove headerUp, add headerDn")
        cls.Call("remove", "headerUp")
        cls.Call("add", "headerDn")
    } else {
        log.Debugf("remove headerDn, add headerUp")
        cls.Call("remove", "headerDn")
        cls.Call("add", "headerUp")
    }
    for i, col := range t.cols {
        c := col.Get("classList")
        c.Call("add", "header")
        if i != col_id {
            c.Call("remove", "headerUp", "headerDn")
        }
    }
    t.do_sort(col_id)
    }
}


func (t *Table)do_sort(col_id int) {
    s := t.cols[col_id].Call("getAttribute", "stype").String()
    typ := sort_type[strings.ToLower(s)]

    ord := 1
    if t.cols[col_id].Get("classList").Call("contains", "headerUp").Bool() {
        ord = -1
    }

    sort.Slice(t.rows, less(t.rows, col_id, typ, ord))
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


// return init sort column
func (t *Table)get_head() (isc int) {
    isc = -1
    thd := t.elm.Call("getElementsByTagName", "thead")
    if thd.Length() < 1 {
        log.Errorf("Can not find thead of table %s", t.id)
        return
    }
    ths := thd.Index(0).Call("getElementsByTagName", "th")
    log.Debugf("found %d th(s) in table %s", ths.Length(), t.id)
    for i := 0; i < ths.Length(); i++ {
        col := ths.Index(i)
        cls := col.Get("classList")
        typ := col.Call("getAttribute", "stype")
        up := cls.Call("contains", "headerUp").Bool()
        dn := cls.Call("contains", "headerDn").Bool()
        if (typ.Type() == js.TypeNull) && !up && !dn &&
           !cls.Call("contains", "header").Bool() {
            continue
        }
        cls.Call("add", "header")
        cb := js.NewCallback(t.click_col(i))
        ths.Index(i).Call("addEventListener", "click", cb)
        t.cols = append(t.cols, col)
        if up || dn { isc = i }
    }
    return
}


func (t *Table)get_body() {
    tbd := t.elm.Call("getElementsByTagName", "tbody")
    if tbd.Length() < 1 {
        log.Errorf("Can not find tbody of table %s", t.id)
        return
    }
    t.body = tbd.Index(0)
    src := t.body.Call("getAttribute", "src").String()
    // var x = document.URL;
    lo := js.Global().Get("location")
    if src[0] == '/' {
        src = lo.Get("origin").String() + src
    }
    log.Infof("tbody src=%s, lo=%v", src, lo)
    ih, err := t.fetch_data(src)
    log.Infof("ih=%v, err=%v", ih, err)
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
