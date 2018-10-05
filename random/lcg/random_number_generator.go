package lcg

import (
	"encoding/binary"
    "time"
    "math"
)

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

// RandomBytes generate specific size of random bytes
func (rng *RandomNumberGenerator) RandomBytes(len int32) []byte {
    // NOTE: modulus has to be lower or equal to 2^32
    uint32SliceLen := int(math.Ceil(float64(len) / float64(4)))
    output := make([]byte, uint32SliceLen * 4)

    for index := 0; index < uint32SliceLen; index++ {
        // Generate a random uint32 and count it as 4 bytes 
        randomNumber := uint32(rng.NextInt())
        startIndex := index * 4
        binary.LittleEndian.PutUint32(output[startIndex:], randomNumber)
    }

    return output[:len]
}

// SetSeed set the seed for subsequent rng calculations
func (rng *RandomNumberGenerator) SetSeed(seed int64) {
    rng.seed = seed
}

// GetRandomNumberGeneratorCustom return an rng instance 
// with custom params using lcg algorithm
func GetRandomNumberGeneratorCustom(a, c, modulus int64) RandomNumberGenerator {
    seed := time.Now().UnixNano()
    rng := RandomNumberGenerator{
        seed, a, c, modulus,
    }
    return rng
}

// GetRandomNumberGenerator return an rng instance 
// with default params using lcg algorithm
// (Source: Numerical Recipes)[https://en.wikipedia.org/wiki/Linear_congruential_generator]
func GetRandomNumberGenerator() RandomNumberGenerator {
    return GetRandomNumberGeneratorCustom(
        // a
        1664525,
        // c
        1013904223,
        // modulus
        int64(math.Pow(2, 32)),
    )
}