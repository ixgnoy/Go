package main

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math Error: Square root of negative number")
	}
	// Then compute the square root
	return 1, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Data is empty")
	}
	return nil
}

func main() {
	// result, err := sqrt(16)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Square root:", result)

	// result2, err2 := sqrt(-16)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// 	return
	// }
	// fmt.Println("Square root:", result2)

	//Initialize empty data
	data := []byte{}
	// Short hand notation for error handling
	// if err := process(data); err != nil {
	err := process(data)
	if err != nil {
		fmt.Println("Error:",err)
		return
	}
	fmt.Println("Data processed successfully")
	
}
