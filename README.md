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

You can find this example with data in [this gist](https://gist.github.com/fabritsius/6c1f63563616a22119dcba7e43b5e929).

## Features

- use `Read(filename)` to parse a CSV file
- by default first line is used to name all the fields
- function `Read` can take additional functional options:
    - use `Index([slice, of, names])` to set custom names to each column
    - use `Skip(nrows)` to skip N rows from the beginning
    - use `Limit(nrows)` to limit amount of returned rows
    - use `Delimiter('rune')` to change default comma separator

Please see examples of use in [csvier_test.go](./csvier_test.go) file.

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.