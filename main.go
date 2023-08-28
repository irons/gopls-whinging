package main

import (
	"fmt"
	"github.com/drhodes/golorem"
)

func main() {
	fmt.Println(lorem.Sentence(7, 11))
	fmt.Println(lorem.Sentence(2, 5))
}
