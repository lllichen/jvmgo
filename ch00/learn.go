package main

import "fmt"

//func main(){

	//c := [4]byte{0xCA,0xFE,0xBA,0xBE}
	//
	//for _,v := range c {
	//	fmt.Printf("%b\t",v)
	//}
	//fmt.Println(c)
	//
	//d := [8]byte{0xC,0xA,0xF,0xE,0xB,0xA,0xB,0xE}
	//for _,v := range d {
	//	fmt.Printf("%b\t",v)
	//}
	//
	//
	//z := 0XCAFEBABE
	//fmt.Printf("%b",z)

//}

//func main(){
//	defer func(){ // 必须要先声明defer，否则不能捕获到panic异常
//		fmt.Println("c")
//		if err:=recover();err!=nil{
//			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
//		}
//		fmt.Println("d")
//	}()
//	f()
//}
//
//func f(){
//	fmt.Println("a")
//	panic(55)
//	fmt.Println("b")
//	fmt.Println("f")
//}

func main() {
	fmt.Println(1024&1024)
}

