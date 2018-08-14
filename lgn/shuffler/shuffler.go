// Package shuffler shuffles slices of integers
// deterministically based on the ID of the user.
package shuffler

import (
	"log"
	"math/rand"
	"strconv"
)

type shufflerT struct {
	Variations  int // How many different shuffled sets should be assigned to the IDs
	MaxElements int // Max number of elements to shuffle

	ID int // Basis for the variations; typically the user ID
}

// New creates a shufflerT with the reusable data,
// to generate shuffled slices
func New(variatons, maxNumberOfElements int, ID string) *shufflerT {
	s := shufflerT{}
	s.Variations = variatons
	s.MaxElements = maxNumberOfElements
	s.ID, _ = strconv.Atoi(ID)
	return &s
}

func (s *shufflerT) Slice() []int {

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
		gen.Shuffle(s.MaxElements, swapFct)
		log.Printf("User %v is class %v => %+v", s.ID, class, order)

	}

	return order
}
