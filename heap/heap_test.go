package heap

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
			set:  []interface{}{1, 6, 7, 4, 2},
			want: 7,
		},
		{
			set:  []interface{}{8.4, 6.3, 1.2, 7.3, 10.0},
			want: 10.0,
		},
		{
			set:  []interface{}{"a", "d", "c", "b", "a"},
			want: "d",
		},
	}

	for _, tc := range testCases {
		h := Init(func(i, j interface{}) bool {
			switch i.(type) {
			case int:
				return i.(int) <= j.(int)
			case float64:
				return i.(float64) <= j.(float64)
			case string:
				return i.(string) <= j.(string)
			}
			panic("unknown type")
		})

		for _, v := range tc.set {
			h.Push(v)
		}
		if h.Top() != tc.want {
			t.Errorf("want =  %v, but result = %v", tc.want, h.Top())
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
			want: []interface{}{1, 1, 3, 5, 6},
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
		h := Init(func(i, j interface{}) bool {
			switch i.(type) {
			case int:
				return i.(int) >= j.(int)
			case float64:
				return i.(float64) >= j.(float64)
			case string:
				return i.(string) >= j.(string)
			}
			panic("unknown type")
		})

		for _, v := range tc.set {
			h.Push(v)
		}

		popItems := []interface{}{}
		for !h.Empty() {
			popItems = append(popItems, h.Top())
			h.Pop()
		}
		if !reflect.DeepEqual(popItems, tc.want) {
			t.Errorf("want =  %v, but result = %v", tc.want, popItems)
		}
	}
}
