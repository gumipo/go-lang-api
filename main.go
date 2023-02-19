package main

// format package
import (
	"fmt"
	"time"
)

func outer(w int, h int) {
	fmt.Println(w * h)
}

func main() {
	fmt.Println("test")
	fmt.Println(time.Now())

	var neko time.Time = time.Now()

	var t, f bool = true, false

	fmt.Println(t, f)
	fmt.Println(neko)

	var (
		i2 int    = 100
		s2 string = "meiasjf"
	)
	fmt.Println(i2, s2)

	var i3 int // int 初期値０
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "hello"

	fmt.Println(i3, s3)

	// 暗黙的定義
	i4 := 500
	fmt.Println(i4)

	outer(3, 6)
	// i4 := 600 // 再定義できない
}
