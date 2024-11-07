package main

import(
    "fmt"

)

type Node struct{
    value int
    left *Node
    right *Node
}

func createTree(arr *[15]int, index int) *Node {

    if index >= len(arr) {
        return nil
    }

    root := &Node{
        value: arr[index],
        left: nil,
        right: nil,
    }

    root.left = createTree(arr, 2*index + 1)
    root.right = createTree(arr, 2*index + 2)
    

    return root
}


func BFS(root *Node){
    if root == nil {
        return
    }

    queue := []*Node{root}
   
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        fmt.Printf("%d", node.value)
        fmt.Println()

        if node.left != nil{
            queue = append(queue, node.left)
        }
        if node.right != nil{
            queue = append(queue, node.right)
        }
    }
}

func DFS(root *Node){
    if root == nil {
        return
    }

    stack := []*Node{root}
   
    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        fmt.Printf("%d", node.value)
        fmt.Println()

        if node.left != nil{
            stack = append(stack, node.left)
        }
        if node.right != nil{
            stack = append(stack, node.right)
        }
    }
}

func main(){
    array := [15]int{1, 3, 5, 6, 7, 24, 213, 3, 2, 23, 34, 12, 123, 33, 44}
    root := createTree(&array, 0)
    //BFS(root)
    DFS(root)
    
}