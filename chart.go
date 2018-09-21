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

    title := `
<text x="400" y="20" fill="#3E576F" style="color:#3E576F;font-size:16px;font-family:'Lucida Grande', 'Lucida Sans Unicode', Verdana, Arial, Helvetica, sans-serif;" text-anchor="middle" transform="rotate(0 400 20)" class="highcharts-title"><tspan x="400">Daily UR report</tspan></text>
`

    prn := `
<rect x="0.5" y="0.5" width="23" height="19" rx="3" ry="3" fill="url(http://172.16.90.60:8080/#highcharts-3)" transform="translate(740,10)" stroke-width="1" zIndex="19" stroke="#B0B0B0"></rect>
<path d="M 745.5 23.5 L 757.5 23.5 757.5 18.5 745.5 18.5 Z M 748.5 18.5 L 748.5 14.5 754.5 14.5 754.5 18.5 Z M 748.5 23.5 L 747 26.5 756 26.5 754.5 23.5 Z" fill="#B5C9DF" stroke="#A0A0A0" stroke-width="1" zIndex="20"></path>
<rect x="740" y="10" width="24" height="20" fill="rgb(255,255,255)" fill-opacity="0.001" title="Print the chart" zIndex="21" style="cursor:pointer;"></rect>
`
    dwn := `
<rect x="0.5" y="0.5" width="23" height="19" rx="3" ry="3" fill="url(http://172.16.90.45:8000/unrated#highcharts-2)" transform="translate(766,10)" stroke-width="1" zIndex="19" stroke="#B0B0B0"></rect>
<path d="M 771.5 26.5 L 783.5 26.5 783.5 23.5 771.5 23.5 Z M 777.5 23.5 L 774.5 18.5 776.5 18.5 776.5 14.5 778.5 14.5 778.5 18.5 780.5 18.5 Z" fill="#A8BF77" stroke="#A0A0A0" stroke-width="1" zIndex="20"></path>
<rect x="766" y="10" width="24" height="20" fill="rgb(255,255,255)" fill-opacity="0.001" title="Export to raster or vector image" zIndex="21" style="cursor:pointer;"></rect>
`
    path := `
<path d="M 80 310.5 L 750 310.5" fill="none" stroke="#C0C0C0" stroke-width="1"></path>
<path d="M 80 249.5 L 750 249.5" fill="none" stroke="#C0C0C0" stroke-width="1"></path>
<path d="M 80 188.5 L 750 188.5" fill="none" stroke="#C0C0C0" stroke-width="1"></path>
<path d="M 80 127.5 L 750 127.5" fill="none" stroke="#C0C0C0" stroke-width="1"></path>
<path d="M 80 66.5 L 750 66.5" fill="none" stroke="#C0C0C0" stroke-width="1"></path>
`

    c.svg.Set("innerHTML", title + prn + dwn + path)
/*
    inner := `
<defs>
<clipPath id="highcharts-1">
<rect x="0" y="0" width="670" height="260" fill="none"></rect>
</clipPath>
<linearGradient id="highcharts-2" gradientUnits="userSpaceOnUse" x1="0" y1="0" x2="0" y2="20">
<stop offset="0.4" stop-color="#F7F7F7" stop-opacity="1"></stop>
<stop offset="0.6" stop-color="#E3E3E3" stop-opacity="1"></stop>
</linearGradient>
<linearGradient id="highcharts-3" gradientUnits="userSpaceOnUse" x1="0" y1="0" x2="0" y2="20">
<stop offset="0.4" stop-color="#F7F7F7" stop-opacity="1"></stop>
<stop offset="0.6" stop-color="#E3E3E3" stop-opacity="1"></stop>
</linearGradient>
</defs>
<rect x="0" y="0" width="800" height="400" rx="5" ry="5" fill="#FFFFFF" stroke="#4572A7" stroke-width="0"></rect>
<text x="400" y="20" fill="#3E576F" style="color:#3E576F;font-size:16px;font-family:'Lucida Grande', 'Lucida Sans Unicode', Verdana, Arial, Helvetica, sans-serif;" text-anchor="middle" transform="rotate(0 400 20)" class="highcharts-title">
<tspan x="400">Daily NOD report</tspan>
</text>
<g class="highcharts-grid" zIndex="1"></g><g class="highcharts-grid" zIndex="1">
<path d="M 80 310.5 L 750 310.5" fill="none" stroke="#C0C0C0" stroke-width="1"></path>
<path d="M 80 249.5 L 750 249.5" fill="none" stroke="#C0C0C0" stroke-width="1"></path>
<path d="M 80 188.5 L 750 188.5" fill="none" stroke="#C0C0C0" stroke-width="1"></path>
<path d="M 80 127.5 L 750 127.5" fill="none" stroke="#C0C0C0" stroke-width="1"></path>
<path d="M 80 66.5 L 750 66.5" fill="none" stroke="#C0C0C0" stroke-width="1"></path>
</g>
<rect x="0.5" y="0.5" width="23" height="19" rx="3" ry="3" fill="url(http://172.16.90.60:8080/#highcharts-2)" transform="translate(766,10)" stroke-width="1" zIndex="19" stroke="#B0B0B0"></rect>
<rect x="0.5" y="0.5" width="23" height="19" rx="3" ry="3" fill="url(http://172.16.90.60:8080/#highcharts-3)" transform="translate(740,10)" stroke-width="1" zIndex="19" stroke="#B0B0B0"></rect>
<path d="M 771.5 26.5 L 783.5 26.5 783.5 23.5 771.5 23.5 Z M 777.5 23.5 L 774.5 18.5 776.5 18.5 776.5 14.5 778.5 14.5 778.5 18.5 780.5 18.5 Z" fill="#A8BF77" stroke="#A0A0A0" stroke-width="1" zIndex="20"></path>
<path d="M 745.5 23.5 L 757.5 23.5 757.5 18.5 745.5 18.5 Z M 748.5 18.5 L 748.5 14.5 754.5 14.5 754.5 18.5 Z M 748.5 23.5 L 747 26.5 756 26.5 754.5 23.5 Z" fill="#B5C9DF" stroke="#A0A0A0" stroke-width="1" zIndex="20"></path>
<rect x="766" y="10" width="24" height="20" fill="rgb(255,255,255)" fill-opacity="0.001" title="Export to raster or vector image" zIndex="21" style="cursor:pointer;"></rect>
<rect x="740" y="10" width="24" height="20" fill="rgb(255,255,255)" fill-opacity="0.001" title="Print the chart" zIndex="21" style="cursor:pointer;"></rect>
<g class="highcharts-series" clip-path="url(http://172.16.90.60:8080/#highcharts-1)" visibility="visible" zIndex="3" transform="translate(80,50)">
<path d="M 6.568627450980392 128.0219170476456 L 225.52287581699346 111.65145745501502 444.47712418300654 22.73385564759488 663.4313725490196 12.38095238095238" fill="none" stroke="rgb(0, 0, 0)" stroke-width="5" isShadow="true" stroke-opacity="0.05" transform="translate(1,1)"></path>
<path d="M 6.568627450980392 128.0219170476456 L 225.52287581699346 111.65145745501502 444.47712418300654 22.73385564759488 663.4313725490196 12.38095238095238" fill="none" stroke="rgb(0, 0, 0)" stroke-width="3" isShadow="true" stroke-opacity="0.1" transform="translate(1,1)"></path>
<path d="M 6.568627450980392 128.0219170476456 L 225.52287581699346 111.65145745501502 444.47712418300654 22.73385564759488 663.4313725490196 12.38095238095238" fill="none" stroke="rgb(0, 0, 0)" stroke-width="1" isShadow="true" stroke-opacity="0.15000000000000002" transform="translate(1,1)"></path>
<path d="M 6.568627450980392 128.0219170476456 L 225.52287581699346 111.65145745501502 444.47712418300654 22.73385564759488 663.4313725490196 12.38095238095238" fill="none" stroke="#4572A7" stroke-width="2"></path><circle cx="663.4313725490196" cy="12.38095238095238" r="4" stroke="#FFFFFF" stroke-width="0" fill="#4572A7"></circle>
<circle cx="444.47712418300654" cy="22.73385564759488" r="4" stroke="#FFFFFF" stroke-width="0" fill="#4572A7"></circle>
<circle cx="225.52287581699346" cy="111.65145745501502" r="4" stroke="#FFFFFF" stroke-width="0" fill="#4572A7"></circle>
<circle cx="6.568627450980392" cy="128.0219170476456" r="4" stroke="#FFFFFF" stroke-width="0" fill="#4572A7"></circle>
</g>
<g class="highcharts-series" clip-path="url(http://172.16.90.60:8080/#highcharts-1)" visibility="visible" zIndex="3" transform="translate(80,50)">
<path d="M 6.568627450980392 259.89570155255734 L 225.52287581699346 259.847704726041 444.47712418300654 259.7976317009222 663.4313725490196 259.6373980205419" fill="none" stroke="rgb(0, 0, 0)" stroke-width="5" isShadow="true" stroke-opacity="0.05" transform="translate(1,1)"></path>
<path d="M 6.568627450980392 259.89570155255734 L 225.52287581699346 259.847704726041 444.47712418300654 259.7976317009222 663.4313725490196 259.6373980205419" fill="none" stroke="rgb(0, 0, 0)" stroke-width="3" isShadow="true" stroke-opacity="0.1" transform="translate(1,1)"></path>
<path d="M 6.568627450980392 259.89570155255734 L 225.52287581699346 259.847704726041 444.47712418300654 259.7976317009222 663.4313725490196 259.6373980205419" fill="none" stroke="rgb(0, 0, 0)" stroke-width="1" isShadow="true" stroke-opacity="0.15000000000000002" transform="translate(1,1)"></path>
<path d="M 6.568627450980392 259.89570155255734 L 225.52287581699346 259.847704726041 444.47712418300654 259.7976317009222 663.4313725490196 259.6373980205419" fill="none" stroke="#AA4643" stroke-width="2"></path>
<path d="M 663.4313725490196 255.63739802054192 L 667.4313725490196 259.6373980205419 663.4313725490196 263.6373980205419 659.4313725490196 259.6373980205419 Z" fill="#AA4643" stroke="#FFFFFF" stroke-width="0"></path>
<path d="M 444.47712418300654 255.79763170092218 L 448.47712418300654 259.7976317009222 444.47712418300654 263.7976317009222 440.47712418300654 259.7976317009222 Z" fill="#AA4643" stroke="#FFFFFF" stroke-width="0"></path>
<path d="M 225.52287581699346 255.847704726041 L 229.52287581699346 259.847704726041 225.52287581699346 263.847704726041 221.52287581699346 259.847704726041 Z" fill="#AA4643" stroke="#FFFFFF" stroke-width="0"></path>
<path d="M 6.568627450980392 255.89570155255734 L 10.568627450980392 259.89570155255734 6.568627450980392 263.89570155255734 2.568627450980392 259.89570155255734 Z" fill="#AA4643" stroke="#FFFFFF" stroke-width="0"></path>
</g>
<g class="highcharts-axis" zIndex="7">
<path d="M 87.5 310 L 87.5 315" fill="none" stroke="#C0D0E0" stroke-width="1"></path>
<text x="87" y="324" fill="#666" style="color:#666;font-size:11px;font-family:'Lucida Grande', 'Lucida Sans Unicode', Verdana, Arial, Helvetica, sans-serif;" text-anchor="middle" transform="rotate(0 87 324)">
<tspan x="87">Sep 1 00:00</tspan>
</text>
<path d="M 196.5 310 L 196.5 315" fill="none" stroke="#C0D0E0" stroke-width="1"></path>
<text x="196" y="324" fill="#666" style="color:#666;font-size:11px;font-family:'Lucida Grande', 'Lucida Sans Unicode', Verdana, Arial, Helvetica, sans-serif;" text-anchor="middle" transform="rotate(0 196 324)">
<tspan x="196">Sep 1 12:00</tspan>
</text>
<path d="M 306.5 310 L 306.5 315" fill="none" stroke="#C0D0E0" stroke-width="1"></path>
<text x="306" y="324" fill="#666" style="color:#666;font-size:11px;font-family:'Lucida Grande', 'Lucida Sans Unicode', Verdana, Arial, Helvetica, sans-serif;" text-anchor="middle" transform="rotate(0 306 324)">
<tspan x="306">Sep 2 00:00</tspan>
</text>
<path d="M 415.5 310 L 415.5 315" fill="none" stroke="#C0D0E0" stroke-width="1"></path>
<text x="415" y="324" fill="#666" style="color:#666;font-size:11px;font-family:'Lucida Grande', 'Lucida Sans Unicode', Verdana, Arial, Helvetica, sans-serif;" text-anchor="middle" transform="rotate(0 415 324)">
<tspan x="415">Sep 2 12:00</tspan>
</text>
<path d="M 524.5 310 L 524.5 315" fill="none" stroke="#C0D0E0" stroke-width="1"></path>
<text x="524" y="324" fill="#666" style="color:#666;font-size:11px;font-family:'Lucida Grande', 'Lucida Sans Unicode', Verdana, Arial, Helvetica, sans-serif;" text-anchor="middle" transform="rotate(0 524 324)">
<tspan x="524">Sep 3 00:00</tspan>
</text>
<path d="M 634.5 310 L 634.5 315" fill="none" stroke="#C0D0E0" stroke-width="1"></path>
<text x="634" y="324" fill="#666" style="color:#666;font-size:11px;font-family:'Lucida Grande', 'Lucida Sans Unicode', Verdana, Arial, Helvetica, sans-serif;" text-anchor="middle" transform="rotate(0 634 324)"><tspan x="634">Sep 3 12:00</tspan></text>
<path d="M 743.5 310 L 743.5 315" fill="none" stroke="#C0D0E0" stroke-width="1"></path>
<text x="743" y="324" fill="#666" style="color:#666;font-size:11px;font-family:'Lucida Grande', 'Lucida Sans Unicode', Verdana, Arial, Helvetica, sans-serif;" text-anchor="middle" transform="rotate(0 743 324)"><tspan x="743">Sep 4 00:00</tspan></text></g>
<path d="M 80 310.5 L 750 310.5" fill="none" stroke="#C0D0E0" stroke-width="1" zIndex="7"></path><g class="highcharts-axis" zIndex="7">
<text x="72" y="313" fill="#666" style="color:#666;font-size:11px;font-family:'Lucida Grande', 'Lucida Sans Unicode', Verdana, Arial, Helvetica, sans-serif;" text-anchor="end" transform="rotate(0 72 313)"><tspan x="72">0</tspan></text>
<text x="72" y="252" fill="#666" style="color:#666;font-size:11px;font-family:'Lucida Grande', 'Lucida Sans Unicode', Verdana, Arial, Helvetica, sans-serif;" text-anchor="end" transform="rotate(0 72 252)"><tspan x="72">500 K</tspan></text>
<text x="72" y="191" fill="#666" style="color:#666;font-size:11px;font-family:'Lucida Grande', 'Lucida Sans Unicode', Verdana, Arial, Helvetica, sans-serif;" text-anchor="end" transform="rotate(0 72 191)"><tspan x="72">1 M</tspan></text>
<text x="72" y="130" fill="#666" style="color:#666;font-size:11px;font-family:'Lucida Grande', 'Lucida Sans Unicode', Verdana, Arial, Helvetica, sans-serif;" text-anchor="end" transform="rotate(0 72 130)"><tspan x="72">1.5 M</tspan></text>
<text x="72" y="69" fill="#666" style="color:#666;font-size:11px;font-family:'Lucida Grande', 'Lucida Sans Unicode', Verdana, Arial, Helvetica, sans-serif;" text-anchor="end" transform="rotate(0 72 69)"><tspan x="72">2 M</tspan></text></g>
<text x="40" y="180" fill="#6D869F" style="color:#6D869F;font-weight:bold;font-family:'Lucida Grande', 'Lucida Sans Unicode', Verdana, Arial, Helvetica, sans-serif;font-size:12px;" text-anchor="middle" transform="rotate(270 40 180)" zIndex="7"><tspan x="40">URL Number</tspan></text><g class="highcharts-legend" zIndex="7" transform="translate(315.5,359)"><rect x="0.5" y="0.5" width="198" height="25" rx="5" ry="5" fill="none" stroke="#909090" stroke-width="1"></rect>
<text x="30" y="18" fill="#3E576F" style="cursor:pointer;color:#3E576F;fill:#3E576F;" zIndex="2"><tspan x="30">Collected</tspan></text>
<path d="M -21 0 L -5 0" fill="none" stroke-width="2" zIndex="2" stroke="#4572A7" transform="translate(30,14)"></path>
<text x="123" y="18" fill="#3E576F" style="cursor:pointer;color:#3E576F;fill:#3E576F;" zIndex="2"><tspan x="123">Submitted</tspan></text>
<path d="M -21 0 L -5 0" fill="none" stroke-width="2" zIndex="2" stroke="#AA4643" transform="translate(123,14)"></path><circle cx="-13" cy="-4" r="4" stroke="#4572A7" stroke-width="0" fill="#4572A7" zIndex="3" transform="translate(30,18)"></circle>
<path d="M -13 -8 L -9 -4 -13 0 -17 -4 Z" fill="#AA4643" stroke="#AA4643" stroke-width="0" zIndex="3" transform="translate(123,18)"></path></g><g class="highcharts-tooltip" zIndex="8" visibility="hidden" transform="translate(612,267)"><rect x="7" y="7" width="109" height="56" rx="5" ry="5" fill="none" fill-opacity="0.85" stroke-width="5" isShadow="true" stroke="rgb(0, 0, 0)" stroke-opacity="0.05" transform="translate(1,1)"></rect><rect x="7" y="7" width="109" height="56" rx="5" ry="5" fill="none" fill-opacity="0.85" stroke-width="3" isShadow="true" stroke="rgb(0, 0, 0)" stroke-opacity="0.1" transform="translate(1,1)"></rect><rect x="7" y="7" width="109" height="56" rx="5" ry="5" fill="none" fill-opacity="0.85" stroke-width="1" isShadow="true" stroke="rgb(0, 0, 0)" stroke-opacity="0.15000000000000002" transform="translate(1,1)"></rect><rect x="7" y="7" width="109" height="56" rx="5" ry="5" fill="rgb(255,255,255)" fill-opacity="0.85" stroke-width="2" stroke="#AA4643"></rect>
<text x="12" y="24" fill="#000000" style="color:#333333;font-size:12px;padding:0;white-space:nowrap;fill:#333333;" zIndex="1"><tspan style="font-weight:bold" x="12">Submitted</tspan><tspan x="12" dy="16">Date: </tspan><tspan style="font-weight:bold" dx="3">2018-09-04</tspan><tspan x="12" dy="16">Number: </tspan><tspan style="font-weight:bold" dx="3">2,969</tspan></text></g>
<text x="790" y="395" fill="#909090" style="cursor:pointer;color:#909090;font-size:10px;font-family:'Lucida Grande', 'Lucida Sans Unicode', Verdana, Arial, Helvetica, sans-serif;" text-anchor="end" transform="rotate(0 790 395)" zIndex="8"><tspan x="790">Highcharts.com</tspan></text><g class="highcharts-tracker" zIndex="9" transform="translate(80,50)">
<path d="M 6.568627450980392 128.0219170476456 L 225.52287581699346 111.65145745501502 444.47712418300654 22.73385564759488 663.4313725490196 12.38095238095238" fill="none" isTracker="true" stroke-opacity="0.005" stroke="rgb(192,192,192)" stroke-width="22" stroke-linecap="round" visibility="visible" zIndex="1" style=""></path>
<path d="M 6.568627450980392 259.89570155255734 L 225.52287581699346 259.847704726041 444.47712418300654 259.7976317009222 663.4313725490196 259.6373980205419" fill="none" isTracker="true" stroke-opacity="0.005" stroke="rgb(192,192,192)" stroke-width="22" stroke-linecap="round" visibility="visible" zIndex="1" style=""></path>
</g>
`
    c.svg.Set("innerHTML", inner)
*/
}
