package metadata

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func TestNewFrame(t *testing.T) {
	dat, err := ioutil.ReadFile("../fixtures/time.txt")
	assert.Nil(t, err)

	strs := strings.Split(string(dat), "\n")
	f, err := NewFrame(strs[2], strs[3])

	assert.Nil(t, err)
	assert.Equal(t, 1, f.Index)
	assert.Equal(t, 1867, f.Pts)
	assert.Equal(t, 1.867, f.PtsTime)
	assert.Equal(t, 0.002595, f.SceneScore)
}

func TestParseError(t *testing.T) {
	dat, _ := ioutil.ReadFile("../fixtures/time.txt")
	strs := strings.Split(string(dat), "\n")

	// bad first line
	_, err := NewFrame("x", strs[1])
	assert.ErrorIs(t, err, ParseError)

	_, err = NewFrame("frame:0 pts:1825 pts_time:1.8.2.5", strs[1])
	assert.Error(t, err)

	// bad second line
	_, err = NewFrame(strs[2], "x")
	assert.ErrorIs(t, err, ParseError)

	_, err = NewFrame(strs[2], "lavfi.scene_score=0.011.56.8")
	assert.Error(t, err)
}