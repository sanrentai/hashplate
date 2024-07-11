package hashpalte

func splitMix32(seed uint32) func() float64 {
	// Return a function that generates a pseudo-random number
	return func() float64 {
		seed |= 0
		seed = seed + 0x9e3779b9
		hash := seed ^ (seed >> 16)
		hash = hash * 0x21f0aaad
		hash = hash ^ (hash >> 15)
		hash = hash * 0x735a2d97
		hash = hash ^ (hash >> 15)
		return float64(hash) / 4294967296.0
	}
}

func getSeedFromString(value string) uint32 {
	seed := uint32(0)
	for _, char := range value {
		seed = (seed*31 + uint32(char))
	}
	return seed
}
