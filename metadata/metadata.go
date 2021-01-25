package metadata

import (
	"bufio"
	"errors"
	"regexp"
	"strconv"
)

type Metadata struct {
	Frames []*Frame
}

var (
	frameRegexp = regexp.MustCompile(`^\s*frame:(\d+)\s+pts:(\d+)\s+pts_time:([\d\.]+)\s*$`)
	scoreRegexp = regexp.MustCompile(`^\s*lavfi.scene_score=([\d\.]+)\s*$`)
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
	score, _, err := r.ReadLine()

	if err != nil {
		return nil, err
	}

	m1 := frameRegexp.FindAllSubmatch(frame, -1)
	m2 := scoreRegexp.FindAllSubmatch(score, -1)

	if len(m1) != 1 || len(m2) != 1 || len(m1[0]) != 4 || len(m2[0]) != 2 {
		return nil, ErrBadFrame
	}

	f, err := strconv.Atoi(string(m1[0][1]))
	if err != nil {
		return nil, err
	}

	p, err := strconv.Atoi(string(m1[0][2]))
	if err != nil {
		return nil, err
	}

	pt, err := strconv.ParseFloat(string(m1[0][3]), 64)
	if err != nil {
		return nil, err
	}

	s, err := strconv.ParseFloat(string(m2[0][1]), 64)
	if err != nil {
		return nil, err
	}

	line := &Frame{
		Index:      f,
		Pts:        p,
		PtsTime:    pt,
		SceneScore: s,
	}

	return line, nil
}
