package main

import (
	"fmt"
	"github.com/drhodes/golorem"
)

func main() {
	fmt.Println(golorem.Sentence(7, 11))
	fmt.Println(golorem.Paragraph(7, 11))
}
