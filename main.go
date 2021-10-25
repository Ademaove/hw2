package hw2

import (
	"fmt"
	"github.com/Ademaove/hw2Lib"
)

func main() {
	fmt.Printf(hw2Lib.AlterString("aBcZYop-iqz."))

	x1, x2 := hw2Lib.GetRoots(1, 9, 3)
	fmt.Println(x1, x2)

	fmt.Println(hw2Lib.GetUUID())
}