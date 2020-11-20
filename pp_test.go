package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrettyPrint(t *testing.T) {
	t.Run("Commandline Argument", func(t *testing.T) {
		str := "{\"key1\":\"val1\",\"key2\":\"val2\",\"key3\":123,\"arr\":[101,201,303,400]}"
		buf := []byte(str)

		res, err := PrettyPrint(buf, 2)

		assert.Nil(t, err)
		assert.Equal(t, `{
  "arr": [
    101,
    201,
    303,
    400
  ],
  "key1": "val1",
  "key2": "val2",
  "key3": 123
}`, *res)

		t.Logf("%s", *res)
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		str := "{key1:\"val1\",key2:\"val2\"}"
		buf := []byte(str)

		res, err := PrettyPrint(buf, 2)

		assert.Nil(t, res)
		assert.NotNil(t, err)

		t.Log(err)
	})
}

func ExamplePrettyPrintOutput() {
	str := "{\"key1\":\"val1\",\"key2\":\"val2\",\"key3\":123,\"arr\":[101,201,303,400]}"
	buf := []byte(str)

	PrettyPrintOutput(buf, 2)
	// Output:
	// {
	//   "arr": [
	//     101,
	//     201,
	//     303,
	//     400
	//   ],
	//   "key1": "val1",
	//   "key2": "val2",
	//   "key3": 123
	// }
}
