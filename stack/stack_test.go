package stack

import (
	"reflect"
	"testing"
)

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
			t.Errorf("want =  %v, but result = %v", tc.want, stk.Top())
		}
	}
}

func TestPop(t *testing.T) {
	type testCase[T any] struct {
		set  []T
		want []T
	}

	testCases := []testCase[interface{}]{
		{
			set:  []interface{}{3, 1, 6, 5, 1},
			want: []interface{}{1, 5, 6, 1, 3},
		},
		{
			set:  []interface{}{1.1, 2.2, 3.3, 4.4, 5.5},
			want: []interface{}{5.5, 4.4, 3.3, 2.2, 1.1},
		},
		{
			set:  []interface{}{"a", "b", "c", "d", "e"},
			want: []interface{}{"e", "d", "c", "b", "a"},
		},
	}

	stk := Init[interface{}]()
	for _, tc := range testCases {
		for _, v := range tc.set {
			stk.Push(v)
		}

		popItems := []interface{}{}
		for !stk.Empty() {
			popItems = append(popItems, stk.Top())
			stk.Pop()
		}
		if !reflect.DeepEqual(popItems, tc.want) {
			t.Errorf("want =  %v, but result = %v", tc.want, popItems)
		}
	}
}
