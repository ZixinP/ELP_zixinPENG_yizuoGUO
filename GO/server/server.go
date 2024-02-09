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

	e.activeConnections++

	return &Connection{Conn: conn}, nil
}

func (e *Environment) closeConnection(connection *Connection) {
	e.connectionMutex.Lock()
	defer e.connectionMutex.Unlock()

	connection.Conn.Close()

	e.activeConnections--
}


	

// fonction for image processing
func image_process(image_init image.Image, index int) image.Image {  
	var wg sync.WaitGroup

	ch := make(chan image.Image) // create a channel to store the image data
	var n int  // n is the number of total steps in the image processing

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go func(step int) {   // create a goroutine for each step
			defer wg.Done()
			img := <-ch
			result := processStep(step, index, img)
			ch <- result
		}(i)
	}

	// initialize the channel with the initial image data
	ch <- image_init

	wg.Wait()

	// get the final result from the channel
	finalResult := <-ch

	return finalResult
}

// for now we only have one step for handling the image: edge detection
func processStep(step, input int, img image.Image) image.Image {
	switch step {
	case 1:
		return hdl.handle_image(img, input)   // call the handle_image function from the handle package
	case 2:
		return img
	case 3:
		return img
	// ...	
	default:
		return img
	}
}


func handleConnection(conn net.Conn) { 

	image, index := fct.Decode_image(conn) 

	image_end := image_process(image, index) 

	image_strings,err:=fct.Encode_image(image_end) 

	if err != nil {
		fmt.Println("Error encoding image:", err)
		return
	}

	conn.Write([]byte(image_strings)) // send the image data back to the client

}

func main() {
	//  create a new environment with a maximum of 5 connections
	env := NewEnvironment(5)
	
	listener, err := net.Listen("tcp", ":8080") // listen on port 8080
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

		connection, err := env.createConnection(conn) // create a new connection
		if err != nil {
			fmt.Println("Error creating connection:", err)
			return
		}
        defer env.closeConnection(connection) // close the connection when the function returns

		handleConnection(connection.Conn)

	}
}






// an idea for using a connection pool,but it's not necessary for this project
/*
type Environment struct {     // struct for the environment
	pool    *sync.Pool
	maxSize int
	mu      sync.Mutex
}

type Connection struct {      // struct for the connection
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


func (e *Environment) getConnection() *Connection {     
	e.mu.Lock()         // apply a lock to the environment to ensure that only one goroutine can access the environment at a time
	defer e.mu.Unlock()

	conn := e.pool.Get() 
	return conn.(*Connection)    // convert the interface{} type to *Connection type
}

func (e *Environment) releaseConnection(conn *Connection) {       // release the connection
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.pool == nil {             //if the pool is not initialized, create a new pool
		e.pool = new(sync.Pool)
	}

	e.pool.Put(conn)
}
*/
