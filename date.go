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
    DpDiv   js.Value
    Syear   js.Value
    Tbody   js.Value
)


func initPickDate() {
    doc := js.Global().Get("document")
    st := doc.Call("createElement", "style")
    st.Set("innerHTML", `
            input.DatePicker {
            }
            div.datepicker {
                position: absolute;
                top: 83.8px;
                left: 219.838px;
                z-index: 1;
                display: block;
                width: auto;
                font-size: 9px;
                padding: .2em;
                border-radius: 6px;
                border: 1px solid #dddddd;
                background-color: #f2f5f7;
                color: #362b36;
            }
            div.datepicker table {
                width: 100%;
                font-size: 1em;
                font-weight: bold;
                display: table;
                border-spacing: 2px;
            }
            div.datepicker table th {
                padding: .7em .3em;
                text-align: center;
                border: 0;
            }
            div.datepicker table td {
                text-align: right;
                cursor: pointer;
                padding: 0 .4em ;
                border: 1px solid #aed0ea;
                background: #d7ebf9;
                color: #2779aa;
            }
            div.datepicker table td.other {
                opacity: .7;
                font-weight: normal;
            }
            div.datepicker table.weekend {

            }
            div.datepicker table.head {
                border: 1px solid #aed0ea;
                border-radius: 6px;
            }
            div.datepicker table.head select {
                font-size: 1em;
                width: 100%;
                padding: 0;
                font-weight: bold;
            }
            div.datepicker table.head td {
                border: 0;
                padding: 0;
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
        <table class="head"><tr>
            <td>lf</td>
            <td><select>
                <option value=1>Jan</option>
                <option value=2>Feb</option>
                <option value=3>Mar</option>
                <option value=4>Apr</option>
                <option value=5>May</option>
                <option value=6>Jun</option>
                <option value=7>Jul</option>
                <option value=8>Aug</option>
                <option value=9>Sep</option>
                <option value=10>Oct</option>
                <option value=11>Nov</option>
                <option value=12>Dec</option>
            </select></td>
            <td><select id="date_picker_sele_year">
                <option value=2017>2017</option>
                <option value=2018>2018</option>
                <option value=2019>2019</option>
            </select></td>
            <td>rt</td>
        </tr></table>
        <table class="week">
            <thead><tr>
                <th>Su</th><th>Mo</th><th>Tu</th><th>We</th><th>Th</th><th>Fr</th><th>Sa</th>
            </tr></thead>
            <tbody id="date_picker_dis_month">
            <tr><td>1</td><td>6</td><td>8</td><td>0</td><td>1</td><td>2</td><td>2</td></tr>
            <tr><td>2</td><td>3</td><td>5</td><td>8</td><td>0</td><td>1</td><td>2</td></tr>
            <tr><td>2</td><td>4</td><td>6</td><td>8</td><td>9</td><td>1</td><td>2</td></tr>
            <tr><td>2</td><td>4</td><td>5</td><td>8</td><td>0</td><td>1</td><td>2</td></tr>
            <tr><td>2</td><td>4</td><td>6</td><td>7</td><td>9</td><td>1</td><td>3</td></tr>
            </tbody>
        </table>
           `)
    /*
    tbs := DpDiv.Call("getElementsByTagName", "tbody")
    for i := 0; i < tbs.Length(); i++ {
        trs := tbs.Index(i).Call("getElementsByTagName", "tr")
        if trs.Length() == 5 {
            Tbody = tbs.Index(i)
            break
        }
    }
    */
    //Syear = doc.Call("getElementById", "date_picker_sele_year")
    //Tbody = doc.Call("getElementById", "date_picker_dis_month")
    //bd := doc.Get("body") // dom3 not support yet
    //if bd.Type() == js.TypeNull {
    bd := doc.Call("getElementsByTagName", "body")
    if bd.Length() < 1 {
        log.Errorf("Can not find <body> in html ")
        return
    }
    bd.Index(0).Call("appendChild", DpDiv)
    Syear = doc.Call("getElementById", "date_picker_sele_year")
    Tbody = doc.Call("getElementById", "date_picker_dis_month")
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
            //h := fmt.Sprintf(`<a href="#">%d</a>`, dt.Day())
            h := fmt.Sprintf(`%d`, dt.Day())
            tds.Index(c).Set("innerHTML", h)
            dt = dt.AddDate(0, 0, 1)
        }
    }
}
