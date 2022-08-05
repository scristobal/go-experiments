package decoder

import (
	"testing"
)

func Test(t *testing.T) {

	type testStruct struct {
		TestContent string `json:"test_content"`
	}

	var testByte = []byte(`{"test_content":"test"}`)

	var testDecodable = Decodable[testStruct](testByte)

	testStructOut, err := testDecodable.Decode()

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if testStructOut.TestContent != "test" {
		t.Errorf("did not decode object got %s", testStructOut.TestContent)
	}
}
