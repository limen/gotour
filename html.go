package main

import (
    "fmt"
    "html"
)

var p = `
hello


world!
<script>alert(111);</script>
    my name is        alice;
`

func main() {
    fmt.Println(p)
    fmt.Println(html.EscapeString(p))
}
