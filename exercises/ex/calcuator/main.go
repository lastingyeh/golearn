package main

import "fmt"

func main(){

	var n1, n2, op, result int

	for {
		fmt.Println("-----caculator-----")
		fmt.Println("number1, number2")
		fmt.Scanf("%d %d", &n1, &n2)
		fmt.Println("1. add")
		fmt.Println("2. sub")
		fmt.Println("3. mul")
		fmt.Println("4. div")
		fmt.Println("0. exit")
	
		fmt.Scanln(&op)

		switch op {
		case 0:
			break
		case 1:
			result = n1 + n2
		case 2:
			result = n1 - n2
		case 3:
			result = n1 * n2
		case 4:
			result = n1 / n2
		default:
			continue
		}
		fmt.Println("result: ", result)
	}
}