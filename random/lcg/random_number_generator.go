package lcg

import "time"

// RandomNumberGenerator define basic params of
// a *Linear Congruential Generator*
type RandomNumberGenerator struct {
    // seed
    seed int64;
    // multiplier
    a int64;
    // increment
    c int64;
    // modulus
    modulus int64;
}

// NextInt get the next randome integer between 0 to rng.modulus
func (rng *RandomNumberGenerator) NextInt() int64 {
    rng.seed = (rng.seed * rng.a + rng.c) % rng.modulus
    return rng.seed
}

// NextFloat get the next randome integer between 0 to 1
func (rng *RandomNumberGenerator) NextFloat() float64 {
    random := float64(rng.NextInt()) / float64(rng.modulus)
    return random
}

// SetSeed set the seed for subsequent rng calculations
func (rng *RandomNumberGenerator) SetSeed(seed int64) {
    rng.seed = seed
}

// GetRandomNumberGenerator return an rng instance using lcg algorithm
func GetRandomNumberGenerator(a, c, modulus int64) RandomNumberGenerator {
    seed := time.Now().UnixNano()
    rng := RandomNumberGenerator{seed, a, c, modulus}
    return rng
}
