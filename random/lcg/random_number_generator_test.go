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

func ExampleRandomNumberGenerator_RandomBytes() {
    var seed, a, c, modulus int64
    a = 1664525
    c = 1013904223
    modulus = int64(math.Pow(2, 32))
    seed = 2601234783
    rng := GetRandomNumberGenerator(a, c, modulus)
    rng.SetSeed(seed)
    fmt.Printf("%x\n", rng.RandomBytes(20))
    // Output: 32cff916e9648e2b34e99f0903838c3686cc0fc9
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

func TestRandomNumberGenerator_RandomBytes(test *testing.T) {
    var seed, a, c, modulus int64
    a = 1664525
    c = 1013904223
    modulus = int64(math.Pow(2, 32))
    seed = 1538749367
    rng := GetRandomNumberGenerator(a, c, modulus)
    rng.SetSeed(seed)

    results := []byte{
        170, 189, 185, 73, 1, 81, 9, 130, 108, 118,
    }

    output := rng.RandomBytes(10)

    for index, expected := range results {
        if output[index] != expected {
            test.Errorf("RandomBytes()[%d] == (%d), %d expected",
                index,
                output[index],
                expected)
        }
    }
}