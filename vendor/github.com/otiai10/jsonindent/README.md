JSON Indent for Go Just for Enjoying Code Golf :golf:

# You Ain't Gonna Need This

Because this is just a shorthand for [`json.Encoder.SetIndet`](https://godoc.org/encoding/json#Encoder.SetIndent).

```go
jsonindent.NewEncoder(w).Encode(v)
// is equivalent to
// encoder := json.NewEncoder(w)
// encoder.SetIndent("", "\t")
// encoder.Encode(v)
```

[![Build Status](https://travis-ci.org/otiai10/jsonindent.svg?branch=master)](https://travis-ci.org/otiai10/jsonindent)
[![GoDoc](https://godoc.org/github.com/otiai10/jsonindent?status.svg)](https://godoc.org/github.com/otiai10/jsonindent)
