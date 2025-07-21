package main

import "fmt"

func convertKtoC(tempK float64) float64 {
	return tempK - 273
}

func main() {

	tempK := 373
	fmt.Println(convertKtoC(float64(tempK)))

}
