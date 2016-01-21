package main

//package assignment1

import (
     "testing"
     "time" 
     "fmt"
     "os"
     "net"
     "bufio"
     
     )


 func TestMain(m *testing.M) {
       go serverMain()   // launch the server as a goroutine.
       time.Sleep(1 * time.Second) 
 }
 
 func TestCase1(t * testing.T) {
 
        ip_address, err := net.ResolveTCPAddr("tcp4","localhost:8080")
	//fmt.Print("\nIP Address : ",ip_address)
	if err!=nil {
	t.Errorf(err)	
	} else {
	conn, err  := net.DialTCP("tcp",nil, ip_address)
	if err!=nil {
	t.Errorf(err)	
	} else {
        reader:= bufio.NewReader(conn)
        
        fmt.Fprintf(conn, "write test.txt %v %v \r\n%v\r\n",21,20,"engineering the cloud")
        data, err := reader.ReadString('\n')
	if (err != nil) {
	t.Errorf(err)	
	} else {
	if (!(strings.Compare(string(data[:],"OK 1 \r\n"))==0)) {
        t.Errorf("Not Ok")		
        }
        
        
        
        
        
        
        
        
        }
 
// t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
 }}
 
 func TestCase2(t * testing.T) {
 
// t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
 }
 
 
 func TestCase3(t * testing.T) {
 
// t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
 }
 
 func TestCase4(t * testing.T) {
 
// t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
 }
 
 
 func TestCase5(t * testing.T) {
 
// t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
 }
 
 
 
