package main

import (
	"reflect"
	"testing"
)

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name     string
		args     args
		wantHead *ListNode
	}{
		{
			"identical lists",
			args{&ListNode{1, &ListNode{2, &ListNode{3, nil}}}, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}},
			&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}},
		},
		{
			"only second node differ",
			args{&ListNode{1, &ListNode{2, &ListNode{4, nil}}}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}}},
			&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
		},
		{
			"first list is nil",
			args{nil, &ListNode{0, nil}},
			&ListNode{0, nil},
		},
		{
			"both lists are nil",
			args{nil, nil},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHead := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(gotHead, tt.wantHead) {
				t.Errorf("mergeTwoLists() = %v, want %v", gotHead, tt.wantHead)
			}
		})
	}
}
