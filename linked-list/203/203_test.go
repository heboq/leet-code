package main

import (
	"reflect"
	"testing"
)

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func Test_removeElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"Target is the first and only one node",
			args{&ListNode{6, nil}, 6},
			nil,
		},
		{
			"All nodes are targets",
			args{&ListNode{6, &ListNode{6, &ListNode{6, &ListNode{6, nil}}}}, 6},
			nil,
		},
		{
			"Empty slice",
			args{&ListNode{}, 1},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
