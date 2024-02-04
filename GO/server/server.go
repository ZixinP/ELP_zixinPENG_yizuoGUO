package server


import (
	"fmt"
	"image"
	"net"
	"sync"
	fct "../fonctions"
	hdl "../handle"
)

type Request struct {
	Image        string `json:"image"`
	IntParameter int    `json:"intParameter"`
}

type Connection struct {
	Conn net.Conn
}

type Environment struct {
	activeConnections int
	maxConnections    int
	connectionMutex   sync.Mutex
}

func NewEnvironment(MaxConnections int) *Environment {
	return &Environment{
		maxConnections: MaxConnections,
	}
}

func (e *Environment) createConnection(conn net.Conn) (*Connection, error) {
	e.connectionMutex.Lock()
	defer e.connectionMutex.Unlock()

	if e.activeConnections >= e.maxConnections {
		return nil, fmt.Errorf("exceeded maximum allowed connections")
	}

	// 更新活跃连接数
	e.activeConnections++

	return &Connection{Conn: conn}, nil
}

func (e *Environment) closeConnection(connection *Connection) {
	e.connectionMutex.Lock()
	defer e.connectionMutex.Unlock()

	// 关闭连接
	connection.Conn.Close()

	// 更新活跃连接数
	e.activeConnections--
}


	


func image_process(image_init image.Image, index int) image.Image { // 图像处理函数
	var wg sync.WaitGroup

	ch := make(chan image.Image) // 创建一个无缓冲的通道
	var n int // n 是图片处理需要的总步数，后面改

	// 循环创建goroutine
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go func(step int) {
			defer wg.Done()
			img := <-ch
			result := processStep(step, index, img)
			ch <- result
		}(i)
	}

	// 初始化第一个输入并开始循环

	ch <- image_init

	// 等待所有 goroutine 执行完毕
	wg.Wait()

	// 最终结果
	finalResult := <-ch

	return finalResult
}

func processStep(step, input int, img image.Image) image.Image {
	switch step {
	case 1:
		return hdl.handle_image(img, input)
	case 2:
		return nil
	case 3:
		return nil
	// ...	
	default:
		return nil
	}
}


func handleConnection(conn net.Conn) { // 处理连接

	image, index := fct.Decode_image(conn) // 解码客户端发送的图像数据

	image_end := image_process(image, index) // 图像处理函数

	image_strings,err:=fct.Encode_image(image_end) // 将图像数据编码为 base64 字符串

	if err != nil {
		fmt.Println("Error encoding image:", err)
		return
	}

	conn.Write([]byte(image_strings)) // 将图像数据发送给客户端

}

func main() {
	//  创建一个环境结构体，最多可以容纳10个连接
	env := NewEnvironment(5)
	
	listener, err := net.Listen("tcp", ":8080") //监听8080端口
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	

	fmt.Println("Server listening on :8080")

	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		connection, err := env.createConnection(conn) // 创建连接
		if err != nil {
			fmt.Println("Error creating connection:", err)
			return
		}
        defer env.closeConnection(connection) // 关闭连接

		// 在处理连接的过程中，将 Connection 对象传递给其他函数
		handleConnection(connection.Conn)

	}
}







/*
type Environment struct { // 环境结构体，为了限制连接池的大小
	pool    *sync.Pool
	maxSize int
	mu      sync.Mutex
}

type Connection struct { // 连接结构体
	conn net.Conn
}

func NewEnvironment(Maxsize int) *Environment {   
    return &Environment{
        pool: &sync.Pool{
            New: func() interface{} {
                return &Connection{conn: nil}
            },
        },
        maxSize: Maxsize,
    }
}


func (e *Environment) getConnection() *Connection { // 从连接池中获取连接
	e.mu.Lock() //执行函数时加锁，为了保证在多个goroutine中不会出现竞争条件
	defer e.mu.Unlock()

	conn := e.pool.Get() //从连接池中获取连接
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
*/