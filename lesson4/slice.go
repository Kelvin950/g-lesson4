package main

import (
	"fmt"
	"math"
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
	

	 a := []int {1,2,3,4,5}  
	 start ,stop := 0,  3  

	fmt.Println( a[start:stop])
 
	dst = make([]int ,len(numbers)-1) 

	copy(dst ,append(numbers[:3] , 100)) 
	fmt.Println(dst)
	crecu(a) ;
  
	 
}


func crecu ( r[]int) {


	if(len(r) >0){
 
	  
	midlle := int(math.Round(float64(len(r)/2)))

 

	crecu(r[0:midlle]) 
	    fmt.Println(r , 22) ;
 crecu(r[midlle:len(r)-1])
	}
 
	

}