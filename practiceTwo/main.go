package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

type Students struct {
	Name       string
	Age        string
	Id         int
	SignUpTime time.Time
}


func memUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Allocated : %.2f Mib\n",float64(m.Alloc)/1024/1024)
	fmt.Printf("TotalAllocated : %.2f Mib\n",float64(m.TotalAlloc)/1024/1024)
	fmt.Printf("Sys : %.2f Mib\n",float64(m.Sys)/1024/1024)
	fmt.Println("GC : ",m.NumGC)

}

func main() {

	rand.Seed(time.Now().UnixNano())

	fmt.Println("-------Student Signup-------")
	var s Students
	fmt.Printf("Enter your name : ")
	fmt.Scan(&s.Name)

	for {

		fmt.Printf("Enter your age : ")
		fmt.Scan(&s.Age)

		int_age, err := strconv.Atoi(s.Age)
		if err != nil {
			fmt.Println("invalid age, please enter a number")
			continue
		} else {
			fmt.Println(int_age)
			break
		}

	}

	s.Id = rand.Intn(90000) + 10000
	s.SignUpTime = time.Now()

	fmt.Println("you're succesfully regestered")
	fmt.Println("Name : ",s.Name)
	fmt.Println("Age : ",s.Age)
	fmt.Println("Id : ",s.Id)
	fmt.Println("Sign Up time : ",s.SignUpTime.Format("2006-01-02 15:04:05"))
	fmt.Println("----------------------------------------------")
	memUsage()
}
