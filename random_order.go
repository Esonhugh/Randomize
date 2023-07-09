package randomize

import "math/rand"

// RandomOrder func help you to shuffle the order of string array
// If you got region list or other string Array
// make random order of it to avoid the Cloud Platform Providers' `Cloud Security Alert module` to detect.
func RandomOrder[T any](s []T) []T {
	if !ForceRandomOrderOpt {
		return s
	}
	return ForceRandomOrder(s)
}

// ForceRandomOrder is function to shuffle the order of string array force, ignore the AdvancedEvasion
func ForceRandomOrder[T any](s []T) []T {
	r := rand.New(rand.NewSource(rand.Int63()))
	r.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	return s
}

// RandomOrderMap is a function to shuffle the order of map
// If you got region list or other map
func RandomOrderMap[M map[K]V, K comparable, V any](m M) M {
	if !ForceRandomOrderOpt {
		return m
	}
	return ForceRandomOrderMap(m)
}

// ForceRandomOrderMap is a function to shuffle the order of map force, ignore the AdvancedEvasion
func ForceRandomOrderMap[M map[K]V, K comparable, V any](m M) M {
	r := rand.New(rand.NewSource(rand.Int63()))
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	r.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})
	newMap := make(map[K]V)
	for _, k := range keys {
		newMap[k] = m[k]
	}
	return newMap
}
