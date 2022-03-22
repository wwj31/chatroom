package dirty

type (
	Trie struct {
		root *node
	}
	node struct {
		words   map[byte]*node
		prevEnd bool
	}
)

func New() *Trie {
	return &Trie{root: &node{}}
}

func (t *Trie) Add(str string) {
	next := t.root
	for i, s := range str {
		next = next.add(byte(s), i == len(str)-1)
	}
}

type pair struct {
	begin, end int
}

// Match 返回所有dst中，匹配到的词索引位置[begin,end)
func (t *Trie) Match(dst string) []pair {
	var (
		result []pair
		p      *pair
		offset int
	)
	if len(t.root.words) == 0 || len(dst) == 0 {
		return nil
	}

	for n := 0; n < len(dst); n++ {
		offset = n
		next := t.root
		for {
			// 匹配
			if next.prevEnd {
				p.end = offset
				result = append(result, *p)
				// 最长匹配完成
				if len(next.words) == 0 {
					p = nil
					break
				}
			}

			// 未匹配完成，目标串结束
			if offset == len(dst) {
				p = nil
				break
			}

			// 字符失配
			nd, ok := next.words[dst[offset]]
			if !ok {
				p = nil
				break
			}

			if p == nil {
				p = &pair{begin: n}
			}
			next = nd
			offset++
		}
	}
	return result
}

func (nd *node) add(b byte, end bool) *node {
	if nd.words == nil {
		nd.words = make(map[byte]*node, 1)
	}

	next, ok := nd.words[b]
	if !ok {
		next = &node{}
		nd.words[b] = next
	}
	if !next.prevEnd {
		next.prevEnd = end
	}
	return next
}
