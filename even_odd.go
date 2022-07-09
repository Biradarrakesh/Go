package main

import (
	"fmt"
	"math/rand"
	"time"
)

type numbers []int

func newDeckOfNums() numbers {
	randInts := numbers{}
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)

	for i := 0; i < 10; i++ {
		randomInteger := generator.Int()
		randInts = append(randInts, randomInteger)
	}
	return randInts
}

func (nums numbers) evenOrOdd() {
	for _, number := range nums {
		if number%2 == 0 {
			fmt.Printf("I %d, i am Even \n", number)
		} else {
			fmt.Printf("I %d, i am Odd \n", number)
		}
	}
}
