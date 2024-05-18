package main

import "fmt"

func main() {
	//defer
	defer fmt.Println("Hola")
	fmt.Println("mundo")

	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		//break
		if i == 8 {
			fmt.Println("break")
			break
		}
	}

	

}
