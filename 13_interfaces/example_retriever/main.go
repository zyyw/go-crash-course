package main

import (
	"fmt"
	"time"

	"github.com/zyyw/go-crash-course/13_interfaces/example_retriever/mock"
	"github.com/zyyw/go-crash-course/13_interfaces/example_retriever/real"
)

// interface 不是一个简单的引用，肚子里有两个东西：
// * 类型
// * 值
// 可以通过: switch 或者 type assert 来判断 (查看接口变量的方法)
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"This is a fake imooc.com"}
	r = &retriever
	inspect(r)

	r = &real.Retriever{"Chrome/1.9.3", 41 * time.Second}
	inspect(r)

	// type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)
	if mockRetriver, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriver.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
	// fmt.Println(download(r))
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)

	// type switch
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	default:
		fmt.Printf("Unkonw type: %v\n", v)
	}
}
