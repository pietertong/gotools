package chaintable

import "errors"

/**
链表的概念
单链表是一种链式存取的数据结构，用一组地址任意的存储单元存放线性表中的数据元素。链表中的数据是以结点来表示的，
每个结点的构成：元素(数据元素的映象) + 指针(指示后继元素存储位置)，元素就是存储数据的存储单元，指针就是连接每个结点的地址数据。


链表 -> 初始化
链表 -> 遍历输出
链表 -> 修改
链表 -> 头部插入，尾部插入
链表 -> 删除
*/

type Element interface{}

type LinkNode struct {
	Data Element
	Next *LinkNode
}

type HeadLinkNode struct {
	Length int
	Node   *LinkNode
}

type LinkNoder interface {
	Add(node Element)
	Foreach() []Element
	Remove(index int) error
	Insert(index int, node Element)
	Len() int
	Search(node Element) int
	Get(index int) *LinkNode
}

func New() *HeadLinkNode {
	return &HeadLinkNode{
		Length: 0,
		Node: &LinkNode{
			Data: "",
			Next: nil,
		},
	}
}

func (h *HeadLinkNode) Foreach() (res []Element) {
	l := h.Node
	res = make([]Element, 0, h.Length)
	for {
		if l == nil || (l.Next == nil && l.Data == "") {
			break
		}
		if len(l.Data.(string)) > 0 {
			res = append(res, l.Data)
		}
		l = l.Next
	}
	return res
}

func (h *HeadLinkNode) Add(node Element) {
	l := h.Node
	for {
		if l != nil && l.Next == nil {
			newNode := &LinkNode{
				Data: node,
				Next: nil,
			}
			l.Next = newNode
			break
		} else {
			l = l.Next
		}
	}
	h.Length++
}

func (h *HeadLinkNode) Remove(index int) error {
	length := h.Len()
	if index < 1 || index > length {
		return errors.New("index error")
	}
	l := h.Node

	for i := 0; i < index-1; i++ {
		l = l.Next
	}
	if index == length {
		l.Next = nil
	} else {
		l.Next = l.Next.Next
	}
	h.Length--
	return nil
}
func (h *HeadLinkNode) Insert(index int, node Element) {
	length := h.Len()
	if index < 1 || index > length {
		return
	}
	l := h.Node

	for i := 0; i < index-1; i++ {
		l = l.Next
	}
	newNode := &LinkNode{
		Data: node,
		Next: l.Next,
	}
	l.Next = newNode
	h.Length++
}
func (h *HeadLinkNode) Len() int {
	return h.Length
}
func (h *HeadLinkNode) Search(node Element) int {
	l := h.Node
	index := -1
	for {
		if l == nil {
			index = -1
			break
		}
		if l != nil && l.Next == nil {
			if l.Data == node {
				break
			}
			index = -1
			break
		}
		index++
		if l.Data == node {
			break
		}
		l = l.Next
	}
	return index
}

func (h *HeadLinkNode) Get(index int) *LinkNode {
	length := h.Len()
	if index < 1 || length < index {
		return nil
	}

	l := h.Node
	for i := 0; i < index; i++ {
		l = l.Next
	}
	return l
}
