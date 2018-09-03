// Package shuffler creates slices of integers
// random, but reproducible; based on the ID of the user;
// classes of users see the same random order each time they visit;
// each page has a different randomization of appropriate length.
package shuffler

import (
	"log"
	"math/rand"
	"strconv"
)

type shufflerT struct {
	ID         int // Seed for shuffling; typically the user ID
	Variations int // ID modulo Variations is the actual seed. Determines how many different shuffled sets are derived from various IDs

	MaxElements int // The number of elements to shuffle; typically the largest number of input groups across all pages of a questionaire.
}

// New creates a Shuffler for creating deterministic variations
// of a slice []int{1,2,3...,MaxElements}
//
// ID is the seed for the randomizer
// Variations is the number of classes.
// MaxElements is the slice length.
func New(ID string, variatons int, maxNumberOfElements int) *shufflerT {
	s := shufflerT{}
	s.Variations = variatons

	s.MaxElements = maxNumberOfElements
	s.ID, _ = strconv.Atoi(ID)
	return &s
}

// Slice generates a shuffled slice.
// Param iter gives the number of shufflings; typically the page number
func (s *shufflerT) Slice(iter int) []int {

	order := make([]int, s.MaxElements)
	for i := 0; i < len(order); i++ {
		order[i] = i // []int{0,1,2,3}
	}
	if s.Variations == 0 {
		// keep the slice
	} else {
		class := int64(s.ID % s.Variations) // user 12; variations 5 => class 2
		src := rand.NewSource(class)        // not ...UTC().UnixNano(), but constant init
		gen := rand.New(src)                // generator seeded with the class of the user ID

		// This does not cover all elements equally
		if false {
			for i := 0; i < len(order); i++ {
				order[i] = gen.Int() % s.MaxElements // gen.Int 19; max elements 4 =>
				log.Printf("%2v: User %v is class %v => of %v", i, s.ID, class, order[i])
			}
		}

		swapFct := func(i, j int) {
			order[i], order[j] = order[j], order[i]
		}
		for i := 0; i <= iter; i++ {
			gen.Shuffle(s.MaxElements, swapFct)
			log.Printf("%2v: User %v is class %v => %+v", i, s.ID, class, order)
		}
	}

	return order
}
