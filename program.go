package main

import(
	"fmt"
	"os"
	"strconv"
	"math"
	"runtime"
)


func main(){
	fmt.Println("xfactorial v0.3 -- by x86coder");
	if len(os.Args) < 2{
		fmt.Println("Usage xfactorial {number}");
	} else{
	
		var factorial, targetNumber, unitSize, start float64
		var processors int
		
		c := make(chan float64);
		processors = runtime.GOMAXPROCS(0)
		targetNumber, err := strconv.ParseFloat(os.Args[1], 64)
		
		if err != nil{
			fmt.Println("Error. Input argument not valid!");
			fmt.Println("Usage xfactorial {number}");
		} else{
		
			/* Add if condition for the case where processors = 1 */
			
			unitSize = math.Floor(targetNumber/float64(processors))
			//fmt.Println("[Unit of Execution] size = ", unitSize)
			
			start = 1.0
			for i := 0; i<processors; i++{
				if i == (processors-1){
					fmt.Println(" > go routine() with: ", start, targetNumber);
					go xfactorial(c, start, targetNumber);
				} else{
					fmt.Println(" > go routine() with: ", start, start+unitSize-1.0);
					go xfactorial(c, start, start+unitSize-1.0);
				}
				start = start+unitSize
			}
			
			factorial = 1.0 // Identity value
			for j := 0; j < processors; j++{
				u := <- c	// Receive from channel
				factorial = factorial * u
			}
			
			fmt.Printf(" > %f! = %f", targetNumber, factorial);
		}
	}
}