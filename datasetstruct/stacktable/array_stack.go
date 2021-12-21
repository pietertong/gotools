package stacktable

import (
	"errors"
	"sync"
)

type SizeType int

type Element interface{}

type ArrayStack struct {
	array []Element  `json:"array"`
	size  SizeType   `json:"size"`
	lock  sync.Mutex `json:"lock"`
}

// New 初始化
func New() *ArrayStack {
	return &ArrayStack{
		array: make([]Element, 0),
		size:  0,
		lock:  sync.Mutex{},
	}
}

// Push 入栈
func (a *ArrayStack) Push(elements []Element) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.array = append(a.array, elements...)
	a.size = a.size + SizeType(len(elements))
	return nil
}

// Pop 出栈
func (a *ArrayStack) Pop() (Element, error) {
	length := a.Len()
	if length < SizeType(1) {
		return Element(""), errors.New("stack less")
	}
	a.lock.Lock()
	defer a.lock.Unlock()
	v := a.array[length-1]
	cpArray := make([]Element, length-1, length-1)
	for i := SizeType(0); i < SizeType(length-SizeType(1)); i++ {
		cpArray[i] = a.array[i]
	}
	a.array = cpArray
	a.size = a.size - 1
	return Element(v), nil
}

// Len 长度
func (a *ArrayStack) Len() SizeType {
	return a.size
}

// Search 搜索
func (a *ArrayStack) Search(find Element) (res bool, err error) {
	res = false
	length := a.Len()
	if length < SizeType(1) {
		return res, errors.New("stack less")
	}
	a.lock.Lock()
	defer a.lock.Unlock()
	for i := SizeType(0); i < length-SizeType(1); i++ {
		if a.array[i] == find {
			res = true
			break
		}
	}
	return res, nil
}

// IsEmpty 判断栈是否为空
func (a *ArrayStack) IsEmpty() bool {
	length := a.Len()
	if length < 1 {
		return true
	}
	return false
}
