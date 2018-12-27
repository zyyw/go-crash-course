package main

import "fmt"

/*
Note in-place:
1. A slice has three components: a pointer, a length, and a capacity.

2. This is an array initialization / literal:
a := [...]int{0, 1, 2, 3}
This is a slice initialization / literal:
s := []int{0, 1, 2, 3}

3.arrays are comparable, but slice are not. So we cannot use == to test
whether two slices contain the same elements.
For two slice of bytes, to compare, use bytes.Equal();
for other types of slices, implement it yourself.

4. 数组是值类型
*/
func main() {
	// runArrays()

	arr := [...]int{1, 2, 5}
	fmt.Println(arr)
	printArray(arr)
	fmt.Println(arr) // 数组仍然是 [1, 2, 5]
}

func runArrays() {
	// Arrays
	var fruitArr [2]string
	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Banana"

	// Declare and assign
	fruitArr2 := [2]string{"Apple", "Banana"}
	fmt.Println(fruitArr2)
	fmt.Println(fruitArr2[1])

	// create an array, whose size can be referred by the right-hand side
	// 让编译器数数组中有几个元素
	nums := [...]int{1, 3, 5}
	fmt.Println(nums)

	// 遍历数组
	// method 1:
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	// method 2: use range key word
	// 2.1
	for i := range nums {
		fmt.Println(nums[i])
	}
	// 2.2
	for i, v := range nums {
		fmt.Println(i, v)
	}
	// 2.3: 通过下划线 _ 省略变量
	for _, v := range nums {
		fmt.Println(v)
	}
}

// example for note 4
// 调用 func printArray(nums [3]int) 会“拷贝”一份数组
func printArray(nums [3]int) {
	nums[0] = 100 // 这里不会改变传进来的原来的数组的一个元素值为100；数组是值类型!!
	for i := range nums {
		fmt.Println(nums[i])
	}
}
