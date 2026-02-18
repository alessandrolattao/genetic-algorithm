package main

import "math/rand/v2"

const (
	PopSize     = 50
	Generations = 500
)

func main() {
	population := make([]Individual, PopSize)
	for i := range population {
		population[i] = newIndividual(randomGenome())
	}

	for gen := range Generations {
		sortByFitness(population)

		printSolution(population[0], gen, false)

		newPop := make([]Individual, PopSize)
		newPop[0] = population[0] // elitism
		newPop[1] = population[1]

		for i := 2; i < PopSize; i++ {
			parent1 := population[rand.IntN(PopSize/3)]
			parent2 := population[rand.IntN(PopSize/3)]
			newPop[i] = breed(parent1, parent2)
		}

		population = newPop
	}

	sortByFitness(population)
	printSolution(population[0], Generations, true)
}
