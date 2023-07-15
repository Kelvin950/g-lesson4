package main

import (
	"fmt"


)

 

func main(){
 var numbers [4]int 

 fmt.Printf("%v\n" , numbers) ;
 fmt.Printf("%#v\n" ,numbers)




//ellipsis operator 
a2 := [...]int{1,2,3,4}
fmt.Printf("%#v" , a2) ;
fmt.Println(len(a2))


for i ,v := range(a2) {

	fmt.Println(i , v) ;
}


for i:=0 ; i<len(a2) ;i++{

	fmt.Println(i , a2[i])
}

//2d array

balances :=  [2][3]int {

	{4,5,6} , 
	{21,21,21} ,
}

fmt.Println(balances) ;

accountName :=  [4]string{ 
	0:"kwamr" , 
	2:"keof" ,
	1: "rer" ,
	3:"ueae" ,
}

 
fmt.Println(accountName) 


var city [3]string ;
city[0] =  "Accra" 
city[1] = "Kumasi" 
fmt.Println(city) ;

var grades = [3]float64{1,2,3}   

grades[0] = 2 
  

fmt.Println(grades) ;


taskDone :=[...]bool{}

fmt.Println(taskDone) ;

 
for i:=0 ;i<len(city) ;i++{
 
	fmt.Println(i ,  city[i]) 

}

for index, value := range(grades){


	fmt.Println(index ,value)
}

nums := [...]int{30, -1, -6, 90, -6} 

for _, values :=  range(nums){
 
if values > 0 && values %2 ==0 {
	 
	println(values) ;
}


}

}