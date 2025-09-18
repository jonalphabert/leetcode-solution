type Task struct{
    userID      int
    taskID      int
    priority    int
}

type TaskHeap []*Task

func (heap TaskHeap) Len() int {return len(heap)}
func (heap TaskHeap) Less(i,j int) bool {
    if heap[i].priority == heap[j].priority {
        return heap[i].taskID > heap[j].taskID
    }
    return heap[i].priority > heap[j].priority
}
func (heap TaskHeap) Swap(i,j int){heap[i],heap[j]=heap[j],heap[i]}
func (heap *TaskHeap) Push(x interface{}) {
    *heap = append(*heap, x.(*Task))
}
func (heap *TaskHeap) Pop() interface{} {
    old := *heap
    n := len(old)
    item := old[n-1]
    *heap = old[:n-1]
    return item
}

type TaskManager struct {
    taskMap  map[int]*Task
    taskHeap TaskHeap
}

func Constructor(tasks [][]int) TaskManager {
    tm := TaskManager{
        taskMap     : make(map[int]*Task),
        taskHeap    : TaskHeap{},
    }

    for _, t := range tasks {
        userId, taskId, priority := t[0], t[1], t[2]
        task := &Task{userID: userId, taskID: taskId, priority: priority}
        tm.taskMap[taskId] = task
        heap.Push(&tm.taskHeap, task)
    }

    return tm
}


func (this *TaskManager) Add(userId int, taskId int, priority int)  {
    task := &Task{userID: userId, taskID: taskId, priority: priority}
    this.taskMap[taskId] = task
    heap.Push(&this.taskHeap, task)
}


func (this *TaskManager) Edit(taskId int, newPriority int)  {
    task := this.taskMap[taskId]
    newTask := &Task{userID: task.userID, taskID: task.taskID, priority: newPriority}
    this.taskMap[taskId] = newTask
    heap.Push(&this.taskHeap, newTask)
}


func (this *TaskManager) Rmv(taskId int)  {
    delete(this.taskMap, taskId)
}


func (this *TaskManager) ExecTop() int {
    for this.taskHeap.Len() > 0 {
        top := this.taskHeap[0]
        heap.Pop(&this.taskHeap)

        if t, ok := this.taskMap[top.taskID]; ok && t == top {
            delete(this.taskMap, top.taskID)
            return top.userID
        }
    }
    return -1
}



/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */