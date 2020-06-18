package datastruct

type Heap struct {
	N int
	Q int
	A []int
}

func NewHeap(n int) *Heap {
	return &Heap{N: n}
}

func (h *Heap) Insert(d int) {
	if h.Q >= h.N {
		return
	}
	h.Q++
	h.A = append(h.A, d)
	i := h.Q
	for i>>1 > 0 && h.A[i] > h.A[i>>1] {
		h.A[i], h.A[i>>1] = h.A[i>>1], h.A[i]
		i = i >> 1
	}
}
