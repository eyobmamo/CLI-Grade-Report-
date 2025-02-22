package main

import (
	"fmt"
)

func main(){
	fmt.Println("welcome to grede calculator")
	var name string
	var no_courses int
	subjects := make(map[string]float32)

	fmt.Println("Inter Your name: ")
	fmt.Scanln(&name)
	fmt.Println("Inter no subject you take:")
	fmt.Scanln(&no_courses)

	for current_course:=1; current_course <= no_courses; current_course +=1 {
		var course_name string 
		var course_result float32 
		fmt.Println("Enter the cousrse name:")
		fmt.Scanln(&course_name)
		

		for {
			fmt.Printf("Eneter course result of %s . \n",course_name)
			fmt.Scanln(&course_result)
			if course_result >= 0 && course_result <=100 {
				break
			} else {
				fmt.Println("invalid resutl inter value b/n 0 and 100")
			}
		}
		subjects[course_name] = course_result
	} 

	fmt.Println("welcome to Gread report")
	fmt.Printf("Name : %s \n",name)
	count := 1
	for name,result := range subjects {
		fmt.Println("-----------------------------------------------------------------------")
		fmt.Printf("| %d. subject : %s  | result : %f \n |",count,name,result )
		
		count += 1

	}
	Average := Avrage_calculator(subjects)
	
		fmt.Println("-----------------------------------------------------------------------")
		fmt.Printf("|  the Avrage Result : %.2f \n|",Average)
		fmt.Println("-----------------------------------------------------------------------")


	var input string
	fmt.Scanln(&input)


	
}

func Avrage_calculator(result map[string]float32) float32 {
	var Total float32
	var Avrage float32
	amount := len(result)

	for _,res := range result {
		Total +=res
	}

	Avrage = Total/ float32(amount)
	return Avrage
}
