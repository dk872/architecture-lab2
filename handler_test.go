package lab2

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompute(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		errMsg   string
	}{
		{
			name:     "valid expression",
			input:    "* ^ 2 3 - 5 1",
			expected: "(* (pow 2 3) (- 5 1))\n",
		},
		{
			name:   "empty expression",
			input:  " ",
			errMsg: "empty expression",
		},
		{
			name:   "invalid expression",
			input:  "- 3",
			errMsg: "syntax error: incorrect input: invalid expression",
		},
		{
			name:   "invalid characters",
			input:  "- d 10 2 3",
			errMsg: "syntax error: incorrect input: invalid character in expression",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := &strings.Builder{}
			handler := ComputeHandler{
				Input:  strings.NewReader(tt.input),
				Output: output,
			}

			err := handler.Compute()

			if tt.errMsg == "" {
				require.NoError(t, err, "unexpected error for test: %s", tt.name)
				assert.Equal(t, tt.expected, output.String(), "incorrect output for test: %s", tt.name)
			} else {
				require.Error(t, err, "expected error for test: %s", tt.name)
				assert.EqualError(t, err, tt.errMsg, "wrong error for test: %s", tt.name)
			}
		})
	}
}
