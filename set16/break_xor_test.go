package set16

import (
	"reflect"
	"testing"
)

func TestTranspose(t *testing.T) {
	input := [][]byte{[]byte("hello"), []byte("world")}
	expectedOutput := [][]byte{[]byte("hw"), []byte("eo"), []byte("lr"), []byte("ll"), []byte("od")}
	output := transpose(input)

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("expected: %v, actual %v", expectedOutput, output)
	}
}
