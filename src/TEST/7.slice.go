package main
 
import (
	"fmt"
)
 
func main(){
	var array [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var slice = array[0:2] 
	fmt.Printf("%d\n",slice);
	
	for i,v:= range slice {
		fmt.Printf("%d  %d\n",i,v);

	}

	fmt.Printf("%d",slice[0]);
 
}