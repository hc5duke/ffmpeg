package metadata

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestParseSingleFrame(t *testing.T) {
	data := "frame:0    pts:1825    pts_time:1.825\nlavfi.scene_score=0.011568"
	br := bytes.NewReader([]byte(data))
	r := bufio.NewReader(br)

	f, err := ParseSingleFrame(r)
	assert.Nil(t, err)
	assert.NotNil(t, f)
}

func TestParseSingleFrameNotEnoughData(t *testing.T) {
	data := "frame:0    pts:1825    pts_time:1.825\n"
	br := bytes.NewReader([]byte(data))
	r := bufio.NewReader(br)

	f, err := ParseSingleFrame(r)
	assert.Equal(t, err, io.EOF)
	assert.Nil(t, f)
}

func TestParseSingleFrameBadData(t *testing.T) {
	data := "frame:0    pts:1825    pts_time:1.825\nsomething"
	br := bytes.NewReader([]byte(data))
	r := bufio.NewReader(br)

	f, err := ParseSingleFrame(r)
	assert.Equal(t, err, ErrBadFrame)
	assert.Nil(t, f)
}

func TestNewMetadata(t *testing.T) {
	filename := "./fixtures/time.txt"
	f, _ := os.Open(filename)
	r := bufio.NewReader(f)

	m, e := New(r)
	assert.Nil(t , e)
	assert.NotNil(t , m)
}