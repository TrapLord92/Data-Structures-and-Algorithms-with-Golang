As of Go 1.20, the `rand.Seed` function has not been deprecated. However, if you are using an alternative library or there are specific requirements in your context indicating that `rand.Seed` should not be used directly, you can achieve similar functionality using `rand.New` and `rand.NewSource`. Here is how you can update your function to accommodate this:

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

// randomSequence generates a slice of random integers of length num.
func randomSequence(num int) []int {
    sequence := make([]int, num)
    source := rand.NewSource(time.Now().UnixNano())
    r := rand.New(source)
    for i := 0; i < num; i++ {
        // Generating a random integer between -999 and 999
        sequence[i] = r.Intn(1999) - 999
    }
    return sequence
}

func main() {
    num := 10
    sequence := randomSequence(num)
    fmt.Println("Random sequence:", sequence)
}
```

### Explanation of Changes:

1. **Using `rand.NewSource` and `rand.New`**:
   - `rand.NewSource(time.Now().UnixNano())` creates a new random source seeded with the current time in nanoseconds.
   - `rand.New(source)` creates a new random number generator that uses the specified source.
   - `r.Intn(1999) - 999` generates a random number between `-999` and `999`.

This approach ensures that you can manage multiple random number generators independently, each with its own seed, if needed.