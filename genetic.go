package main

import (
	"math/rand/v2"
	"slices"
)

// Individual represents a candidate solution. The genome is a bitmask
// where each bit indicates whether the corresponding item is selected.
type Individual struct {
	Genome  uint64
	Fitness int
}

// newIndividual creates an individual and evaluates its fitness.
func newIndividual(genome uint64) Individual {
	return Individual{
		Genome:  genome,
		Fitness: fitness(genome),
	}
}

// randomGenome generates a random 64-bit genome.
func randomGenome() uint64 {
	return rand.Uint64()
}

// breed combines two parents using single-point crossover.
// A random split point divides the genome: bits above the point come from
// parent1, bits below from parent2. The result is then mutated.
func breed(parent1, parent2 Individual) Individual {
	breedPoint := rand.IntN(GenomeSize)
	var mask uint64
	if breedPoint > 0 {
		mask = (uint64(1) << breedPoint) - 1
	}
	babyGenome := (parent1.Genome & ^mask) | (parent2.Genome & mask)
	babyGenome = mutate(babyGenome, 0.015)
	return newIndividual(babyGenome)
}

// mutate flips each bit independently with the given probability.
func mutate(genome uint64, rate float64) uint64 {
	for i := range GenomeSize {
		if rand.Float64() < rate {
			genome ^= uint64(1) << i
		}
	}
	return genome
}

// sortByFitness sorts the population in descending order by fitness.
func sortByFitness(pop []Individual) {
	slices.SortFunc(pop, func(a, b Individual) int {
		return b.Fitness - a.Fitness
	})
}

// isBitSet returns true if the bit at the given index is 1.
func isBitSet(data uint64, index int) bool {
	return (data & (uint64(1) << index)) != 0
}
