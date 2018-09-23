package otp

import "fmt"
import "testing"
import "github.com/ziliwesley/serious-cryptography/common"

func ExampleDecrypt() {
    fmt.Println(Decrypt("3c3b6eed4f", "545e028120"))
    // Output: hello
}

func TestDecrypt(test *testing.T) {
    fixtures := []common.Fixture {
        { Input: "07ad19a9d6", Expected: "hello" },
        { Input: "18a707a9dd", Expected: "world" },
        { Input: "1ba118a0ca", Expected: "times" },
    }
    for _, c := range fixtures {
        output := Decrypt(c.Input, "6fc875c5b9")
        if output != c.Expected {
            test.Errorf("Decrypt(%q, 2) == (%q), %q expected",
                c.Input,
                output,
                c.Expected)
        }
    }
}