package link

import (
	"testing"
)

func TestNewAndString(t *testing.T) {
	if New().String() != "" {
		t.Error("link.New() error")
	}
	if New(1).String() != "1" {
		t.Error("link.New(1) error")
	}
	if New(1, 2, 3).String() != "1->2->3" {
		t.Error("link.New(1,2,3) error")
	}
}

func TestNewAndStringV2(t *testing.T) {
	if New().StringV2() != "" {
		t.Error("link.New() error")
	}
	if New(1).StringV2() != "1" {
		t.Error("link.New(1) error")
	}
	if New(1, 2, 3).StringV2() != "1->2->3" {
		t.Error("link.New(1,2,3) error")
	}
}

func TestNewAndStringV3(t *testing.T) {
	if New().StringV3() != "" {
		t.Error("link.New() error")
	}
	if New(1).StringV3() != "1" {
		t.Error("link.New(1) error")
	}
	if New(1, 2, 3).StringV3() != "1->2->3" {
		t.Error("link.New(1,2,3) error")
	}
}

func TestCmp(t *testing.T) {
	if !New().Cmp(New()) || !New().Cmp(nil) {
		t.Error("error: New() != New()")
	}
	if !New(1, 2, 3).Cmp(New(1, 2, 3)) {
		t.Error("error: New(1,2,3) != New(1,2,3)")
	}

	l := New(1, 2, 3, 4, 5)
	if !l.Cmp(l) {
		t.Error("error: New(1,2,3,4,5) != self")
	}
}

func TestReverse(t *testing.T) {
	if New().Reverse().String() != "" {
		t.Error("error: New().Reverse()")
	}

	if New(1, 2, 3).Reverse().String() != "3->2->1" {
		t.Error("error: New(1,2,3).Reverse()")
	}
}

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = New(1, 2, 3).String()
		_ = New(1, 2, 3, 4, 5, 6, 7, 8, 9).String()
		_ = New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20).String()
	}
}

func BenchmarkStringV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = New(1, 2, 3).StringV2()
		_ = New(1, 2, 3, 4, 5, 6, 7, 8, 9).StringV2()
		_ = New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20).StringV2()
	}
}

func BenchmarkStringV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = New(1, 2, 3).StringV3()
		_ = New(1, 2, 3, 4, 5, 6, 7, 8, 9).StringV3()
		_ = New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20).StringV3()
	}
}
