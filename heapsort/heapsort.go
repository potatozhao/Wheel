package main

import "fmt"
//小根堆

//下沉
func down(list []int, dad int){
	tmp := dad
	for{
		son := tmp * 2 + 1
		if son >= len(list){
			break
		}
		if son + 1 < len(list) && list[son] > list[son+1]{
			son += 1
		}
		if list[son] < list[tmp]{
			list[son],list[tmp] = list[tmp],list[son]
			tmp = son
		}else{
			break
		}
	}
}

func up(list []int, son int){
	tmp := son
	for{
		dad := (tmp-1) / 2
		if tmp == son || list[dad] < list[tmp]{
			break
		}
		list[dad],list[son] = list[son], list[dad]
	}
}

func Push(list []int, dad int) {
	down(list, dad)
}

func Pop(list []int) int {
	if len(list) <= 0 {
		return -1
	}
	if len(list) == 1 {
		return list[0]
	}
	list[0], list[len(list)-1] = list[len(list)-1], list[0]
	tmp := list[len(list)-1]
	list = list[:len(list) - 1]
	Push(list, 0)
	return tmp
}

func MakeHeap(list []int){
	for i := len(list) / 2 - 1; i >= 0; i-- {
		Push(list, i)
	}
}


//大根堆正序，小根堆倒叙，此处需要重新手写。一个完整的heap操作
func HeapSort(list []int){
	MakeHeap(list)
	fmt.Println(list)
	tmp := list
	for i := len(tmp) - 1; i>=1; i--{
		tmp[0], tmp[i] = tmp[i], tmp[0]
		tmp = tmp[:len(tmp) - 1]
		Push(tmp, 0)
	}
}


func main() {
	fmt.Printf("hello, world\n")
	b := []int{1,4,2,6,5,8,9}
	HeapSort(b)
	fmt.Println(b)
}