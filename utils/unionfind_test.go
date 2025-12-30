package utils

import "testing"

func newUF() *UnionFind[int] {
	return &UnionFind[int]{
		data: make(map[int]int),
		rank: make(map[int]int),
		size: make(map[int]int),
	}
}

func TestAddAndFind(t *testing.T) {
	uf := newUF()

	uf.Add(1)
	uf.Add(2)

	if uf.Find(1) != 1 {
		t.Fatalf("expected root of 1 to be 1")
	}
	if uf.Find(2) != 2 {
		t.Fatalf("expected root of 2 to be 2")
	}

	// Adding the same node again should be a no-op
	uf.Add(1)
	if uf.size[1] != 1 {
		t.Fatalf("expected size of node 1 to remain 1")
	}
}

func TestJoinAndFind(t *testing.T) {
	uf := newUF()

	uf.Add(1)
	uf.Add(2)
	uf.Add(3)

	if !uf.Join(1, 2) {
		t.Fatalf("expected Join(1,2) to return true")
	}

	root1 := uf.Find(1)
	root2 := uf.Find(2)

	if root1 != root2 {
		t.Fatalf("expected 1 and 2 to be in the same set")
	}

	// Joining already-connected nodes should return false
	if uf.Join(1, 2) {
		t.Fatalf("expected Join on same set to return false")
	}

	// Joining another node
	if !uf.Join(2, 3) {
		t.Fatalf("expected Join(2,3) to return true")
	}

	if uf.Find(3) != root1 {
		t.Fatalf("expected 3 to be connected to the same root as 1")
	}
}

func TestSizeAndRank(t *testing.T) {
	uf := newUF()

	uf.Add(1)
	uf.Add(2)
	uf.Add(3)
	uf.Add(4)

	uf.Join(1, 2)
	uf.Join(3, 4)
	uf.Join(1, 3)

	root := uf.Find(1)

	if uf.size[root] != 4 {
		t.Fatalf("expected size of the set to be 4, got %d", uf.size[root])
	}

	if uf.rank[root] < 2 {
		t.Fatalf("expected rank to increase after unions")
	}
}

