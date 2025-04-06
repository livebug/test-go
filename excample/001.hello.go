package excample

import "fmt"

// hello.go

func Hello() {

	fmt.Println("Hello, World!")
	fmt.Println("This is a simple Go program.")

	// 读取名称，输出hello + 名称
	name := ""
	fmt.Print("请输入你的名字: ")
	fmt.Scanln(&name)
	fmt.Printf("Hello, %s!\n", name)
	fmt.Println("欢迎使用Go语言！")
	fmt.Println("这是一个简单的Go程序。")
	fmt.Println("Go语言是一种编译型、并发型、垃圾回收型的编程语言。")
	fmt.Println("它的设计目标是提高编程的效率和可维护性。")
	fmt.Println("Go语言的语法简洁明了，易于学习和使用。")
	fmt.Println("Go语言的标准库提供了丰富的功能，可以满足大部分开发需求。")
	fmt.Println("Go语言的并发模型非常强大，可以轻松地处理并发任务。")
	fmt.Println("Go语言的社区活跃，有很多优秀的开源项目和库可以使用。")
	fmt.Println("Go语言的生态系统不断发展，有很多新的工具和框架可以使用。")
	fmt.Println("Go语言的未来非常光明，是一种值得学习和使用的编程语言。")
	fmt.Println("希望你能喜欢Go语言！")
	fmt.Println("如果你有任何问题，请随时问我。")
	fmt.Println("祝你编程愉快！")
	fmt.Println("再见！")
}
