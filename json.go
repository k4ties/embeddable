package embeddable

import "encoding/json"

// JSON ...
var JSON = jsonMarshaller{}

// jsonMarshaller is an implementation of Marshaller to json.
type jsonMarshaller struct{}

// Unmarshal ...
func (j jsonMarshaller) Unmarshal(data []byte, dest any) error {
	return json.Unmarshal(data, dest)
}
