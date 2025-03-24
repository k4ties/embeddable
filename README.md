# Embeddable

Embeddable is package that allows you to embed any struct to file, with specified marshal format (json, toml, yaml)
It is very helpful for small configs.

### For example:
```go
//go:embed filename.json 
var b []byte
var conf = embeddable.MustExtract[map[string]any](embeddable.JSON, b)
// do something with conf
``` 

Examples are located in the `example` folder.
