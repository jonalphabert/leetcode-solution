// Double linked list
type Node struct {
    key, value int
    prev, next *Node
}

type LRUCache struct {
    capacity int
    cache map[int]*Node
    head *Node
    tail *Node
}


func Constructor(capacity int) LRUCache {
    head, tail := &Node{}, &Node{}
    head.next = tail
    tail.prev = head
    return LRUCache{
        capacity: capacity,
        cache : make(map[int]*Node),
        head : head,
        tail : tail,
    }
}

func (this *LRUCache) remove(node *Node) {
    prev, next := node.prev, node.next
    prev.next = next
    next.prev = prev
}

func (this *LRUCache) insert(node *Node){
    node.next = this.head.next
    node.prev = this.head
    this.head.next.prev = node
    this.head.next = node
}

func (this *LRUCache) Get(key int) int {
    if node, ok := this.cache[key]; ok {
        this.remove(node)
        this.insert(node)
        return node.value
    }

    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    // Update kalau key sudah ada
    if node, ok := this.cache[key]; ok {
        this.remove(node)
        node.value = value
        this.insert(node)
        return
    }

    // check panjang cache apakah sudah mencapai capacity
    if len(this.cache) == this.capacity {
        lru := this.tail.prev
        this.remove(lru)
        delete(this.cache, lru.key)
    }

    newNode := &Node{
        key: key,
        value: value,
    }
    this.cache[key] = newNode
    this.insert(newNode)
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */