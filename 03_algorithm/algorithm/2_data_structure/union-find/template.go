package union_find

// UnionFindTemplate 使用路径压缩和按秩合并
type UnionFindTemplate struct {
	fa   []int
	rank []int
}

func NewUnionFindTemplate(num int) *UnionFindTemplate {
	fa := make([]int, num)
	rank := make([]int, num)
	for i := 0; i < num; i++ {
		fa[i] = i
		rank[i] = 1
	}
	return &UnionFindTemplate{
		fa:   fa,
		rank: rank,
	}
}

func (u *UnionFindTemplate) Find(x int) int {
	// 路径压缩
	if x != u.fa[x] {
		u.fa[x] = u.fa[u.Find(x)]
	}
	return u.fa[x]
}

func (u *UnionFindTemplate) Union(x, y int) {
	// 按秩合并
	x, y = u.Find(x), u.Find(y)
	if x != y {
		r1, r2 := u.rank[x], u.rank[y]
		if r1 < r2 {
			x, y = y, x
		}
		u.rank[x]++
		u.fa[y] = x
	}
}
