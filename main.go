package main

import "fmt"

func add(a, b, c int) int {
	return a + b + c
}

func change(sls []int) {
	sls[2] = 50
	sls[3] = 0
}

func main() {
	// Hàm tính tổng
	sum := add(5, 4, 10)
	fmt.Println("5 + 4 + 10 =", sum)

	// Mảng, vòng lặp
	arr := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// Slices
	// Khai báo 1 slices num, thêm 2 số bất kỳ vào num, in ra len(num) và cap(num)
	// sao chép từ num vào num2, cắt để lấy các pt bắt đầu từ pt thứ 4 trong num2
	// duyệt for range và in ra tất cả các pt trong num2
	num := []int{10, 11, 12, 13, 14}
	num = append(num, 20, 22)
	fmt.Println("\nlen(num) =", len(num), ", cap(num) =", cap(num))

	num2 := make([]int, len(num))
	copy(num2, num)
	num2 = num2[3:]

	for _, value := range num2 {
		fmt.Println(value)
	}

	// Pointer
	// Cho số nguyên a = 100, in ra địa chỉ của a thông qua con trỏ b
	// cho con trỏ c trỏ vào b, in ra giá trị c
	a := 100
	var b *int = &a
	fmt.Println("\nĐịa chỉ của a là:", b)

	c := b
	fmt.Println("Giá trị của c là:", *c)

	// thay đổi 1 vài phần tử trong num2 thông qua con trỏ trong hàm
	change(num2[:])
	fmt.Println("num2 sau khi thay đổi là:", num2)
}
