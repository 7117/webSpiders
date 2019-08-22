package main
 
import (
	"fmt"
)
 
func main(){
	// 定义赋值一个数组方式一
	var a [10]byte
	a[0]='b';
	fmt.Printf("%c\n",a);
 
	// 定义赋值一个数组方式二
	var b [10]int = [10]int{1,2,3}
	fmt.Printf("%d\n",b);
 
	// 定义赋值一个数组方式三
	c:=[10]int{1,2,3}
	fmt.Printf("%d\n",c);
 
}
 
// [1 0 0 0 0 0 0 0 0 0]
// [1 2 3 0 0 0 0 0 0 0]
// [1 2 3 0 0 0 0 0 0 0]