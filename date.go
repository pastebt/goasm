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
    DpDiv   js.Value    // date elem div
    Month   js.Value    // date elem div month select
    Syear   js.Value    // date elem div year select
    Sdays   js.Value    // date elem div day select

    DpAct   *DatePicker // active DatePicker
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
    DpDiv = doc.Call("createElement", "div")
    DpDiv.Get("classList").Call("add", "datepicker")
    DpDiv.Set("innerHTML", `
        <table class="head"><tr>
            <td class="arrow"><circle><arrow class="left" title="prev" ></arrow></circle></td>
            <td><select id="date_picker_sel_month">
            </select></td>
            <td><select id="date_picker_sele_year">
            </select></td>
            <td class="arrow"><circle><arrow class="right" title="next"></arrow></circle></td>
        </tr></table>
        <table class="week">
            <thead><tr>
                <th>Su</th><th>Mo</th><th>Tu</th><th>We</th><th>Th</th><th>Fr</th><th>Sa</th>
            </tr></thead>
            <tbody id="date_picker_sele_days">
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
    //bd := doc.Get("body") // dom3 not support yet
    //if bd.Type() == js.TypeNull {
    bd := doc.Call("getElementsByTagName", "body")
    if bd.Length() < 1 {
        log.Errorf("Can not find <body> in html ")
        return
    }
    bd.Index(0).Call("appendChild", DpDiv)
    Month = doc.Call("getElementById", "date_picker_sel_month")
    Syear = doc.Call("getElementById", "date_picker_sele_year")
    Sdays = doc.Call("getElementById", "date_picker_sele_days")
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
    d.elm.Call("addEventListener", "keypress", js.NewCallback(d.keypress))
    d.elm.Call("addEventListener", "change", js.NewCallback(d.change))
    //d.add_btn()
    d.btn = doc.Call("createElement", "button")
    d.btn.Set("innerText", "...")
    d.btn.Call("addEventListener", "click", js.NewCallback(d.click_btn))
    //inserAfter(d.elm, d.btn)
    d.elm.Get("parentNode").Call("insertBefore",
                                 d.btn, d.elm.Get("nextSibling"))
}


// update picker div,
func (d *DatePicker)keypress(vs []js.Value) {
    DpAct = d
    log.Debugf("keypress DpAct=%v, vs=%v", d, vs)
}

// hide date picker div
// update date
func (d *DatePicker)change(vs []js.Value) {
    DpAct = d
    log.Debugf("change DpAct=%v, vs=%v", d, vs)
}

func (d *DatePicker)click_btn(vs []js.Value) {
    st := DpDiv.Get("style")
    if DpAct == d {
        // switch display block/none
        if st.Get("display").String() == "block" {
            st.Set("display", "none")
        } else {
            st.Set("display", "block")
        }
    } else {
        DpAct = d
        // keep display block
        st.Set("display", "block")
    }
    top := d.elm.Get("offsetTop").Float() + d.elm.Get("offsetHeight").Float()
    lft := d.elm.Get("offsetLeft").Float()
    st.Set("top", top)
    st.Set("left", lft)
    log.Debugf("click_btn DpAct=%v, vs=%v, top=%v, left=%v", d, vs, top, lft)
}


func update_table(n time.Time) {
    //log.Debugf("n = %v", n)
    dt := n.AddDate(0, 0, 1 - n.Day())   // 1st day of month
    //log.Debugf("dt = %v", dt)
    dt = dt.AddDate(0, 0, -int(dt.Weekday()))   // Sunday
    //log.Debugf("dt = %v", dt)
    trs := Sdays.Call("getElementsByTagName", "tr")
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
    ms := ""
    for m := time.January; m <= time.December; m++ {
        s := ""
        if m == n.Month() { s = ` selected="selected"` }
        ms += fmt.Sprintf("<option value=%d%s>%s</option>\n",
                          m, s, m.String()[:3])
    }
    Month.Set("innerHTML", ms)

    ms = ""
    ty := n.Year()
    for y := ty - 10; y <= ty + 10; y++ {
        s := ""
        if y == ty { s = ` selected="selected"` }
        ms += fmt.Sprintf("<option value=%d%s>%d</option>\n", y, s, y)
    }
    Syear.Set("innerHTML", ms)
}
