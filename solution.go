package main

import (
	"fmt"
	"github.com/kyokomi/emoji/v2"
)

func main(){
	s1:=emoji.Sprint("Hello :world_map:")
	fmt.Println(s1)
}

