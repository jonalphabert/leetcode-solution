/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    dummy := &ListNode{Next: head}
    prev := dummy
    currentNode := head

    for currentNode != nil && currentNode.Next != nil {
        first := currentNode
        second := currentNode.Next

        prev.Next = second
        first.Next = second.Next
        second.Next = first

        prev = first
        currentNode = prev.Next
    }

    return dummy.Next
}