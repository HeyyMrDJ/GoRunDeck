package main
import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	//"crypto/tls"
	"io/ioutil"
	//"bytes"
	//"encoding/json"
	//"strings"
)

func main() {

	TestAuthentication2()
}

func TestAuthentication(){
	const (
		baseURL = "http://192.168.86.25:4440"
	)
	
	var (
		username = "admin"
		password = "admin"
	)
	type App struct {
		Client *http.Client
	}

	jar, _ := cookiejar.New(nil)
	app := App{Client: &http.Client{Jar: jar}}
	client := app.Client
	loginURL := baseURL + "/j_security_check"
	data := url.Values{
		"j_username": {username},
		"j_password": {password},
	}
	response, err := client.PostForm(loginURL, data)
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


func TestAuthentication2() {
	jar, _ := cookiejar.New(nil)
	app := &http.Client{Jar: jar}
	response, err := app.PostForm("http://192.168.86.25:4440/j_security_check?j_username=admin&j_password=admin", nil)
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

