package main

import "fmt"

func main() {
	fmt.Println("Go Loops Phase")

	days := []string{"Sunday","Monday", "Tuesday", "Wenesday", "Thusday", "Friday", "Saturday"}
	fmt.Println(days)

	// for i:=0; i< len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days{

	// 	fmt.Println(days[i])
	// }

	// for _, day := range days{
	// 	fmt.Printf("The  values are %v\n ", day)
	// }

	rougueValue := 1


	// for rougueValue <10{
	// 	fmt.Println("The rougue value is:", rougueValue)
	// 	rougueValue++
	// }

	for rougueValue <10{
		if rougueValue ==5 {

			
			rougueValue++
			continue
			
		}
		fmt.Println("The rougue value is:", rougueValue)
		rougueValue++

	}
}