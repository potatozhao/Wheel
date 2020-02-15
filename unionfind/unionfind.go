package main

import (
	"fmt"
)

// UnionFind union find struct
type UnionFind struct {
	Prev []int
	Rank []int
	Count int
}

// NewUnionFind ...
func NewUnionFind(n int) UnionFind{
	uf := UnionFind{
		Prev: make([]int, n),
		Rank: make([]int, n),
		Count: n,
	}
	for i := range uf.Prev{
		uf.Prev[i] = i
	}
	return uf
}

// Find ...
func (this *UnionFind)Find(s int) int{
	if s >= len(this.Prev){
		return -1
	}
	if s == this.Prev[s]{
		return s
	}
	this.Prev[s] = this.Find(this.Prev[s])
	return this.Prev[s]
}

// Unite ...
func (this *UnionFind)Unite(a,b int){
	if a >= len(this.Prev) || b >= len(this.Prev) || a <0 || b < 0{
		return 
	}
	pa := this.Find(a)
	pb := this.Find(b)
	if pa == pb{
		return
	}
	if this.Rank[pa] < this.Rank[pb]{
		this.Prev[pa] = this.Prev[pb]
		this.Rank[pb] += this.Rank[pa] + 1
	}else{
		this.Prev[pb] = this.Prev[pa]
		this.Rank[pa] += this.Rank[pb] + 1
	}
	this.Count--
	return
}

// Same ... 
func (this *UnionFind)Same(a,b int) bool{
	return this.Find(a) == this.Find(b)
}

func (this *UnionFind)GetCount() int{
	return this.Count
}

func main(){
	fmt.Println("helloworld")
	uf := NewUnionFind(10)
	uf.Unite(1, 9)
	uf.Unite(9, 5)
	fmt.Println(uf.Find(5))
	fmt.Println(uf.Rank)
	return
}