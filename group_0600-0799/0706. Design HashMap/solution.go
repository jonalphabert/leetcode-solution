type Bucket struct {
    key, value int
}

type MyHashMap struct {
    capacity     int
    mapStructure [][]Bucket
}

func Constructor() MyHashMap {
    capacity := 100003 // bilangan prima dekat 1e5
    myMap := make([][]Bucket, capacity)

    return MyHashMap{
        capacity:     capacity,
        mapStructure: myMap,
    }
}

func (this *MyHashMap) hash(key int) int {
    return key % this.capacity
}

func (this *MyHashMap) Put(key int, value int) {
    idx := this.hash(key)

    for i, bucket := range this.mapStructure[idx] {
        if bucket.key == key {
            // update kalau sudah ada
            this.mapStructure[idx][i].value = value
            return
        }
    }

    // kalau belum ada, tambahkan
    this.mapStructure[idx] = append(this.mapStructure[idx], Bucket{key: key, value: value})
}

func (this *MyHashMap) Get(key int) int {
    idx := this.hash(key)

    for _, bucket := range this.mapStructure[idx] {
        if bucket.key == key {
            return bucket.value
        }
    }
    return -1
}

func (this *MyHashMap) Remove(key int) {
    idx := this.hash(key)

    buckets := this.mapStructure[idx]
    for i, bucket := range buckets {
        if bucket.key == key {
            // hapus elemen dengan slicing
            this.mapStructure[idx] = append(buckets[:i], buckets[i+1:]...)
            return
        }
    }
}



/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */