package main

import (
    "fmt"
    "time"
    "syscall/js"

    "fgdwcfgo/log"
)


type DatePicker struct {
    id   string
    elm  js.Value       // input js element
    btn  js.Value       // generated button
}

var (
    DpDiv js.Value
    Tbody js.Value
)


func initPickDate() {
    doc := js.Global().Get("document")
    st := doc.Call("createElement", "style")
    st.Set("innerHTML",
           `table.sorted th {
                background-color: #d0dddd;
            }
            input.DatePicker {
                #background-image: url(/css/images/bg.gif);
                #background-repeat: no-repeat;
                #background-position: center right;
                cursor: pointer;
                #background-color: #d0dddd;
                #background-position: center left;
            }
            `)
    //hd := doc.Get("head")
    //if hd.Type() == js.TypeNull {
    hd := doc.Call("getElementsByTagName", "head")
    if hd.Length() < 1 {
        log.Errorf("Can not find <head> in html ")
        return
    }
    hd.Index(0).Call("insertBefore", st, hd.Index(0).Get("firstChild"))

    DpDiv := doc.Call("createElement", "div")
    DpDiv.Get("classList").Call("add", "datepicker")
    DpDiv.Set("innerHTML", `
        <table class="year">
            <tr><td>lf</td><td>Mon</td><td>Yea</td><td>rt</td></tr>
        </table>
        <table class="week">
            <thead><tr>
                <th>Su</th><th>Mo</th><th>Tu</th><th>We</th><th>Th</th><th>Fr</th><th>Sa</th>
            </tr></thead>
            <tbody>
            <tr><td>1</td><td>6</td><td>8</td><td>0</td><td>1</td><td>2</td><td>2</td></tr>
            <tr><td>2</td><td>3</td><td>5</td><td>8</td><td>0</td><td>1</td><td>2</td></tr>
            <tr><td>2</td><td>4</td><td>6</td><td>8</td><td>9</td><td>1</td><td>2</td></tr>
            <tr><td>2</td><td>4</td><td>5</td><td>8</td><td>0</td><td>1</td><td>2</td></tr>
            <tr><td>2</td><td>4</td><td>6</td><td>7</td><td>9</td><td>1</td><td>3</td></tr>
            </tbody>
        </table>
           `)
    tbs := DpDiv.Call("getElementsByTagName", "tbody")
    for i := 0; i < tbs.Length(); i++ {
        trs := tbs.Index(i).Call("getElementsByTagName", "tr")
        if trs.Length() == 5 {
            Tbody = tbs.Index(i)
            break
        }
    }
    //bd := doc.Get("body") // dom3 not support yet
    //if bd.Type() == js.TypeNull {
    bd := doc.Call("getElementsByTagName", "body")
    if bd.Length() < 1 {
        log.Errorf("Can not find <body> in html ")
        return
    }
    bd.Index(0).Call("appendChild", DpDiv)
    update_table(time.Now())
}


func NewDate(id string) *DatePicker{
    return &DatePicker{id: id}
}


func (d *DatePicker)Init() {
    doc := js.Global().Get("document")
    d.elm = doc.Call("getElementById", d.id)
    if d.elm.Type() == js.TypeNull {
        log.Errorf("Can not find input with id=%s", d.id)
        return
    }
    //d.add_btn()
    d.btn = doc.Call("createElement", "button")
    d.btn.Set("innerText", "...")
    //inserAfter(d.elm, d.btn)
    d.elm.Get("parentNode").Call("insertBefore",
                                 d.btn, d.elm.Get("nextSibling"))
}


func update_table(n time.Time) {
    //log.Debugf("n = %v", n)
    dt := n.AddDate(0, 0, 1 - n.Day())   // 1st day of month
    //log.Debugf("dt = %v", dt)
    dt = dt.AddDate(0, 0, -int(dt.Weekday()))   // Sunday
    //log.Debugf("dt = %v", dt)
    trs := Tbody.Call("getElementsByTagName", "tr")
    for r := 0; r < 5; r++ {
        tds := trs.Index(r).Call("getElementsByTagName", "td")
        for c := 0; c < 7; c++ {
            //days[r][c] = d
            cs := tds.Index(c).Get("classList")
            if dt.Month() == n.Month() {
                cs.Call("remove", "other")
            } else {
                cs.Call("add", "other")
            }
            h := fmt.Sprintf(`<a href="#">%d</a>`, dt.Day())
            tds.Index(c).Set("innerHTML", h)
            dt = dt.AddDate(0, 0, 1)
        }
    }
}
