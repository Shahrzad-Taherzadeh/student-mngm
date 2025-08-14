package main

import "fmt"


func  main()  {
	
	fmt.Println("hello world!")

	type information struct{
		name string
		age int
	}

	checkInformation := map[int]information{
		1234 : {"Ali", 17},
		12 : {"Shahrzad", 18},
		34 : {"Amir", 20},
		56 : {"Zeynab", 55},
		78 : {"Mahmoud", 20},
	}

	var idNumber int

	for{

	fmt.Print("Enter ur ID number => ")
	fmt.Scan(&idNumber)
	
	v, ok := checkInformation[idNumber]
	if ok{
		fmt.Printf("name:%s age:%d", v.name, v.age)
	}else{
		fmt.Println("...Not found...")
		continue
	}
	break
	}
}

