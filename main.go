package main

import (
	"fmt"

	"GoPlayGround/nodelist"
)

func main() {
	a := nodelist.NodeList[string]{Next: nil, Val: "Hello"}
	a.Append("Goobye")
	fmt.Println(a)
}
