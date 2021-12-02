package union_find

type UnionFindSimple struct {
	parent []int
}

func NewUnionFindSimple(num int) *UnionFindSimple {
	return &UnionFindSimple{parent: make([]int, num)}
}

func (u *UnionFindSimple) Find(x int) int {
	if u.parent[x] == x {
		return x
	}
	return u.Find(u.parent[x])
}

func (u *UnionFindSimple) Merge(x, y int) {
	x, y = u.Find(x), u.Find(y)
	if x == y {
		return
	}

	u.parent[x] = y
}
