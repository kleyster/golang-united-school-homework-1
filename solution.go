package solution

import (
	"fmt"
	"github.com/kyokomi/emoji/v2"
)

func main(){
	return GetMessage()
}

func GetMessage() string {
	return emoji.Sprint("Hello :world_map:")
}