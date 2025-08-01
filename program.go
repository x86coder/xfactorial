package main

import(
	"fmt"
	"os"
	"strconv"
	"math"
	"runtime"
)

func main(){
	fmt.Println("xfactorial v0.32 -- by x86coder");
	if len(os.Args) < 2{
		fmt.Println("Usage: xfactorial {natural_number_or_zero}");
	} else{
	
		var factorial, targetNumber, unitSize, start float64
		var processors int
		
		c := make(chan float64);
		processors = runtime.GOMAXPROCS(0)
		targetNumber, err := strconv.ParseFloat(os.Args[1], 64)
		
		if err != nil || targetNumber < 0.0{
			fmt.Println("Error. Input argument not valid!");
			fmt.Println("Usage: xfactorial {natural_number_or_zero}");
		} else{
		
			targetNumber = math.Floor(targetNumber)
			
			/* Do not enter loop if machine is single-core */
			if processors == 1 || targetNumber < 4.0{
				// Stop value is 1.0, factorial(x) starts multiplying from above
				factorial = factorial1(targetNumber, 1.0);
			} else{
				unitSize = math.Floor(targetNumber/float64(processors))
				//fmt.Println("[Unit of Execution] size = ", unitSize)
				
				// Start each worker thread ...
				start = 1.0
				for i := 0; i<processors; i++{
					// Wait for the last iteration to input/process the remainder from the division above (targetNumber/processors) e.g. If targetNumber is not even, then always will be a remainder because # of cores is mostly always multiple of 2.
					if i == (processors-1){
						fmt.Println(" > go routine() with: ", start, targetNumber);
						go xfactorial(c, start, targetNumber);
					} else{
						fmt.Println(" > go routine() with: ", start, start+unitSize-1.0);
						go xfactorial(c, start, start+unitSize-1.0);
					}
					start = start+unitSize
				}
				
				// <- c means we receive from channel "C", channel is automatically
				//  locked (on it's thread) until it is available for read
				factorial = 1.0 // Identity value
				for j := 0; j < processors; j++{
					u := <- c	// Receive from channel
					factorial = factorial * u
				}
			}
			
			fmt.Printf(" > (%f)! = %f", targetNumber, factorial);
		}
	}
}