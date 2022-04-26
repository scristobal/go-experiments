package decoder

import "encoding/json"

type Raw[T any] []byte

func (in Raw[T]) Decode() (T, error) {
	var out T

	err := json.Unmarshal(in, &out)

	return out, err

}

func Decodable[T any](data []byte) Raw[T] {
	var out Raw[T] = data
	return out
}
