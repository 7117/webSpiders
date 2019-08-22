package main
 
import (
	"fmt"
)
 
func main(){
	// 方式一 :原生普通方式
	var mapp map[string]float32;
	mapp = make(map[string]float32);
	mapp["name"]=11.11;
	fmt.Printf("%v\n",mapp);
	fmt.Println(mapp);
 
 
	// 方式二：短声明
	mappp:=make(map[string]float32);
	mappp["sex"]=22.22;
	fmt.Printf("%v\n",mappp);
	fmt.Println(mappp);
 
	// 方式三：结合
	var mapppp = make(map[string]float32);
	mapppp["school"]=33.33
	mapppp["school2"]=33.33
	fmt.Println(mapppp);
 
}