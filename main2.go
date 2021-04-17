package main
import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"crypto/tls"
	"io/ioutil"
)

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

func (app *App) login()  {
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

func main() {
	jar, _ := cookiejar.New(nil)
	tr := &http.Transport{
	  TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	  }
	  app := App{
	  Client: &http.Client{Jar: jar,Transport: tr},
	}
	app.login()
}