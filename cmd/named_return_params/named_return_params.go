package main

import "fmt"

func main() {
	a, b, c, d := namedReturnParams()
	fmt.Printf("a = %v\nb = %v\nc = %v\nd = %d\n", a, b, c, d)
}

func namedReturnParams() (
	a *string,
	b *string,
	c string,
	d int,
) {
	algo := "alalalala"
	a = &algo
	return
}
