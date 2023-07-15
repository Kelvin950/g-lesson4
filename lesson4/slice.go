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
       

	  //append to a slice or copy to a slice
    
	 numbers1 := []int{1,2,3,4} 

     	numbers=  append(numbers1 , 10 )
    fmt.Println(numbers) ;
	n3 := []int{100 , 200} 
	numbers1 =  append(numbers1, n3...) 
	
	   src:= []int{10,20,30} 
	   dst:=  make([]int,  len(src)) ;
	       copy(dst , src) ;
	   
	   fmt.Println(src , dst ) ;
	

	
}