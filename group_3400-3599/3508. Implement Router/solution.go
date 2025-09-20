type Packet struct {
	source      int
	destination int
	timestamp   int
}

type Router struct {
	memoryLimit int
	queue       []Packet                  
	seen        map[[3]int]bool           
	destMap     map[int][]int             
}

func Constructor(memoryLimit int) Router {
	return Router{
		memoryLimit: memoryLimit,
		queue:       make([]Packet, 0),
		seen:        make(map[[3]int]bool),
		destMap:     make(map[int][]int),
	}
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
	key := [3]int{source, destination, timestamp}
	if this.seen[key] {
		return false
	}

	if len(this.queue) == this.memoryLimit {
		oldest := this.queue[0]
		this.queue = this.queue[1:]
		delete(this.seen, [3]int{oldest.source, oldest.destination, oldest.timestamp})

		tsArr := this.destMap[oldest.destination]
		idx := sort.SearchInts(tsArr, oldest.timestamp)
		if idx < len(tsArr) && tsArr[idx] == oldest.timestamp {
			this.destMap[oldest.destination] = append(tsArr[:idx], tsArr[idx+1:]...)
		}
	}

	p := Packet{source, destination, timestamp}
	this.queue = append(this.queue, p)
	this.seen[key] = true
	this.destMap[destination] = append(this.destMap[destination], timestamp)

	return true
}

func (this *Router) ForwardPacket() []int {
	if len(this.queue) == 0 {
		return []int{}
	}

	p := this.queue[0]
	this.queue = this.queue[1:]
	delete(this.seen, [3]int{p.source, p.destination, p.timestamp})

	tsArr := this.destMap[p.destination]
	idx := sort.SearchInts(tsArr, p.timestamp)
	if idx < len(tsArr) && tsArr[idx] == p.timestamp {
		this.destMap[p.destination] = append(tsArr[:idx], tsArr[idx+1:]...)
	}

	return []int{p.source, p.destination, p.timestamp}
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
	tsArr, ok := this.destMap[destination]
	if !ok || len(tsArr) == 0 {
		return 0
	}

	left := sort.Search(len(tsArr), func(i int) bool { return tsArr[i] >= startTime })
	right := sort.Search(len(tsArr), func(i int) bool { return tsArr[i] > endTime })

	return right - left
}