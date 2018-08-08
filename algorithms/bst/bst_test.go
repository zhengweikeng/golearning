package bst

import (
	"testing"
)

func TestBst_Get(t *testing.T) {
	tree := BST{
		Key: "S",
		Value: &Data{
			Value: "hello S",
		},
		Left: &BST{
			Key: "E",
			Value: &Data{
				Value: "hello E",
			},
			Left: &BST{
				Key: "A",
				Value: &Data{
					Value: "hello A",
				},
				Right: &BST{
					Key: "C",
					Value: &Data{
						Value: "hello C",
					},
				},
			},
			Right: &BST{
				Key: "R",
				Value: &Data{
					Value: "hello R",
				},
				Left: &BST{
					Key: "H",
					Value: &Data{
						Value: "hello H",
					},
					Right: &BST{
						Key: "M",
						Value: &Data{
							Value: "hello M",
						},
					},
				},
			},
		},
		Right: &BST{
			Key: "X",
			Value: &Data{
				Value: "hello X",
			},
		},
	}

	value := tree.Get("R")
	if value == nil {
		t.Error("R not found")
	}
	if value.Value != "hello R" {
		t.Error("R value is not hello R")
	}

	value = tree.Get("T")
	if value != nil {
		t.Error("T was found")
	}
}

func TestBST_Size(t *testing.T) {
	tree := &BST{
		Key: "S",
		Value: &Data{
			Value: "hello S",
		},
		N: 1,
	}
	if tree.Size() != 1 {
		t.Error("size shoule be 1")
	}
}

func TestBST_Put(t *testing.T) {
	tree := &BST{
		Key: "S",
		Value: &Data{
			Value: "hello S",
		},
		N: 1,
	}

	tree.Put("E", &Data{
		Value: "hello E",
	})
	if tree.Left.Value.Value != "hello E" {
		t.Errorf("Key = %s should be \"should be E\"\n", "E")
	}
	if tree.Size() != 2 {
		t.Errorf("size shoule be 2, but got %d", tree.Size())
	}
	if tree.Left.Size() != 1 {
		t.Errorf("size shoule be 1, but got %d", tree.Left.Size())
	}

	tree.Put("R", &Data{
		Value: "hello R",
	})
	if tree.Left.Right.Value.Value != "hello R" {
		t.Errorf("Key = %s should be \"should be R\"\n", "R")
	}
	if tree.Size() != 3 {
		t.Errorf("size shoule be 2, but got %d", tree.Size())
	}
}
