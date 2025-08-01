package main

func factorial1(n float64, stop float64) float64{
	if n == stop{
		return n
	} else{
		if n == 0.0{
			return 1.0
		} else{
			return n * factorial1(n-1.0, stop);
		}
	}
}

func xfactorial(channel chan float64, start float64, end float64){
	// Reverse order, calculate higher value first. Start value is stop condition for recursive function factorial(x)
	channel <- factorial1(end, start);
}