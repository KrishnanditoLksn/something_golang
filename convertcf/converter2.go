package convertcf

import (
	"fmt"
	"os"
	"somthing_golang/tempconv"
)

func ConvertCelcius() {
	var t float64

	fmt.Print("Masukkan suhu: ")
	_, err := fmt.Scanln(&t)
	if err != nil {
		fmt.Fprintln(os.Stderr, "input error:", err)
		os.Exit(1)
	}

	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)

	fmt.Printf("%s = %s\n", f, tempconv.FToC(f))
	fmt.Printf("%s = %s\n", c, tempconv.CToF(c))
}
