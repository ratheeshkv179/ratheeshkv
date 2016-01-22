package main

//package assignment1

import ( 
	"fmt"
	"sync"
	"net"
	"strings"
	"bufio"
	"strconv"
	"io"
//	"io/ioutil"
	"time"
	"github.com/syndtr/goleveldb/leveldb"   
	)
 
var mutex = &sync.Mutex{}
var file_version uint64 = 0
var version = make(map [string] uint64)
var expiry = make(map [string] int64)
var exp_sec = make(map [string] int64)

var db, err = leveldb.OpenFile("/tmp/files.db", nil)
	
func serverMain(){

	ip_address, err := net.ResolveTCPAddr("tcp4","localhost:8080")
	if err!=nil {
	//fmt.Print("\nError : ",err)
//	os.Exit(1)
	} else {
	//fmt.Print("\nServer started, Listening on port :8080")
	go checkexpiry()
	
	listen, _  := net.ListenTCP("tcp", ip_address)

	
	for {
	conn, err := listen.Accept()
	if( err!=nil){
	//fmt.Print("\nError : ",err)
//	os.Exit(1)
	} else {
	go client(conn)

	}}}
}


func checkexpiry(){

        for true {
        time.Sleep(1000000000)
        for k := range expiry {
        cur_time  := time.Now().Unix()
        if(exp_sec[k] != -1 && exp_sec[k]!=0) {
        if (expiry[k]<=cur_time) {
        delete(expiry,k)
        delete(exp_sec,k)
        delete(version,k)
        err = db.Delete([]byte(k), nil)
        }}
    }}
}


func client(conn net.Conn){

	//fmt.Println("\nClient connected, connection ID :",conn)
 
       	reader := bufio.NewReader(conn)
        		
	for true {

	data, err := reader.ReadBytes('\n')

	if(err== io.EOF) {
	//conn.Close()
	//fmt.Print("conn closed")
	break
	} else	if (err!=nil){
	//fmt.Print("\nError : ",err)
	fmt.Fprintf(conn,"ERR_INTERNAL\r\n")
	
	} else {
	
	//fmt.Println("\nNo of Bytes read : ",data)	
	//fmt.Println("\nNo of Bytes read1 : ",err)	

	if(data[len(data)-2]!='\r') {
	//fmt.Print("\nError : ",err)
	fmt.Fprintf(conn,"ERR_INTERNAL\r\n")
	} else {
	command := string(data[0:len(data)-2])
	part := strings.Fields(command)



	if((strings.Compare(part[0],"write"))==0) {

	if(len(part)==3 ||len(part)==4){
	numbytes,err := strconv.Atoi(part[2])
	//fmt.Print(numbytes)		
	if(err != nil ) { // number of bytes in not a number
	//fmt.Print("\nError : ",err)
	fmt.Fprintf(conn,"ERR_INTERNAL\r\n")
	} else {

        flag := 0
        expiry_time := 0

	if(len(part)==4){            // Expiry time present
	expiry_time,err = strconv.Atoi(part[3])
	//fmt.Print(expiry_time)		
	if(err != nil ) { // number of bytes in not number
	flag = 1
	//fmt.Print("\nError : ",err)
	fmt.Fprintf(conn,"ERR_INTERNAL\r\n")
	}}
	
       if ( flag == 0){
	 
	data, err := reader.ReadBytes('\n')
	var content [] byte
	for (data[len(data)-2] != '\r') {
	content = append(content,data...)
	data, err = reader.ReadBytes('\n')
	}
	content = append(content,data[:len(data)-2]...)
	//content = bytes.Trim(content," ")
	//fmt.Print("\n",len(content),content,numbytes,"#")
	if(len(content)!=numbytes) {
	//fmt.Print("\nError : nmbytes mismatch")
	fmt.Fprintf(conn,"ERR_INTERNAL\r\n")
	} else {	
//	_, err = file.Write(content)
//        err = db.Put([]byte("key"), []byte("novalue"), nil)
        err = db.Put([]byte(string(part[1])), []byte(content), nil)
        	
	if (err!=nil) {
	//fmt.Print("\nError : ",err)
	fmt.Fprintf(conn,"ERR_INTERNAL\r\n")
	} else {
	
        mutex.Lock()
	file_version  ++
	version[part[1]] =  file_version // Update file entry
	fmt.Fprintf(conn,"OK %v\r\n",version[part[1]])	
	if (len(part)==4) {
	exp_sec[part[1]] = int64(expiry_time)
        expiry[part[1]] = int64(expiry_time) + time.Now().Unix()
        } else {
        exp_sec[part[1]] = -1
        }
        mutex.Unlock()
        }
	content = nil
        }}}
        } else {
        fmt.Fprintf(conn,"ERR_CMD_ERR\r\n")
        }
        
        
        
        
        }  else if ((strings.Compare(part[0],"cas"))==0) {

	if((len(part)==4 ||len(part)==5) && (len(part[1])<=256)){
      	_, err := db.Get([]byte(string(part[1])), nil)
	if(err != nil){
	//fmt.Print("\nError : ",err)
	fmt.Fprintf(conn,"ERR_FILE_NOT_FOUND\r\n")	
	} else {
	versn,err := strconv.Atoi(part[2])
	if(err != nil ) { // version in not a number
	//fmt.Print("\nError : ",err)
	fmt.Fprintf(conn,"ERR_INTERNAL\r\n")
	} else {
	numbytes,err := strconv.Atoi(part[3])
	if(err != nil ) { // number of bytes in not a number
	//fmt.Print("\nError : ",err)
	fmt.Fprintf(conn,"ERR_INTERNAL\r\n")
	} else {
	if(version[part[1]] != uint64(versn)) {
	//fmt.Print("\nError : version mismatch ")
	fmt.Fprintf(conn,"ERR_VERSION\r\n")
	} else {
	flag := 0
	exp := 0
	if(len(part)==5) {
	exp,err = strconv.Atoi(part[4])
	if(err != nil ) { // number of bytes in not a number
	//fmt.Print("\nError : ",err)
	fmt.Fprintf(conn,"ERR_INTERNAL\r\n")
	flag = 1
	}}
	if (flag == 0){
	data, _ := reader.ReadBytes('\n')
	var content [] byte
	for (data[len(data)-2] != '\r') {
	content = append(content,data...)
	data, err = reader.ReadBytes('\n')
	}
	content = append(content,data[:len(data)-2]...)
	//content = bytes.Trim(content," ")
        //fmt.Print("\n",len(content),numbytes,"#")
        
	if(len(content)!=numbytes) {
	//fmt.Print("\nError : size mismatch")
	fmt.Fprintf(conn,"ERR_INTERNAL\r\n")
	} else {
	//delete
        err = db.Put([]byte(string(part[1])), []byte(content), nil)
	if (err!=nil) {
	//fmt.Print("\nError : ",err)
	fmt.Fprintf(conn,"ERR_INTERNAL\r\n")
	} else {
	
        mutex.Lock()
	if (len(part)==4) {
        exp_sec[part[1]]	= -1
	} else {
	exp_sec[part[1]] = int64(exp)
        expiry[part[1]]	=  int64(exp) + time.Now().Unix()
	}
	file_version  ++
	version[part[1]] =  file_version // Update file entry
        fmt.Fprintf(conn,"OK %v\r\n",file_version)	
        mutex.Unlock()
        
	}
	content = nil
	//file.Close()
	}}}}}}
	} else {
	fmt.Fprintf(conn,"ERR_CMD_ERR\r\n")
	}
 
	} else if ((strings.Compare(part[0],"read"))==0) {
	if(len(part)==2 && len(part[1])<=256){

        mutex.Lock()
        contents, err := db.Get([]byte(string(part[1])), nil)
        mutex.Unlock()
        
	//fmt.Print("\nFile Name : ",len(string(part[1])))
	if(err !=nil) {
	
	fmt.Fprintf(conn,"ERR_FILE_NOT_FOUND\r\n")
	}else {
	//fmt.Print("\nFile Content :",string(contents))
	
	if (exp_sec[string(part[1])] == -1){
	fmt.Fprintf(conn,"CONTENTS %v %v\r\n%v\r\n",version[part[1]],len(contents),string(contents))	
	}else {
	fmt.Fprintf(conn,"CONTENTS %v %v %v\r\n%v\r\n",version[part[1]],len(contents),exp_sec[string(part[1])],string(contents))	
	}}
	} else {
	fmt.Fprintf(conn,"ERR_CMD_ERR\r\n")
	}

	} else if ((strings.Compare(part[0],"delete"))==0) {

	//fmt.Print("\nDELETE")
	if(len(part)==2 && len(part[1])<=256){
	
        mutex.Lock()
	
        err = db.Delete([]byte(string(part[1])), nil)
        //fmt.Print("\nError : ",err)        
	if(err != nil) {
	fmt.Fprintf(conn,"ERR_FILE_NOT_FOUND\r\n")
	} else {
	delete (version, part[1])
	delete (expiry,  part[1])
	fmt.Fprintf(conn,"OK\r\n")
	}
        mutex.Unlock()
	} else {
	fmt.Fprintf(conn,"ERR_CMD_ERR\r\n")
	}
	} else {
	fmt.Fprintf(conn,"ERR_CMD_ERR\r\n")
	}
}}}	}
 

func main(){
        defer db.Close()
	serverMain()
}


