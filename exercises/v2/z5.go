package v2

type ArrayStack struct{
	Elements []int
}

type LinkedStack struct{
	Tops *Node
}

type Node struct{
	Value int
	Next *Node
}

func (s *ArrayStack) Pop() int{
	if len(s.Elements) == 0{
		panic("Stack is empty")
	}
	element := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return element
}

func (s *LinkedStack) Pop() int{
	if s.Tops == nil{
		panic("Stack is empty")
	}
	value := s.Tops.Value
	s.Tops = s.Tops.Next
	return value
}

func (s *ArrayStack) Push(value int){
	s.Elements = append(s.Elements, value)
}

func (s *LinkedStack) Push(value int){
	s.Tops = &Node{Value: value, Next: s.Tops}
}

func (s *ArrayStack) Top() int{
	if len(s.Elements) == 0{
		panic("Stack is empty")
	}
	element := s.Elements[len(s.Elements)-1]
	return element
}

func (s *LinkedStack) Top() int{
	if s.Tops == nil{
		panic("Stack is empty")
	}
	return s.Tops.Value
}

func (s *ArrayStack) IsEmpty() bool{
	return len(s.Elements) == 0
}

func (s *LinkedStack) IsEmpty() bool{
	return s.Top == nil
}