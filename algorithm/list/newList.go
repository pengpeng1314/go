package main

import (
	"fmt"
)

//定义节点
type LinkNode struct {
	Data interface{}
	Next *LinkNode
}

//在头结点后面添加元素
func (head *LinkNode) Add(data interface{}) {
	var newNode LinkNode
	newNode.Data = data
	newNode.Next = head.Next
	head.Next = &newNode
}

//在链尾部添加元素
func (head *LinkNode) Append(data interface{}) {
	p := head
	for p.Next != nil {
		p = p.Next
	}
	var newNode LinkNode
	newNode.Data = data
	p.Next = &newNode
	newNode.Next = nil
}

//获取长度,不算头结点
func (head *LinkNode) GetLength() int {
	p := head
	length := 0
	for p.Next != nil {
		length++
		p = p.Next
	}
	return length
}

//删除指定pos位置元素，首元结点为0位置
func (head *LinkNode) Delete(pos int) interface{} {
	length := head.GetLength()
	if pos >= 0 && pos < length {
		pre := head
		var n *LinkNode
		if pos == 0 {
			n = pre.Next
		} else {
			for i := 0; i < pos; i++ {
				pre = pre.Next
			}
			n = pre.Next
		}
		pre.Next = n.Next
		n.Next = nil
		return (*n).Data
	} else {
		fmt.Println("Please give a pos between 0 to length-1!")
		return nil
	}
}

//在指定pos位置前插入元素，首元结点为0位置
func (head *LinkNode) Insert(pos int, data interface{}) {
	var n LinkNode
	n.Data = data
	length := head.GetLength()
	if pos == 0 {
		n.Next = head.Next
		head.Next = &n
	} else if pos > 0 && pos < length {
		pre := head
		for i := 0; i < pos; i++ {
			pre = pre.Next
		}
		n.Next = pre.Next
		pre.Next = &n
	} else {
		fmt.Println("Please give a pos between 0 to length-1!")
	}
}

//获取指定pos位置的元素，首元结点为0位置
func (head *LinkNode) GetData(pos int) interface{} {
	length := head.GetLength()
	if pos >= 0 && pos < length {
		pre := head
		for i := 0; i <= pos; i++ {
			pre = pre.Next
		}
		return (*pre).Data
	} else {
		fmt.Println("Please give a pos between 0 to length-1!")
		return nil
	}
}

//打印所有节点
func (head *LinkNode) Print() {
	p := head
	fmt.Printf("head->")
	for p.Next != nil {
		p = p.Next
		fmt.Printf("%v->", (*p).Data) //写成p.Data也对
	}
	fmt.Printf("nil\n")
}

func main() {
	var head = LinkNode{Data: 0, Next: nil}
	head.Insert(0, 1)
	head.Add(2)
	head.Insert(1, 3)
	head.Append(4)
	head.Print()
	head.Delete(0)
	head.Print()
	fmt.Println(head.GetData(1))

}
