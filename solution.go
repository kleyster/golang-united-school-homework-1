package solution

import (
	"fmt"
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	s1 := emoji.Sprint("Hello :world_map:")
	fmt.Println(s1)
}
