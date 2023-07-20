package main

import (
	"fmt"
	"math"
	"unsafe"
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
  
	 

	//let see some changes  with the internal part of slice

	arr1 := []int{1,2,3,44} 

	slice1 ,slice2 := make([]int , 4) , make([]int, 5) 

	copy(slice1 , arr1) 
	copy(slice2,arr1)
 
   
arr1[1] =  100 

fmt.Println(slice1, slice2)

 //checking memory size 

  aew := [5]int{1,2,3,4,5} 
 aswe := []int{1,2,3,4,5,6} 

 fmt.Println(unsafe.Sizeof(aew))
 fmt.Println(unsafe.Sizeof(aswe))

 
 muslice := []float64{1.2, 1.3 , 1.2} 
 muslice =  append(muslice, []float64{1.3,1.4,5.5}...)
  

 fmt.Println(muslice[:3])
 var muslice2  =  []float64{}
 
 muslice2 =  append(muslice2, muslice...)
 fmt.Print(muslice2)

  years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
  
  newYears :=  []int{} 

   newYears =   append(years[:3], years[len(years)-3:]...)

   fmt.Println(newYears)

    
}


func crecu ( r[]int) {


	if(len(r) >0){
 
	  
	midlle := int(math.Round(float64(len(r)/2)))

 

	crecu(r[0:midlle]) 
	    fmt.Println(r , 22) ;
 crecu(r[midlle:len(r)-1])
	}
 

	



}