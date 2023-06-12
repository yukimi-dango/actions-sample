package util_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	util "github.com/yukimi-dango/actions-sample/lib"
)

func TestMax(t *testing.T) {

	tcs := []struct {
		a    int64
		b    int64
		want int64
	}{
		{
			a:    1,
			b:    2,
			want: 2,
		},
		{
			a:    1,
			b:    1,
			want: 1,
		},
		{
			a:    -1,
			b:    1,
			want: 1,
		},
	}

	for _, tc := range tcs {
		got := util.Max(tc.a, tc.b)
		assert.Equal(t, tc.want, got)
	}
}
