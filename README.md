# csvier

This Go module simplifies reading of CSV files. Parsed data is represented as a slice of maps.

## Example

The most basic example:

```go
package main

import (
    "fmt"

    "github.com/fabritsius/csvier"
)

func main() {
    data, err := csvier.Read("data.csv")
    if err != nil {
        panic(err)
    }

    for _, r := range data {
        fmt.Printf("%s (age %s) believes that %s.\n", r["NAME"], r["AGE"], r["BELIEF"])
    }
}
```

Output:
```
Tommy (age 6) believes that dragons are real.
Steve (age 29) believes that you can always win with a clear conscience.
Bill (age 49) believes that the Earth is flat.
```

Read function can take optional parameters. You can find more examples in [csvier_test.go](./csvier_test.go) file. By default first row of the file is used for naming each field.