package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 0; i < 65000; i++ {
		fmt.Printf("%d \t %b \t %#x \t %c \t %U \t %q \n", i, i, i, i, i, i)

	}

}
