type Food struct {
    name   string
    rating int
}

type PriorityQueue []Food

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    if pq[i].rating == pq[j].rating {
        return pq[i].name < pq[j].name
    }
    return pq[i].rating > pq[j].rating
}

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
    *pq = append(*pq, x.(Food))
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[:n-1]
    return item
}

type FoodRatings struct {
    foodCuisine map[string]string   
    foodRating  map[string]int      
    cuisineHeap map[string]*PriorityQueue 
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
    fr := FoodRatings{
        foodCuisine: make(map[string]string),
        foodRating:  make(map[string]int),
        cuisineHeap: make(map[string]*PriorityQueue),
    }

    for i, food := range foods {
        cuisine := cuisines[i]
        rating := ratings[i]

        fr.foodCuisine[food] = cuisine
        fr.foodRating[food] = rating

        if fr.cuisineHeap[cuisine] == nil {
            pq := &PriorityQueue{}
            heap.Init(pq)
            fr.cuisineHeap[cuisine] = pq
        }
        heap.Push(fr.cuisineHeap[cuisine], Food{name: food, rating: rating})
    }

    return fr
}

func (fr *FoodRatings) ChangeRating(food string, newRating int) {
    cuisine := fr.foodCuisine[food]
    fr.foodRating[food] = newRating
    heap.Push(fr.cuisineHeap[cuisine], Food{name: food, rating: newRating})
}

func (fr *FoodRatings) HighestRated(cuisine string) string {
    pq := fr.cuisineHeap[cuisine]
    for {
        top := (*pq)[0]
        if fr.foodRating[top.name] == top.rating {
            return top.name
        }
        heap.Pop(pq) 
    }
}