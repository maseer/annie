package yinyuetai

import (
	"testing"

	"github.com/maseer/annie/extractors/types"
	"github.com/maseer/annie/test"
)

func TestDownload(t *testing.T) {
	tests := []struct {
		name string
		args test.Args
	}{
		{
			name: "normal test",
			args: test.Args{
				URL:     "http://v.yinyuetai.com/video/3386385",
				Title:   "什么是爱/ What is Love",
				Size:    20028736,
				Quality: "流畅",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			New().Extract(tt.args.URL, types.Options{})
		})
	}
}
