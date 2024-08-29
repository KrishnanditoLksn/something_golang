package main

import "fmt"

func main()  {
	type footballer string
	var a int32 = 12321
	var b int16 = 32123
	fmt.Println("Testing" , a)
	fmt.Println("Konversi nilai : " , b)

	var name = "Dito"
	var e uint8 = name[0]
	var str = string(e)


	var yourname footballer = "Dino"
	fmt.Println("Converted Nilai" , str)

	fmt.Println(yourname)

	var messi footballer = "Lionel"
	var namer footballer = footballer(messi)
	fmt.Println(namer)
}