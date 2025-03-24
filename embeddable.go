package embeddable

import "fmt"

// Marshaller is an interface to unmarshal data and extract it to destination.
type Marshaller interface {
	Unmarshal(data []byte, dest any) error
}

// Extract extracts data from marshalled bytes via specified T and Marshaller. It returns
// a typed T object if there was no error.
func Extract[T comparable](m Marshaller, d []byte) (T, error) {
	var zero T
	if err := m.Unmarshal(d, &zero); err != nil {
		return zero, fmt.Errorf("embeddable: %w", err)
	}
	return zero, nil
}

// MustExtract doing same job as Extract but don't return an error. It will also panic
// if there will be one.
func MustExtract[T comparable](m Marshaller, d []byte) T {
	res, err := Extract[T](m, d)
	if err != nil {
		panic(err.Error())
	}
	return res
}
