package cache

// 参考这位大佬 https://books.halfrost.com/leetcode/ChapterThree/LRUCache/
type LRUCache struct {
	head, tail *Node
	Keys       map[int]*Node
	Cap        int
}

type Node struct {
	Prev, Next *Node
	Key, Val   int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{Keys: make(map[int]*Node), Cap: capacity}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Keys[key]; ok {
		// return val
		this.Remove(node)
		this.Add(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Keys[key]; ok {
		node.Val = value
		this.Remove(node)
		this.Add(node)
		return
	} else {
		node := &Node{Prev: nil, Next: nil, Key: key, Val: value}
		this.Add(node)
		this.Keys[key] = node
		if len(this.Keys) > this.Cap {
			delete(this.Keys, this.tail.Key)
			this.Remove(this.tail)
		}
	}
}

// remove node from doubly-linked list
func (this *LRUCache) Remove(node *Node) {
	if node == this.head {
		this.head = node.Next
		node.Next = nil
		return
	}
	if node == this.tail {
		this.tail = node.Prev
		node.Prev.Next = nil
		node.Prev = nil
		return
	}

	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

// add node to head
func (this *LRUCache) Add(node *Node) {
	node.Prev = nil
	node.Next = this.head
	if this.head != nil {
		this.head.Prev = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = node
		this.tail.Next = nil
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
