package main

import (
	"fmt"

	"GoPlayGround/NodeList"
)

func main(){
	a := NodeList.NodeList[string]{Next: nil, Val: "Hello"}
	a.Append("Goobye")
	fmt.Println(a)
}
