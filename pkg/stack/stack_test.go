package stack_test

import (
	"testing"

	"github.com/phrkdll/monomatch/pkg/stack"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type testCase struct {
		name  string
		input []any
	}

	testCases := []testCase{
		{
			name:  "string array input",
			input: []any{"1", "2", "3"},
		},
		{
			name:  "int array input",
			input: []any{1, 2, 3},
		},
		{
			name:  "empty array input",
			input: []any{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stack := stack.New(tc.input)

			assert.Equal(t, stack.Len(), len(tc.input))
		})
	}
}

func TestPushTopPop(t *testing.T) {
	type testCase struct {
		name   string
		first  int
		second int
	}

	testCases := []testCase{
		{
			name:   "add 1 then 2",
			first:  1,
			second: 2,
		},
		{
			name:   "add 2 then 1",
			first:  2,
			second: 1,
		},
		{
			name:   "add 0 then 0",
			first:  0,
			second: 0,
		},
	}

	for _, tc := range testCases {
		testStack := stack.New([]int{})

		t.Run(tc.name, func(t *testing.T) {
			assert.Zero(t, testStack.Len())

			_, err := testStack.Top()
			assert.NotNil(t, err)

			testStack.Push(tc.first)
			testStack.Push(tc.second)

			top, err := testStack.Top()
			assert.Equal(t, &tc.second, top)

			testStack.Pop()

			top, err = testStack.Top()
			assert.Equal(t, &tc.first, top)
		})
	}
}
