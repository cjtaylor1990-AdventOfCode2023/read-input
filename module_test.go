package readinput

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testFile   = "resources/test.txt"
	noSuchFile = "resources/nosuchfile.txt"
)

func TestCanReadATextFile(t *testing.T) {
	actual, err := ReadInput(testFile)
	expected := []string{"Hello, World!", "Goodnight, World!"}

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestReturnsErrorWhenNoSuchFileExists(t *testing.T) {
	actual, err := ReadInput(noSuchFile)

	assert.Error(t, err)
	assert.ErrorContains(t, err, "there was a problem reading the file: ")
	assert.Nil(t, actual)
}
