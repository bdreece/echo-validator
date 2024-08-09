[![.github/workflows/build.yml](https://github.com/bdreece/echo-validator/actions/workflows/build.yml/badge.svg)](https://github.com/bdreece/echo-validator/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/bdreece/echo-validator.svg)](https://pkg.go.dev/github.com/bdreece/echo-validator)

# echo-validator

A basic implementation of [echo.Validator] using [github.com/goplayground/validator/v10].

## Usage

The [echo-validator] package provides an easy-to-use global instance of `Validator` that can
be configured with `echo.Echo` as follows:

```go
package main

import (
    echovalidator "github.com/bdreece/echo-validator"
    "github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()
    e.Validator = echovalidator.Default

    /* ... */
}
```

## License

MIT License

Copyright (c) 2024 Brian Reece

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

[github.com/goplayground/validator/v10]: https://github.com/goplayground/validator
[echo.Validator]: https://pkg.go.dev/github.com/labstack/echo/v4#Validator
[echo-validator]: https://pkg.go.dev/github.com/bdreec/echo-validator
