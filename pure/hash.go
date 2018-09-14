package pure

type Hash struct {
	size    int
	table   []int
	visited map[int]int
}

func NewCuckooHash(size int) *Hash {
	return &Hash{
		size:  size,
		table: make([]int, size),
	}
}

func (h *Hash) Insert(v int) bool {
	idx := h.h1(v)

	// stuck in a cycle
	if h.visited[v] > 5 {
		return false
	}

	if h.table[idx] == 0 {
		h.table[idx] = v
		h.visited[v]++
		return true
	}

	idx = h.h2(v)
	if h.table[idx] == 0 {
		h.table[idx] = v
		h.visited[v]++
		return true
	}

	temp := h.table[idx]
	h.table[idx] = v
	h.visited[v]++
	return h.Insert(temp)
}

func (h *Hash) Delete(v int) bool {
	idx := h.h1(v)

	if h.table[idx] == 0 {
		return false
	}

	if h.table[idx] == v {
		h.table[idx] = 0
		return true
	}

	idx = h.h2(v)
	if h.table[idx] == 0 {
		return false
	}

	if h.table[idx] == v {
		h.table[idx] = 0
		return true
	}

	return false
}

func (h *Hash) Find(v int) bool {
	return h.table[h.h1(v)] == v || h.table[h.h2(v)] == v
}

func (h *Hash) h1(v int) int {
	return v * 3 % h.size
}

func (h *Hash) h2(v int) int {
	return v + 8%h.size
}
