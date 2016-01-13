package main
import ( 
	"fmt"
	"os"
	"net"
	"io/ioutil"
	)

func main(){

	if(len(os.Args)<2 || len(os.Args)>2 ) {
		
		fmt.Printf("Input Format : go run client.go ip_address:port_no\r\n")				
		os.Exit(1)
	}else {
	
	ip_address, err := net.ResolveTCPAddr("tcp4",os.Args[1])
	fmt.Print(ip_address)
	fmt.Print(err)
	conn, err  := net.DialTCP("tcp",nil, ip_address)
	fmt.Print(err)
	fmt.Print(conn)
	check, err := conn.Write([]byte("hello ratheesh\n"))
	fmt.Print(err)
	fmt.Print(check)
	result, err := ioutil.ReadAll(conn)
	fmt.Print(err)
	fmt.Print(result)

	os.Exit(0)


	

	}


}

