package main

import "fmt"

func main() {

	for i := 0; i < 100; i++ {

		if i%3 == 0 {
			fmt.Print("fizz \n")
		}
		if i%5 == 0 {
			fmt.Print("buzz \n")
		}

		if i%3 == 0 && i%5 == 0 {
			fmt.Print("fizzbuzz \n")
		}
		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
		}
	}

}
