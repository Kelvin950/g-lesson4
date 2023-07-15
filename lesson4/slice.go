package main

import (
	"fmt"

	
)


func main(){
 
	var city []string;
	fmt.Println("cities is equal to nill" , city==nil) ; 
   fmt.Printf("%#v\n" , city) 
    fmt.Println(len(city))
 
	
	numbers := []int{1,2,3,4} 
	fmt.Println(len(numbers)) 
	  

	nums:= make([]int , 2) ;
	fmt.Printf("%#v", nums) 

	type names []string 
	friend := names{"Dan", "Maria"} 

	fmt.Printf("%#v" , friend) 
      fmt.Println(friend[0]) 

	  friend[1] = "gab"


	 
	  nums1 :=  numbers ;
	  nums1[1] =  3 
	  fmt.Println(numbers) 

}