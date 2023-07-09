package randomize

import (
	"math/rand"
	"time"
)

// RandomChoice will return a random element from the given arguments.
func RandomChoice[T any](args ...T) T {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	i := rand.Intn(len(args))
	return args[i]
}

// RandomChoiceArray will return a slice of chosen random elements from the given arguments.
// But the result may a slice contain duplicate elements and empty slice.
func RandomChoiceArray[T any](args ...T) []T {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	size := rand.Intn(len(args))
	return RandomChoiceArrayBySize(size, args...)
}

// RandomChoiceArrayNoEmpty will return a slice of chosen random elements from the given arguments.
// But the slice will not be empty. still may contain duplicate elements.
func RandomChoiceArrayNoEmpty[T any](args ...T) []T {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	size := rand.Intn(len(args))
	if size == 0 {
		size = 1
	}
	return RandomChoiceArrayBySize(size, args...)
}

// RandomChoiceArrayBySize will return a slice of chosen random elements from the given arguments.
// if size < 0 return nil.
// if size > len(args) return args.
func RandomChoiceArrayBySize[T any](size int, args ...T) []T {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	if size <= 0 {
		return nil
	}
	if size > len(args) {
		size = len(args)
		return args
	}

	arr := make([]T, 1)
	for j := 0; j < size; j++ {
		arr = append(arr, args[rand.Intn(len(args))])
	}
	return arr
}
