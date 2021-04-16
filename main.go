package main

import (
	"fmt"

	"example.com/testgo/mascot"
	"rsc.io/quote"
)

func main() {
	fmt.Println(mascot.BestMascot())
	fmt.Println(quote.Go())
	fmt.Println("한글 출력 테스트")
}
