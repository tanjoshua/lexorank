package lexorank_test

import (
	"testing"

	"github.com/tanjoshua/lexorank"
)

func TestBetween(t *testing.T) {
	tests := []struct {
		name  string
		left  lexorank.LexoRank
		right lexorank.LexoRank
	}{
		{
			name:  "min and max",
			left:  lexorank.Min(),
			right: lexorank.Max(),
		},
		{
			name:  "no space",
			left:  lexorank.LexoRank{Bucket: lexorank.DEFAULT_BUCKET, Value: "a"},
			right: lexorank.LexoRank{Bucket: lexorank.DEFAULT_BUCKET, Value: "b"},
		},
		{
			name:  "no space 2",
			left:  lexorank.LexoRank{Bucket: lexorank.DEFAULT_BUCKET, Value: "aaa"},
			right: lexorank.LexoRank{Bucket: lexorank.DEFAULT_BUCKET, Value: "aan"},
		},
		{
			name:  "varying sizes",
			left:  lexorank.LexoRank{Bucket: lexorank.DEFAULT_BUCKET, Value: "aaaaaaa"},
			right: lexorank.LexoRank{Bucket: lexorank.DEFAULT_BUCKET, Value: "b"},
		},
		{
			name:  "varying sizes 2",
			left:  lexorank.LexoRank{Bucket: lexorank.DEFAULT_BUCKET, Value: "a"},
			right: lexorank.LexoRank{Bucket: lexorank.DEFAULT_BUCKET, Value: "bbaaaa"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			between, err := lexorank.Between(&tt.left, &tt.right)
			if err != nil {
				t.Error(err)
			}

			if between.String() <= tt.left.String() || between.String() >= tt.right.String() {
				t.Errorf("Between() should be between min and max")
			}

			// Check that Between is between min and max
			if tt.left.String() > between.String() {
				t.Errorf("min should be less than Between()")
			}
			if between.String() > tt.right.String() {
				t.Errorf("Between() should be less than max")
			}
			t.Logf("Between(%s, %s) = %s", tt.left, tt.right, between)
		})
	}
}
