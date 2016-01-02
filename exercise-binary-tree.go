package main

import ("golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	WalkRecursivo(t,ch)
	close(ch)
}

//necessario pq nao eh possivel fazer recursao no Walk, se nao estariamos fechando o canal antes de termimar a recursao
func WalkRecursivo(t *tree.Tree, ch chan int){
	if(t.Left != nil) {
		WalkRecursivo(t.Left, ch)
	}

	ch <- t.Value

	if(t.Right != nil){
		WalkRecursivo(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := range ch1 {
		if (i != <-ch2) {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(Same(tree.New(2),tree.New(1)))
	fmt.Println(Same(tree.New(1),tree.New(1)))
}

