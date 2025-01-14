# Go Race Condition Example

This repository demonstrates a common race condition in Go programs involving concurrent access to a shared counter without proper synchronization. The `bug.go` file contains the buggy code, while `bugSolution.go` shows the corrected version using a mutex or atomic operations.

## Problem
The program attempts to increment a counter 1000 times using goroutines.  Without proper synchronization mechanisms, multiple goroutines can access and modify the counter simultaneously, leading to data races and incorrect results. The output is likely to be less than 1000.

## Solution
The corrected code in `bugSolution.go` uses a mutex to protect the counter from concurrent access.  This ensures that only one goroutine can modify the counter at a time, preventing race conditions and producing the expected result (1000). Alternatively, atomic operations offer a more efficient solution.
