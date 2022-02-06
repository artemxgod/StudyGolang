package main

import (
	"fmt"
	"net"
	"bufio"
)

func main(){
	//looking for connect on port 5000
	listener, _ := net.Listen("tcp", ":5000")
	fmt.Println("we are running on http://127.0.0.1:5000")
	for {
		//waiting to accept a connect
		conn, err := listener.Accept()
		if err != nil{
			fmt.Println("Connection went wrong ;(")
			conn.Close()
			continue
		}

		fmt.Println("connected")

		//creating a reader to read from connect
		bufReader := bufio.NewReader(conn)
		fmt.Println("statr reading")

		//creating a goroutine to get all the connections
		go func(conn net.Conn){
			defer conn.Close()

			for {
				//reading by byte
				rbyte, err := bufReader.ReadByte()
				if err != nil{
					fmt.Println("Cant read", err)
					break 
				}
	
				fmt.Print(string(rbyte))
				conn.Write([]byte("received")) // write on server
			}
			

		}(conn)
		
	}
}
