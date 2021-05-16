package metadata

import (
	"bufio"
	"errors"
)

type Metadata struct {
	Frames []*Frame
}

var (
	ErrBadFrame = errors.New("each frame must contain two valid lines of metadata")
	ErrBadSequence = errors.New("input frames sequence is invalid")
)

func New(r *bufio.Reader) (m *Metadata, err error) {
	var (
		f *Frame
		fs []*Frame
	)

	// call ReadLine until EOL
	for err == nil {
		f, err = ParseSingleFrame(r)

		// bad parse
		if f == nil {
			break
		}

		// bad frame index
		if len(fs) != f.Index {
			return nil, ErrBadSequence
		}

		fs = append(fs, f)
	}

	return &Metadata{fs}, nil
}

func ParseSingleFrame(r *bufio.Reader) (*Frame, error) {

	frame, _, err := r.ReadLine()
	if err != nil {
		return nil, err
	}

	score, _, err := r.ReadLine()
	if err != nil {
		return nil, err
	}

	return NewFrame(string(frame), string(score))
}
