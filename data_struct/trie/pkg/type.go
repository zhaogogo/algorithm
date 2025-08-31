package pkg

type trieNode struct {
	// 子节点列表.
	// trie树通过这一项，形成了父子节点一对多的映射关系，最终形成了多叉树的拓扑结构
	nexts []*trieNode

	// 记录通过当前节点的单词数量.
	// 以根节点到当前节点所形成字符串作为前缀的单词数量
	passCnt int

	// 标识是否存在单词以当前节点为结尾.
	// 存在以从根节点到当前节点所形成字符串作为内容的单词
	end bool

	value rune
}

type Trie struct {
	root *trieNode
}

// trie 中需要持有一个根节点 root 即可.
// 其中 root 是所有 trieNode 节点的始祖，其本身对应的字符为空.
func NewTrie() *Trie {
	return &Trie{
		root: &trieNode{
			value: rune('/'),
		},
	}
}

// 检索一个单词，判断其是否存在于trie树的处理流程
func (t *Trie) Search(word string) bool {
	// 查找目标节点，使得从根节点开始抵达目标节点沿路字符形成的字符串恰好等于 word
	node := t.search(word)
	return node != nil && node.end
}

func (t *Trie) search(target string) *trieNode {
	// 移动指针从根节点出发
	move := t.root
	// 依次遍历 target 中的每个字符
	for index, ch := range target {
		// 倘若 nexts 中不存在对应于这个字符的节点，说明该单词没插入过，返回 nil
		if move.value != ch {
			return nil
		}
		move =
	}
}
