package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPrefixToLisp(t *testing.T) {
	tests := []struct {
		name,
		input,
		expected,
		errMsg string
	}{
		{name: "3 operators", input: "* ^ 2 3 - 5 1", expected: "(* (pow 2 3) (- 5 1))", errMsg: ""},
		{name: "7 operators", input: "/ * + ^ 7 2 3 - 8 2 - + 10 5 1", expected: "(/ (* (+ (pow 7 2) 3) (- 8 2)) (- (+ 10 5) 1))", errMsg: ""},
		{name: "empty", input: " ", expected: "", errMsg: "invalid input: empty expression"},
		{name: "few operands", input: "- 3", expected: "", errMsg: "incorrect input: invalid expression"},
		{name: "few operators", input: "- 10 2 3", expected: "", errMsg: "incorrect input: invalid expression"},
		{name: "invalid characters", input: "- d 10 2 3", expected: "", errMsg: "incorrect input: invalid character in expression"},
	}

	for _, tt := range tests {
		res, err := PrefixToLisp(tt.input)

		if tt.errMsg == "" {
			require.Nil(t, err, "Unexpected error for input: %q", tt.input)
			assert.Equal(t, tt.expected, res, "expected: %q, got: %q", tt.expected, res)
		} else {
			require.Error(t, err, "Expected error for input: %q", tt.input)
			assert.EqualError(t, err, tt.errMsg, "expected: %q, got: %q", tt.errMsg, err)
		}
	}
}

func ExamplePrefixToLisp() {
	res, _ := PrefixToLisp("* + 2 2 ^ 4 5")
	fmt.Println(res)

	// Output:
	// (* (+ 2 2) (pow 4 5))
}
