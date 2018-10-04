package lcg

import "fmt"
import "testing"
import "math"

func ExampleRandomNumberGenerator_NextInt() {
    var seed, a, c, modulus int64
    a = 1664525
    c = 1013904223
    modulus = int64(math.Pow(2, 32))
    seed = 3820783849
    rng := GetRandomNumberGenerator(a, c, modulus)
    rng.SetSeed(seed)
    fmt.Println(rng.NextInt())
    // Output: 1246739764
}

func ExampleRandomNumberGenerator_NextFloat() {
    var seed, a, c, modulus int64
    a = 1664525
    c = 1013904223
    modulus = int64(math.Pow(2, 32))
    seed = 1106396931
    rng := GetRandomNumberGenerator(a, c, modulus)
    rng.SetSeed(seed)
    fmt.Println(rng.NextFloat())
    // Output: 0.052043945994228125
}

func TestRandomNumberGenerator_NextInt(test *testing.T) {
    var seed, a, c, modulus int64
    a = 1664525
    c = 1013904223
    modulus = int64(math.Pow(2, 32))
    seed = 3820783849
    rng := GetRandomNumberGenerator(a, c, modulus)
    rng.SetSeed(seed)
    
    results := []int64{1246739764, 1106396931, 223527046, 2943229485, 4151404968}
    for _, expected := range results {
        output := rng.NextInt()
        if output != expected {
            test.Errorf("NextInt() == (%d), %d expected",
                output,
                expected)
        }
    }
}

func TestRandomNumberGenerator_NextFloat(test *testing.T) {
    var seed, a, c, modulus int64
    a = 1664525
    c = 1013904223
    modulus = int64(math.Pow(2, 32))
    seed = 4151404968
    rng := GetRandomNumberGenerator(a, c, modulus)
    rng.SetSeed(seed)
    
    results := []float64{
        0.3074074329342693,
        0.5933728874661028,
        0.2415774876717478,
        0.0037347888574004173,
        0.8854908372741193,
    }
    for _, expected := range results {
        output := rng.NextFloat()
        if output != expected {
            test.Errorf("NextFloat() == (%f), %f expected",
                output,
                expected)
        }
    }
}