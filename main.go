package main

import (
	"fmt"

	"github.com/rcsolis/nutrition_calculator/internal"
)

/**
 * Entrypoint for the application
 */
func main() {
	// Get nutritional score
	ns := internal.Run()
	// Print result
	fmt.Println("Nutritional Score:")
	fmt.Println(ns)
}
