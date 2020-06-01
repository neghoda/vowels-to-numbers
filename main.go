package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/neghoda/vowels-to-numbers/vwlstonums"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text: ")
	text, _ := stdin.ReadString('\n')

	fmt.Println("Enter 1 to encoding, 2 for decoding ")
	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch num {
	case 1:
		fmt.Println(vwlstonums.Encode(text))
	case 2:
		fmt.Println(vwlstonums.Decode(text))
	default:
		fmt.Println("Unsuported ")
	}
}
