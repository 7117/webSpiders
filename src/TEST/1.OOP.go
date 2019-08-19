package main

import (
	"fmt"
)

type Point struct {
	px float32;
	py float32;
}
// 记住哦  *Point前面的pont只是一个形参  就是实例化对象的形参  只是在此函数中有用
// 出了此函数  没有意义了
// 后面在进行实例化  调用此函数的时候  还会进行再次实例化  那才是实参
// 所以说前面的虚参实例化是没有意义的，其实际地址是不一样的  只是名称可能会一样  
// 所以即使实例化对象的名称一样  他们的值也没有办法进行跨函数传值设置值  所以只能通过引用传值进行传值
// 对类的属性进行设置值
// 这样就保存在了类的成员属性之中里面了
func (pont *Point) setName(px,py float32){
	pont.px=px;
	pont.py=py;
	//不加*返回1200  加*返回1212
	//不加*的话  意味着这边只是复制值  复制给了point.px point.py
	// 但是之后的数值没有更改  说明只是对当前的方法里面的值进行了更改  在方法之外不会有效
	fmt.Print(pont.px,pont.py);
}

func (poin *Point) getName()(float32,float32){
	return poin.px,poin.py;
}


func main(){
	point:=new(Point);
	point.setName(1.0,2.0);
	px,py:=point.getName();
	fmt.Print(px,py);

}