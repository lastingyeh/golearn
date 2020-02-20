package main

import "fmt"

type Point struct {
	x int
	y int
}

type Rect1 struct {
	left, right Point
}

type Rect2 struct {
	left, right *Point
}

func main(){
	r1 := Rect1{Point{1,2}, Point{3,4}}
	// continuously allocated memory 
	fmt.Printf("r1.left.x %p, r1.left.y %p\nr1.right.x %p, r1.right.y %p\n",
		&r1.left.x, &r1.left.y, &r1.right.x, &r1.right.y)

		// self-mem (continuously)
		r2 := Rect2{&Point{10,20}, &Point{30, 40}}
		fmt.Printf("r2.left %p, r2.right %p\n",&r2.left,&r2.right)

		// value-mem (discontinuously)
		fmt.Printf("r2.left %p, r2.right %p\n",r2.left, r2.right)

}