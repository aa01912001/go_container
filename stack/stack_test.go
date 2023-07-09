package stack

import "testing"

func TestPush(t *testing.T) {
	type testCase[T any] struct {
		set  []T
		want T
	}

	testCases := []testCase[interface{}]{
		{
			set:  []interface{}{3, 1, 6, 5, 1},
			want: 1,
		},
		{
			set:  []interface{}{3.1, 1.2, 6.1, 5.7, 1.1},
			want: 1.1,
		},
		{
			set:  []interface{}{"e", "d", "c", "b", "a"},
			want: "a",
		},
	}

	stk := Init[interface{}]()
	for _, tc := range testCases {
		for _, v := range tc.set {
			stk.Push(v)
		}
		if stk.Top() != tc.want {
			t.Errorf("set value =  %v, but want value = %v", tc.set, tc.want)
		}
	}
}
