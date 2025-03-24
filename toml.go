package embeddable

import "github.com/pelletier/go-toml"

// TOML ...
var TOML = tomlMarshaller{}

// tomlMarshaller is an implementation of Marshaller to toml.
type tomlMarshaller struct{}

// Unmarshal ...
func (t tomlMarshaller) Unmarshal(data []byte, dest any) error {
	return toml.Unmarshal(data, dest)
}
