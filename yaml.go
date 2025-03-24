package embeddable

import "gopkg.in/yaml.v3"

// YAML ...
var YAML = yamlMarshaller{}

// yamlMarshaller is an implementation of Marshaller to yaml.
type yamlMarshaller struct{}

// Unmarshal ...
func (y yamlMarshaller) Unmarshal(data []byte, dest any) error {
	return yaml.Unmarshal(data, dest)
}
