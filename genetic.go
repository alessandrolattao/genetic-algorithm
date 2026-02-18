package main

import (
	"math/rand/v2"
	"slices"
)

type Individual struct {
	Genome  uint64
	Fitness int
}

func newIndividual(genome uint64) Individual {
	return Individual{
		Genome:  genome,
		Fitness: fitness(genome),
	}
}

func randomGenome() uint64 {
	return rand.Uint64()
}

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

func mutate(genome uint64, rate float64) uint64 {
	for i := range GenomeSize {
		if rand.Float64() < rate {
			genome ^= uint64(1) << i
		}
	}
	return genome
}

func sortByFitness(pop []Individual) {
	slices.SortFunc(pop, func(a, b Individual) int {
		return b.Fitness - a.Fitness
	})
}

func isBitSet(data uint64, index int) bool {
	return (data & (uint64(1) << index)) != 0
}
