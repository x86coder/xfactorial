package main

import(
	"fmt"
	"os"
)

func main(){
	fmt.Println("xfactorial v0.1 -- by x86coder");
	if os.Args[1] == ""{
		fmt.Println("Usage xfactorial \"number\"");
	} else{
		fmt.Println("3");
	}
}