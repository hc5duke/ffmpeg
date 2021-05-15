package metadata

import (
	"errors"
	"regexp"
	"strconv"
)

type Frame struct {
	Index      int
	Pts        int
	PtsTime    float64
	SceneScore float64
}

var (
	//frame:0    pts:1825    pts_time:1.825
	r1, _ = regexp.Compile(`frame:(\d+)\s+pts:(\d+)\s+pts_time:([\d.]+)\b`)

	//lavfi.scene_score=0.011568
	r2, _ = regexp.Compile(`scene_score=([\d.]+)\b`)

	ParseError = errors.New("Unexpected input string")
)

func NewFrame(s0 string, s1 string) (*Frame, error) {
	m0 := r1.FindStringSubmatch(s0)
	m1 := r2.FindStringSubmatch(s1)

	if len(m0) != 4 {
		return nil, ParseError
	}

	if len(m1) != 2 {
		return nil, ParseError
	}

	f, err := strconv.Atoi(m0[1])
	if err != nil {
		return nil, err
	}
	p, err := strconv.Atoi(m0[2])
	if err != nil {
		return nil, err
	}
	pt, err := strconv.ParseFloat(m0[3], 64)
	if err != nil {
		return nil, err
	}
	ss, err := strconv.ParseFloat(m1[1], 64)
	if err != nil {
		return nil, err
	}

	return &Frame{
		f,
		p,
		pt,
		ss,
	}, nil
}
