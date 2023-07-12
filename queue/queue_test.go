package queue

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
			want: 3,
		},
		{
			set:  []interface{}{3.1, 1.2, 6.1, 5.7, 1.1},
			want: 3.1,
		},
		{
			set:  []interface{}{"e", "d", "c", "b", "a"},
			want: "e",
		},
	}

	for _, tc := range testCases {
		q := Init[interface{}]()
		for _, v := range tc.set {
			q.Push(v)
		}
		if q.Front() != tc.want {
			t.Errorf("want =  %v, but result = %v", tc.want, q.Front())
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
			want: []interface{}{3, 1, 6, 5, 1},
		},
		{
			set:  []interface{}{1.1, 2.2, 3.3, 4.4, 5.5},
			want: []interface{}{1.1, 2.2, 3.3, 4.4, 5.5},
		},
		{
			set:  []interface{}{"a", "b", "c", "d", "e"},
			want: []interface{}{"a", "b", "c", "d", "e"},
		},
	}

	for _, tc := range testCases {
		q := Init[interface{}]()
		for _, v := range tc.set {
			q.Push(v)
		}

		popItems := []interface{}{}
		for !q.Empty() {
			popItems = append(popItems, q.Front())
			q.Pop()
		}
		if !reflect.DeepEqual(popItems, tc.want) {
			t.Errorf("want =  %v, but result = %v", tc.want, popItems)
		}
	}
}
