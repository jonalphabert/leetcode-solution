import (
    "container/heap"
)

type ClassInfo struct {
    pass, total int
}

type MaxHeap []ClassInfo

// Nilai yang akan diurutkan berdasarkan penambahan pass ratio satu siswa
// gain = (pass+1)/(total+1) - pass/total
func gain(c ClassInfo) float64 {
    pass, total := float64(c.pass), float64(c.total)
    return (pass+1)/(total+1) - pass/total
}

func (h MaxHeap) Len() int {return len(h)}
func (h MaxHeap) Less(i,j int) bool {return gain(h[i]) > gain(h[j])}
func (h MaxHeap) Swap(i,j int) {
    h[i], h[j] = h[j], h[i]
}
func (h *MaxHeap) Push(x interface{}){
    *h = append(*h, x.(ClassInfo))
}
func (h *MaxHeap) Pop() interface{}{
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0:n-1]
    return x
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
    n := len(classes)
    heapGain := &MaxHeap{}

    // Inisiasi awal heap
    for _, class := range classes {
        heap.Push(heapGain, ClassInfo{class[0], class[1]})
    }

    // Iterasi sebanyak extraStudent
    // Ambil gain tertinggi lalu update nilainya
    for extraStudents > 0 {
        classUpdate := heap.Pop(heapGain).(ClassInfo)
        classUpdate.pass++
        classUpdate.total++
        heap.Push(heapGain, classUpdate)
        extraStudents--
    }

    // Iterasi semua data di heap lalu ambil rata" nya
    sum := 0.0
    for _, class := range *heapGain {
        sum += float64(class.pass)/float64(class.total)
    }

    return sum/float64(n)
}