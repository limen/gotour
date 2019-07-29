package main

import (
	"fmt"
	"strings"
)

func main() {
	p := "China 中国 USA 美国"
	fmt.Println(p)
	fmt.Println(len(p))
	fmt.Println(p[0:11])
	fmt.Println(p[0:12])
	r := []rune(p[0:11])
	fmt.Println(r)
	fmt.Println(string(r[0:7]))
	p = `
    


aaa bbb

ccc


eee    

	
				

`
	fmt.Println(strings.TrimSpace(p))
}
