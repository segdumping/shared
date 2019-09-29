package stack

import "testing"

func TestStack(t *testing.T) {
	var stack Stack
	stack.Push(1)
	stack.Push(2)

	t.Logf("stack top: %v", stack.Top())
	t.Logf("stack pop: %v", stack.Pop())
	t.Logf("stack empty: %v", stack.Empty())
}
