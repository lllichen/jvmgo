package main

import "fmt"

func main(){

	c := [4]byte{0xCA,0xFE,0xBA,0xBE}

	for _,v := range c {
		fmt.Printf("%b\t",v)
	}
	fmt.Println(c)

	d := [8]byte{0xC,0xA,0xF,0xE,0xB,0xA,0xB,0xE}
	for _,v := range d {
		fmt.Printf("%b\t",v)
	}
	

	z := 0XCAFEBABE
	fmt.Printf("%b",z)

}
