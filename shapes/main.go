package main

func main() {
	t := triangle{
		heightLength: 23.05677,
		baseLength:   2.333,
	}

	s := square{
		sideLength: 23.45,
	}

	printArea(t)
	printArea(s)
}
