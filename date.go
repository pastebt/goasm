package main

import (
    "syscall/js"

    "fgdwcfgo/log"
)


type DatePicker struct {
    id   string
    elm  js.Value       // input js element
    btn  js.Value       // generated button
}


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
    hd := doc.Call("getElementsByTagName", "head")
    if hd.Length() < 1 {
        log.Errorf("Can not find <head> in html ")
        return
    }
    hd.Index(0).Call("insertBefore", st, hd.Index(0).Get("firstChild"))

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

