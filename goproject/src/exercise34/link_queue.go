package main
import (
	"fmt"
)

type Node struct {
	name string
	next *Node
}

func(this *Node) AddNext(next *Node) {
	this.next = next
}


func Put(head *Node, name string) {
	newNode := Node{
		name: name,
	}
	//寻找队尾节点
	tailNode := head
	for {
		if tailNode.next == nil {
			break
		}
		tailNode = tailNode.next
	}
	tailNode.AddNext(&newNode)
	fmt.Printf("%v入队\n", name)
}

func Get(head *Node) {
	if head.next == nil {
		fmt.Println("队列为空")
		return
	}
	node := head.next
	head.next = node.next
	fmt.Printf("%v出队\n", node.name)
}

func Show(head *Node) {
	if head.next == nil {
		fmt.Println("队列为空")
		return
	}

	curNode := head.next
	for {
		fmt.Printf("%v<%p> -> ", curNode, curNode)
		curNode = curNode.next
		if curNode == nil {
			break
		}
	}
}


func main() {
	//初始化一个头节点
	head := Node{}
	Show(&head)
	Put(&head, "孙悟空")
	Put(&head, "猪八戒")
	Put(&head, "沙悟净")
	Show(&head)
	fmt.Println()
	Get(&head)
	Get(&head)
	Get(&head)
	Get(&head)
}