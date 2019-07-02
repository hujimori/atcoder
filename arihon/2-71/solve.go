package main

import "fmt"

type Heap struct {
	n  []int
	sz int
}

func (h *Heap) push(n int) {
	// 自分のノードの番号
	i := h.sz
	h.sz++

	for i > 0 {
		// 親のノードの番号
		p := (i - 1) / 2

		// 逆転していないから抜ける
		if h.n[p] <= n {
			break
		}
		// 親のノードの数字を下ろして、自分は上に
		h.n[i] = h.n[p]
		i = p
	}
	h.n = append(h.n, n)
	// h.n[i] = n
}

func (h *Heap) pop() int {
	// 最小値
	ret := h.n[0]

	// 根に持ってくる値
	h.sz--
	x := h.n[h.sz]

	// 根から下ろしていく
	i := 0
	for i*2+1 < h.sz {
		// 子同士を比較
		a := i*2 + 1
		b := i*2 + 2
		if b < h.sz && h.n[b] < h.n[a] {
			a = b
		}

		// もう逆転していないら終わり
		if h.n[a] >= x {
			break
		}

		// 子の数字を持ち上げる
		h.n[i] = h.n[a]
		i = a
	}

	h.n[i] = x

	return ret

}

func main() {
	h := Heap{}
	h.push(1)
	h.push(2)
	h.push(3)
	fmt.Println(h)

}
