package main

import (
    "syscall/js"
)


type Chart struct {
    svg     js.Value    // svg element of chart
}


func initChart() {
}


func NewChart(id string) (c *Chart) {
    c = new(Chart)
    c.svg = js.Global().Get("document").Call("getElementById", id)
    return
}


func (c *Chart)Init() {
}
