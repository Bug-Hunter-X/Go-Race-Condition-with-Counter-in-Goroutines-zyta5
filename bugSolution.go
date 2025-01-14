```go
package main

import (

        "fmt"
        "sync"
        "sync/atomic"
)

func main() {
        var wg sync.WaitGroup
        var count int64 // Use int64 for atomic operations
        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func() {
                        defer wg.Done()
                        atomic.AddInt64(&count, 1) // Use atomic operation
                }()
        }
        wg.Wait()
        fmt.Println(count)
}
```