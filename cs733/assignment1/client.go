package main
//package assignment1

import ( 
	"fmt"
	"os"
//	"time"
	"net"
//	"io/ioutil"
	"bufio"
	)


func main(){
	
	ip_address, err := net.ResolveTCPAddr("tcp4","localhost:8080")
	//fmt.Print("\nIP Address : ",ip_address)
	if err!=nil {
	//fmt.Print("\nError : ",err)
	os.Exit(1)
	}
	conn, err  := net.DialTCP("tcp",nil, ip_address)
	if err!=nil {
	//fmt.Print("\nError : ",err)
	os.Exit(1)
	} else {
	//fmt.Print("\nConnection established : ",conn)

        reader:= bufio.NewReader(conn)
 
 
 /*
        
	content := "ratheeshkv"
//	fmt.Fprintf(conn, "write test.txt %v %v \r\n%v\r\n",20,20,content)
//	fmt.Fprintf(conn, "cas %v 2 20\r\n%v\r\nread %v","test.txt","12345678912345678912","test.txt")
//     	fmt.Fprintf(conn, "read %v\r\n","test.txt")
//     	fmt.Fprintf(conn, "delete %v\r\n","test.txt")


	fmt.Fprintf(conn, "write test.txt %v\r\n%v\r\nwrite test1.txt %v %v\r\n%v\r\n",10,content,10,10,content)
//        fmt.Fprintf(conn, "delete %v\r\n","test.txt")	
//	fmt.Fprintf(conn, "read %v\r\n","test.txt")
  //      fmt.Fprintf(conn, "cas %v 1 20\r\n%v\r\n","test.txt","12345678912345678912")

	data, err := reader.ReadString('\n')
	if (err != nil) {
	fmt.Print(err)	
	} else {
	fmt.Print(string(data[:]))
        }
        
    */    
    //    fmt.Fprintf(conn, "read %v\r\n","test.txt")
      
      fmt.Fprintf(conn, "write test.txt %v %v \r\n%v\r\n",21,20,"engineering the cloud")  

data, err := reader.ReadString('\n')
	if (err != nil) {
	fmt.Print(err)	
	} else {
	fmt.Print(string(data[:]))
        }
      


        
         
    

	}
}

