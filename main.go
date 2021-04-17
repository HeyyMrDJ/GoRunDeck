package main

import (
	"fmt"
	"net"
	"time"
	"os"
	"net/http"
	//"bytes"
	//"encoding/json"
	"io/ioutil"
	"net/url"
	"net/http/cookiejar"
	//"crypto/tls"
)




type RundeckServer struct {
	hostname string
	portnumber string
	username string
	password string

}



func testconnection(hostname string){

	portname := "4440"
	seconds := 2
	timeout := time.Duration(seconds) * time.Second

	conn, err := net.DialTimeout("tcp", hostname+":"+portname, timeout)

	if err != nil {

		fmt.Println(err)
	}

	if conn != nil {
		fmt.Printf(hostname + " is available \n")
	}

}

func Authenticate(server RundeckServer){
	baseurl := "http://" + server.hostname+  ":" + server.portnumber
	loginurl := baseurl + "/j_security_check"
	jar, _ := cookiejar.New(nil)
	app := &http.Client{Jar: jar}
	body := url.Values{
		"j_username": {server.username},
		"j_password": {server.password},
	}
	response, err := app.PostForm(loginurl, body)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	fmt.Println("HTTP Response Status:", response.StatusCode, http.StatusText(response.StatusCode))
	if response.StatusCode >= 200 && response.StatusCode <= 299 {
		fmt.Println("HTTP Status is in the 2xx range")
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		} else {
		fmt.Println("Argh! Broken")
		}
	}
	

func main(){
	
	rundeck := RundeckServer{os.Args[1], "4440", "admin", "admin"}
	testconnection(rundeck.hostname)
	Authenticate(rundeck)
	
}