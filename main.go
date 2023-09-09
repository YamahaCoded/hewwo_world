package main

import "fmt"

func main() {
	var rainbowColors = [6]string{"\033[31m",
		"\033[33m",
		"\033[32m",
		"\033[36m",
		"\033[34m",
		"\033[35m",
	}
	for true {
		for i := 0; i < 6; i++ {
			fmt.Println(rainbowColors[i] + "hewwo, world!!" + "\033[0m")
		}
	}
}