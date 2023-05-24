package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceChunk(t *testing.T) {
	src1 := []string{"1", "2", "3", "4", "5"}
	assert.Equal(t, [][]string{{"1", "2"}, {"3", "4"}, {"5"}}, SliceChunk(src1, 2))
	assert.Equal(t, [][]string{{"1", "2", "3", "4", "5"}}, SliceChunk(src1, 5))
	assert.Equal(t, [][]string{{"1", "2", "3", "4", "5"}}, SliceChunk(src1, 6))
	assert.Equal(t, [][]string{{"1"}, {"2"}, {"3"}, {"4"}, {"5"}}, SliceChunk(src1, 1))
	assert.Equal(t, [][]string{{"1"}, {"2"}, {"3"}, {"4"}, {"5"}}, SliceChunk(src1, 0))
	assert.Equal(t, [][]string{{"1"}, {"2"}, {"3"}, {"4"}, {"5"}}, SliceChunk(src1, -100))
	assert.Equal(t, [][]string{}, SliceChunk(nil, 5))
}
