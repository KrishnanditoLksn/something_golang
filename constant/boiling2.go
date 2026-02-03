package constant

import "fmt"

func Boiling2(){
	const freezingF , boilingF = 32.0,212.0;
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF)) // "212°F = 100°C"
}

func fToC(f float64) float64 {
	return (f-32) * 5 / 9;
}



func Fib(n int )int{
	x , y := 0,1
	for range n {
		x , y = y , x + y
	}
	return x ;
}