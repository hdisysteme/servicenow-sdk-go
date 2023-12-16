package attachmentapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt_UnmarshalJSON_Success(t *testing.T) {

	rawJSON := []byte("\"1\"")

	intVal := Int(0)
	err := (&intVal).UnmarshalJSON(rawJSON)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, intVal, Int(1))
}

func TestInt_UnmarshalJSON_Failed(t *testing.T) {

	rawJSON := []byte("\"s\"")

	intVal := Int(0)
	err := (&intVal).UnmarshalJSON(rawJSON)
	assert.Error(t, err)
}
