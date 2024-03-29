# poptest

[![Go](https://github.com/royge/poptest/actions/workflows/go.yml/badge.svg)](https://github.com/royge/poptest/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/royge/poptest)](https://goreportcard.com/report/github.com/royge/poptest)

A utility package that can be used when testing with gobuffalo pop connection.

_The primary goal of this package is to override the database name from the
database.yml programatically._

## How To Use

```
go get github.com/royge/poptest
```

Example:

```go
import (
	"fmt"
	"strings"

	"github.com/royge/poptest"

	pop "github.com/gobuffalo/pop/v6"
)

// ExampleDeets show the primary purpose of this package.
// You can see that I override the test-mydb database name from database.yml
// with random-db-name.
func ExampleDeets() {
	deets, err := poptest.Deets("test", func(d *pop.ConnectionDetails) error {
		fmt.Println("original:", d.URL)

		name := "random-db-name"
		d.URL = strings.Replace(d.URL, "test-mydb", name, 1)

		return nil
	})
	if err != nil {
		panic("don't panic")
	}

	fmt.Println("new:", deets.URL)

	// Output:
	// original: postgres://postgres:password@127.0.0.1:5432/test-mydb?sslmode=disable
	// new: postgres://postgres:password@127.0.0.1:5432/random-db-name?sslmode=disable
}
```
