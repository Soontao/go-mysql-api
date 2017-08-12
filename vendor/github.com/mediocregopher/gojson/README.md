# gojson

A copy of go's `encoding/json` library which doesn't base64 encode/decode your
byte slices.

You can find the docs [here][docs]. You'll notice they're basically verbatum the
`encoding/json` docs.

# Installation

`go get github.com/mediocregopher/gojson`

or if you're using [goat][goat]:

```yaml
    - loc: https://github.com/mediocregopher/gojson.git
      type: git
      ref: v0.0.0
      path: github.com/mediocregopher/gojson
```

# What is it?

This is the exact same as the standard encoding/json package (as of Dec 14,
2013, go 1.2), except that it doesn't have the baffling behavior of
automatically treating []byte fields as base64 strings. So now you can have the
following code work the way you would expect it to:

```go
package main

import (
    "github.com/mediocregopher/gojson"
    "log"
)

type Wat struct {
    A, B []byte
}

func main() {
    b := []byte(`{"A":"foo","B":"bar"}`)
    w := &Wat{}

    err := gojson.Unmarshal(b,w)
    log.Println(string(w.A), string(w.B), err)
}
```

Whereas before the `A` and `B` fields would have been base64 decoded before
being put in the byte slices (which makes no sense and isn't the common case, so
why is it the behavior in the main package? beats me....).

[docs]: http://godoc.org/github.com/mediocregopher/gojson
[goat]: https://github.com/mediocregopher/goat
