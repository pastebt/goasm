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


type DateDiv struct {
    div     js.Value    // date elem div
    sel     time.Time   // div current time
    mon     js.Value    // date elem div month select
    year    js.Value    // date elem div year select
    days    js.Value    // date elem div day select
    mcb     js.Callback // call back for month change
    scb     js.Callback // call back for month/year select change
    dcb     js.Callback // call back for days select
    act     *DatePicker // active DatePicker
    fmt     string      // date format
}


var DD = DateDiv{fmt: "2006-01-02"}


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
            div.datepicker table td.sel {
                border: 1px solid #2694e8;
                background: #3baae3;
            }
            div.datepicker table td.now {
                border: 1px solid #f9dd34;
                background: #ffef8f;
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
                background: none;
                text-align: center;
            }
            div.datepicker table.head td.arrow:hover {
                border: 1px solid #aed0ea;
                border-radius: 5px;
            }
            div.datepicker table.head td.arrow {
                border: 1px solid transparent;
            }
            circle {
                width: 16px;
                height: 16px;
                background: #d7ebf9;
                border-radius: 50%;
                display: inline-block;
            }
            arrow.left {
                border-top: 6px solid transparent;
                border-right: 8px solid black;
                border-bottom: 6px solid transparent;
                display: inline-block;
                margin: 2px 5px 0 0;
            }
            arrow.right {
                border-top: 6px solid transparent;
                border-left: 8px solid black;
                border-bottom: 6px solid transparent;
                display: inline-block;
                margin: 2px 0 0 5px;
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
/*
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

                <option value=2017>2017</option>
                <option value=2018>2018</option>
                <option value=2019>2019</option>
*/
    DD.div = doc.Call("createElement", "div")
    DD.div.Get("classList").Call("add", "datepicker")
    td := `<td onclick="PickDateClickDay(this);"></td>`
    DD.div.Set("innerHTML", `
        <table class="head"><tr>
            <td class="arrow" onclick="PickDateClickLR(-1);">
                <circle><arrow class="left" title="prev" ></arrow></circle></td>
            <td><select id="date_picker_sel_month" onchange="PickDateSelChange(this);">
            </select></td>
            <td><select id="date_picker_sele_year" onchange="PickDateSelChange(this);">
            </select></td>
            <td class="arrow" onclick="PickDateClickLR(1);">
                <circle><arrow class="right" title="next"></arrow></circle></td>
        </tr></table>
        <table class="week">
            <thead><tr>
                <th>Su</th><th>Mo</th><th>Tu</th><th>We</th><th>Th</th><th>Fr</th><th>Sa</th>
            </tr></thead>
            <tbody id="date_picker_sele_days">
            <tr>` + td + td + td + td + td + td + td + `</tr>
            <tr>` + td + td + td + td + td + td + td + `</tr>
            <tr>` + td + td + td + td + td + td + td + `</tr>
            <tr>` + td + td + td + td + td + td + td + `</tr>
            <tr>` + td + td + td + td + td + td + td + `</tr>
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
    //bd := doc.Get("body") // dom3 not support yet
    //if bd.Type() == js.TypeNull {
    bd := doc.Call("getElementsByTagName", "body")
    if bd.Length() < 1 {
        log.Errorf("Can not find <body> in html ")
        return
    }
    bd.Index(0).Call("appendChild", DD.div)
    DD.mon = doc.Call("getElementById", "date_picker_sel_month")
    DD.year = doc.Call("getElementById", "date_picker_sele_year")
    DD.days = doc.Call("getElementById", "date_picker_sele_days")
    DD.sel = time.Now()
    update_table()
}


func NewDate(id string) *DatePicker{
    return &DatePicker{id: id}
}


func (d *DatePicker)Init() {
    //cb := js.Global().Get("PickDateClickDay")
    //if cb.Type() == js.TypeFunction {
    //    log.Debugf("release PickDateClickDay")
    //    
    //}
    DD.scb.Release()
    DD.dcb.Release()
    DD.mcb.Release()
    DD.scb = js.NewCallback(d.mon_year_chg)
    DD.dcb = js.NewCallback(d.click_day)
    DD.mcb = js.NewCallback(d.click_lr)
    js.Global().Set("PickDateSelChange", DD.scb)
    js.Global().Set("PickDateClickDay", DD.dcb)
    js.Global().Set("PickDateClickLR", DD.mcb)

    doc := js.Global().Get("document")
    d.elm = doc.Call("getElementById", d.id)
    if d.elm.Type() == js.TypeNull {
        log.Errorf("Can not find input with id=%s", d.id)
        return
    }
    d.elm.Call("addEventListener", "keypress",
               js.NewCallback(d.input_keypress))
    d.elm.Call("addEventListener", "change",
               js.NewCallback(d.input_change))
    //d.add_btn()
    d.btn = doc.Call("createElement", "button")
    d.btn.Set("innerText", "...")
    d.btn.Call("addEventListener", "click", js.NewCallback(d.click_btn))
    //inserAfter(d.elm, d.btn)
    d.elm.Get("parentNode").Call("insertBefore",
                                 d.btn, d.elm.Get("nextSibling"))
}


// update picker div,
func (d *DatePicker)input_keypress(vs []js.Value) {
    DD.act = d
    log.Debugf("keypress DpAct=%v, vs=%v", d, vs)
}

// hide date picker div
// update date
func (d *DatePicker)input_change(vs []js.Value) {
    DD.act = d
    log.Debugf("change DpAct=%v, vs=%v", d, vs)
}


func (d *DatePicker)click_btn(vs []js.Value) {
    st := DD.div.Get("style")
    if DD.act == d {
        // switch display block/none
        if st.Get("display").String() == "block" {
            st.Set("display", "none")
        } else {
            st.Set("display", "block")
        }
    } else {
        DD.act = d
        // keep display block
        st.Set("display", "block")
    }
    top := d.elm.Get("offsetTop").Float() + d.elm.Get("offsetHeight").Float()
    lft := d.elm.Get("offsetLeft").Float()
    st.Set("top", top)
    st.Set("left", lft)
    log.Debugf("click_btn DpAct=%v, vs=%v, top=%v, left=%v", d, vs, top, lft)
}


func (d *DatePicker)mon_year_chg(vs []js.Value) {
    v := vs[0].Get("value").String()
    log.Debugf("mon_year_chg, vs=%v, text=%s", vs, v)
    var y, m int
    if t, e := time.Parse("1", v); e == nil {
        // month change
        m = int(t.Month() - DD.sel.Month())
    } else if t, e := time.Parse("2006", v); e == nil{
        y = t.Year() - DD.sel.Year()
    }
    if y != 0 || m != 0 {
        DD.sel = DD.sel.AddDate(y, m, 0)
        update_table()
        d.elm.Set("value", DD.sel.Format(DD.fmt))
    }
}


func (d *DatePicker)click_day(vs []js.Value) {
    v := vs[0].Get("innerText").String()
    log.Debugf("click_day, vs=%v, text=%s", vs, v)
    t, e := time.Parse("2", v)
    if e != nil {
        log.Errorf("time.Parse %v err=%v", v, e)
        return
    }
    v = vs[0].Call("getAttribute", "m").String()
    m, e := time.Parse("1", v)
    if e != nil {
        log.Errorf("time.Parse %v err=%v", v, e)
        return
    }
    dt := t.Day() - DD.sel.Day()
    mt := int(m.Month() - DD.sel.Month())
    if dt != 0 || mt != 0{
        DD.sel = DD.sel.AddDate(0, mt, dt)
        update_table()
        d.elm.Set("value", DD.sel.Format(DD.fmt))
    }
}


func (d *DatePicker)click_lr(vs []js.Value) {
    log.Debugf("click_day, vs=%v", vs)
    m := vs[0].Int()
    DD.sel = DD.sel.AddDate(0, m, 0)
    update_table()
    d.elm.Set("value", DD.sel.Format(DD.fmt))
}


func update_table() {
    //log.Debugf("n = %v", n)
    now := time.Now().Format(DD.fmt)
    n := DD.sel
    dt := n.AddDate(0, 0, 1 - n.Day())   // 1st day of month
    //log.Debugf("dt = %v", dt)
    dt = dt.AddDate(0, 0, -int(dt.Weekday()))   // Sunday
    //log.Debugf("dt = %v", dt)
    trs := DD.days.Call("getElementsByTagName", "tr")
    for r := 0; r < 5; r++ {
        tds := trs.Index(r).Call("getElementsByTagName", "td")
        for c := 0; c < 7; c++ {
            cs := tds.Index(c).Get("classList")
            if dt.Month() == n.Month() {
                cs.Call("remove", "other")
                if dt.Day() == n.Day() {
                   cs.Call("add", "sel")
                } else {
                    cs.Call("remove", "sel")
                }
            } else {
                cs.Call("add", "other")
                cs.Call("remove", "sel")
            }
            if dt.Format(DD.fmt) == now {
                cs.Call("remove", "sel")
                cs.Call("add", "now")
            } else {
                cs.Call("remove", "now")
            }
            //h := fmt.Sprintf(`<a href="#">%d</a>`, dt.Day())
            h := fmt.Sprintf(`%d`, dt.Day())
            td := tds.Index(c)
            td.Set("innerHTML", h)
            td.Call("setAttribute", "m", fmt.Sprintf("%d", dt.Month()))
            //td.Call("setAttribute", "onclick", "PickDateClickDay(this);")
            dt = dt.AddDate(0, 0, 1)
        }
    }
    ms := ""
    for m := time.January; m <= time.December; m++ {
        s := ""
        if m == n.Month() { s = ` selected="selected"` }
        ms += fmt.Sprintf("<option value=%d%s>%s</option>\n",
                          m, s, m.String()[:3])
    }
    DD.mon.Set("innerHTML", ms)

    ms = ""
    ty := n.Year()
    for y := ty - 10; y <= ty + 10; y++ {
        s := ""
        if y == ty { s = ` selected="selected"` }
        ms += fmt.Sprintf("<option value=%d%s>%d</option>\n", y, s, y)
    }
    DD.year.Set("innerHTML", ms)
}
