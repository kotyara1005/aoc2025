package utils

type UnionFind[K comparable] struct {
	data map[K]K
	rank map[K]int
	size map[K]int
	length int
}

func NewUnionFind[K comparable]() UnionFind[K] {
	return UnionFind[K]{
		map[K]K{},
		map[K]int{},
		map[K]int{},
		0,
	}
}

func (uf *UnionFind[K]) Size(node K) int {
	return uf.size[node]
}

func (uf *UnionFind[K]) Len() int {
	return uf.length
}

func (uf *UnionFind[K]) Add(node K)  {
	_, prs := uf.data[node]
	if prs {
		return
	}
	uf.data[node] = node
	uf.rank[node] = 1
	uf.size[node] = 1
	uf.length += 1
}

func (uf *UnionFind[K]) Find(node K) K {
	parent := uf.data[node]
	if parent == node {
		return parent
	}
	uf.data[node] = uf.Find(parent)
	return uf.data[node]
}

func (uf *UnionFind[K]) Join(n1, n2 K) bool {
	n1 = uf.Find(n1)
	n2 = uf.Find(n2)

	if n1 == n2 {
		return false
	}

	r1 := uf.rank[n1]
	r2 := uf.rank[n2]
	
	if r1 < r2 {
		n1, n2 = n2, n1
	}

	uf.data[n2] = n1

	if r1 == r2 {
		uf.rank[n1] += 1
	}

	uf.size[n1] += uf.size[n2]
	uf.length -= 1

	return true
}

