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
	
	ip_address, err := net.ResolveTCPAddr("tcp",os.Args[1])
	fmt.Print(ip_address)
	fmt.Print(err)

	listen, err  := net.ListenTCP("tcp", ip_address)
	fmt.Print(err)
//	fmt.Print(conn)

	for {
	conn, err := listen.Accept()
	fmt.Printf("One conn Accepted")

	if err==nil{
 	result, err := ioutil.ReadAll(conn)
	fmt.Print(err)
	fmt.Printf("%s",result)

	check, err := conn.Write([]byte("hello goodevening"))
	fmt.Print(err)
	fmt.Print(check)

	conn.Close()
	}

	}


	 


	

	}


}

