package main

type DSU struct {
	n   int
	p   []int
	cnt []int
}

func (dsu *DSU) init() {
	for i := 0; i < dsu.n; i++ {
		dsu.p[i] = i
	}
}

func (dsu *DSU) leader(v int) int {
	if dsu.p[v] == v {
		return v
	}
	dsu.p[v] = dsu.leader(dsu.p[v])
	return dsu.p[v]
}

func (dsu *DSU) unite(a int, b int) {
	a = dsu.leader(a)
	b = dsu.leader(b)
	if dsu.cnt[a] < dsu.cnt[b] {
		a, b = b, a
	}
	dsu.cnt[a] += dsu.cnt[b]
	dsu.p[b] = a
}

func (dsu *DSU) uniteOld(a int, b int) {
	dsu.p[a] = dsu.leader(b)
}
