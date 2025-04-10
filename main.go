package main

import (
	"english-ai-go/interfaces"
	"english-ai-go/math"
	"english-ai-go/methods"

	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// Hàm tính tổng
	sum := math.Add(5, 4, 10)
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
	math.Change(num2[:])
	fmt.Println("num2 sau khi thay đổi là:", num2)

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Lỗi: Không tìm thấy file .env")
	}

	// Lấy biến từ môi trường
	port := os.Getenv("PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	fmt.Println("\nCổng:", port)
	fmt.Println("DB User:", dbUser)
	fmt.Println("DB Pass:", dbPass)

	// Methods
	p := methods.Person{Name: "Anna", Height: 155}
	p.ShowDetail()

	// receiver
	// vd1
	p.ChangeHeight(170) // ko thay đổi dữ liệu
	fmt.Println("\nvalue receiver height:", p.Height)

	p.ChangeHeight2(162) // thay đổi dữ liệu
	fmt.Println("pointer receiver height:", p.Height)

	// vd2
	p.SetName("Mike")
	p.SetHeight(180)
	fmt.Println("pointer receiver name:", p.Name)
	fmt.Println("pointer receiver height:", p.Height)

	// int
	var n methods.MyInt = 10
	fmt.Println("\nKết quả:", n.CheckInt())
	// string
	var s methods.MyString = "hello"
	fmt.Println(s.CheckString())

	// interface
	var d interfaces.Animal

	d = interfaces.Dog{}
	fmt.Println(d.Speak())
	d = interfaces.Cat{}
	fmt.Println(d.Speak())

	// constructor
	p1 := methods.NewPerson("Jason", 33)
	p2 := methods.NewPerson("Hana", 25)
	fmt.Println("\nBefore")
	p1.ShowValue()
	p2.ShowPointer()

	p1.ChangeValue("David")
	p2.ChangePointer("Rachel")
	fmt.Println("After")
	p1.ShowValue()
	p2.ShowPointer()
}
