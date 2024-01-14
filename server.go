package main

import (
	"fmt"
	"net"
	"sync"
)

type Connection struct {
	ID int
	// 添加其他必要的连接信息
}

type Environment struct { // 环境结构体，为了限制连接池的大小
	pool    *sync.Pool
	maxSize int
	mu      sync.Mutex
}

func NewEnvironment(maxSize int) *Environment { // 创建环境结构体
	return &Environment{
		pool:    &sync.Pool{},
		maxSize: maxSize,
	}
}

func (e *Environment) getConnection() *Connection { // 获取连接
	e.mu.Lock()        //执行函数时加锁，。为了保证在多个goroutine中不会出现竞争条件
	defer e.mu.Unlock()

	conn := e.pool.Get() //从连接池中获取连接
	if conn == nil {
		return &Connection{ID: 1}
	}
	return conn.(*Connection) //将interface{}类型转换为*Connection类型
}

func (e *Environment) releaseConnection(conn *Connection) { // 表示这个函数是属于Environment类型的方法，用于释放连接到连接池
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.pool == nil { //如果还没有初始化连接池的本地池。在这种情况下，使用 new(sync.Pool) 初始化一个新的 sync.Pool 对象。
		e.pool = new(sync.Pool)
	}

	e.pool.Put(conn)
}

func handleConnection(env *Environment, conn net.Conn) {
	defer conn.Close()

	connection := env.getConnection()
	defer env.releaseConnection(connection)

	fmt.Printf("Handling connection %d\n", connection.ID)
	// 加入处理连接的代码
}

func main() {

	listener, err := net.Listen("tcp", ":8080") //监听8080端口
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	//  创建一个环境结构体，最多可以容纳10个连接
	env := NewEnvironment(10)

	fmt.Println("Server listening on :8080")

	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// 为每个连接启动一个goroutine
		go handleConnection(env, conn)
	}
}
