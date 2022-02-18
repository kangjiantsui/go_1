package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTimeUtilImpl_Timestamp2Time(t *testing.T) {
	assert.Equal(t, TimeUtilImpl{}.Timestamp2Time(1645173808).String(), "2022-02-18 16:43:28 +0800 CST")
}
