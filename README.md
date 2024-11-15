# randdetect

This is a small go package which detects if a string consisting of English words looks random.

The implementation is based on ideas from [random-string-detector](https://github.com/mehmedkadric/random-string-detector).

# Usage

```
go get -u github.com/ccojocar/randdetect
```

```go
package main

import (
  "fmt"
  "github.com/ccojocar/randdetect"
)

func main() {
  s := "my string"
  if ok := randdetect.IsRandom(s); ok {
    fmt.Printf("%q: looks random", s)
  }
}

```
