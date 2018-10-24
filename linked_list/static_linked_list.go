package gostruct

import (
	"errors"
	"fmt"
)

//StaticListNode 静态链表的node
type StaticListNode struct {
	data interface{}
	next uint
}

//StaticLinkedList 静态链表
type StaticLinkedList struct {
	data   []StaticListNode
	length uint
}

//NewStaticListNode 静态链表
func NewStaticListNode(value interface{}, index uint) StaticListNode {
	return StaticListNode{data: value, next: index}
}

//GetNextIndex 设置下一个index
func (s *StaticListNode) GetNextIndex() uint {
	return s.next
}

//GetData
func (s *StaticListNode) GetData() interface{} {
	return s.data
}

//SetData
func (s *StaticListNode) SetData(value interface{}) {
	s.data = value
}

//SetNextIndex
func (s *StaticListNode) SetNextIndex(index uint) {
	s.next = index
}

//NewStaticLinkedList
func NewStaticLinkedList(cap uint) (*StaticLinkedList, error) {
	if cap <= 2 {
		return nil, errors.New("cap need > 2")
	}
	data := make([]StaticListNode, cap, cap)

	var length uint
	for ; length < cap; length++ {
		//0 1
		//1 1
		//2 1
		data[length].SetNextIndex(length + 1)
	}
	data[cap-2].SetNextIndex(0)
	data[cap-1].SetNextIndex(0)
	return &StaticLinkedList{data: data, length: 0}, nil
}

//Len
func (s *StaticLinkedList) Len() uint {
	return s.length
}

func (s *StaticLinkedList) olen() uint {
	return uint(len(s.data))
}

//IsIndexOutOffRange 索引是否越界, true 越界， false 不越界
func (s *StaticLinkedList) IsIndexOutOffRange(index uint) bool {
	//0 [1] [2] [3] [4] 5
	if index >= 1 && index < s.olen()-1 {
		return false
	}
	return true
}

//FreeNodeIndex 获取下一个空链的索引
func (s *StaticLinkedList) FreeNodeNextIndex() uint {
	return s.data[0].GetNextIndex()
}

//IsFull
func (s *StaticLinkedList) IsFull() bool {
	if s.FreeNodeNextIndex() == s.HeadNodeNextIndex() {
		return true
	}
	return false
}

//IsEmpty
func (s *StaticLinkedList) IsEmpty() bool {
	if 0 == s.HeadNodeNextIndex() {
		return true
	}
	return false
}

//HeadNodeIndex 获取当前链表头节点指向的索引
func (s *StaticLinkedList) HeadNodeNextIndex() uint {
	return s.data[s.olen()-1].GetNextIndex()
}

func (s *StaticLinkedList) PreNodeIndex(index uint) (uint, error) {
	var err error
	stop := false
	headNextIndex := s.HeadNodeNextIndex()
	last := s.HeadNodeNextIndex()
	for !stop {
		node, err := s.FindByIndex(headNextIndex)
		if err != nil {
		} else {
			if node.GetNextIndex() == 0 {
				stop = true
				last = headNextIndex
			} else {
				headNextIndex = node.GetNextIndex()
			}
		}
	}
	return last, err
}

func (s *StaticLinkedList) FindByIndex(index uint) (StaticListNode, error) {
	if s.IsIndexOutOffRange(index) {
		return StaticListNode{}, errors.New("索引越界")
	}
	return s.data[index], nil
}

func (s *StaticLinkedList) InsertAfter(value interface{}, index uint) bool {
	if s.IsIndexOutOffRange(index) {
		return false
	}

	//设置此index对应的value
	s.data[index].SetData(value)

	//现在的freeNodeIndex是这个
	nextIndex := s.data[index].GetNextIndex()

	//如果需要插入的索引与空链表的next索引一致，需要将此index的nextindex 设置为空链的index
	if index == s.FreeNodeNextIndex() {
		s.data[index].SetNextIndex(0)
	}
	s.data[0].SetNextIndex(nextIndex)
	//如果这个链表的长度为0，则需要设置头结点
	if 0 == s.length {
		s.data[s.olen()-1].SetNextIndex(index)
	}

	if index != s.HeadNodeNextIndex() {
		proIndex, err := s.PreNodeIndex(index)
		if err == nil {
			s.data[proIndex].SetNextIndex(index)
		}
	}
	s.length++
	fmt.Println(s.data)
	return true
}

func (s *StaticLinkedList) Insert(value interface{}) bool {
	return s.InsertAfter(value, s.FreeNodeNextIndex())
}

func (s *StaticLinkedList) DeleteNode(index uint) bool {
	if s.IsEmpty() {
		return false
	}
	if index > 2 {
		return false
	}
	//头节点的index
	headNodeIndex := s.HeadNodeNextIndex()
	//index对应的node
	indexNode, err := s.FindByIndex(index)
	if err != nil {
		return false
	}
	preNode, err := s.FindByIndex(index - 1)
	if err == nil {
		s.data[preNode.GetNextIndex()].SetNextIndex(indexNode.GetNextIndex())
	}
	if headNodeIndex == indexNode.GetNextIndex() {
		s.data[headNodeIndex].SetNextIndex(indexNode.GetNextIndex())
	}
	s.data[index] = StaticListNode{next: indexNode.GetNextIndex()}
	s.length--
	return true
}

//Traverse 循环
func (s *StaticLinkedList) Traverse() {
	if s.Len() > 0 {
		index := s.HeadNodeNextIndex()
		for index != 0 {
			node, err := s.FindByIndex(index)
			if err != nil {
				fmt.Println(err)
				break
			} else {
				fmt.Println(node.GetData())
			}
			index = node.GetNextIndex()
		}
	}
	fmt.Println(fmt.Sprintf(" static linked list's length is %d ", s.Len()))
}
