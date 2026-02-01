package constant

import "fmt"

const boilingF = 212.0 ;

func Boil()  {
	var f  = boilingF;
	var c = (f-32) * 5 / 9 ;

	fmt.Printf("Boil Point = %g F or %g C\n" , f,c)

}