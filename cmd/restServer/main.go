package main

import (
	"fmt"
	"net/http"
	"log"
	"strings"
	"io/ioutil"
	"encoding/json"
	"github.com/gin-gonic/gin"
)
func main() {
	router := gin.Default()
	router.POST("/sendsms", func(c *gin.Context) {
		sendMessage(c)
	})

	router.GET("/getAllMessages/", func(c *gin.Context) {
		getAllMessages(c)
	})
	router.GET("/getAllAccounts", func(c *gin.Context) {
		getAllAccounts(c)
	})
	router.POST("/createUser", func(c *gin.Context) {
		createAccount(c)
	})
	router.Run(":8080")

}

func getAllMessages(c *gin.Context){
	username := c.Query("username")
	password := c.Query("password")
	url := "http://localhost:8060/getAllMessages?username="+username+"&password="+password
	method := "GET"
  
	payload := strings.NewReader(``)
  
	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)
  
	if err != nil {
	  fmt.Println(err)
	  return
	}

	res, err := client.Do(req)
	if err != nil {
	  fmt.Println(err)
	  return
	}
	defer res.Body.Close()
  
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
	  fmt.Println(err)
	  return
	}
	fmt.Println(string(body))
	c.String(http.StatusOK, "%s", string(body))
}

func getAllAccounts(c *gin.Context){
	username := c.Query("username")
    	password := c.Query("password")
		url := "http://localhost:8060/getAllAccounts?username="+username+"&password="+password
		method := "GET"
	  
		payload := strings.NewReader(``)
	  
		client := &http.Client {
		}
		req, err := http.NewRequest(method, url, payload)
	  
		if err != nil {
		  fmt.Println(err)
		  return
		}

		res, err := client.Do(req)
		if err != nil {
		  fmt.Println(err)
		  return
		}
		defer res.Body.Close()
	  
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
		  fmt.Println(err)
		  return
		}
		fmt.Println(string(body))
		c.String(http.StatusOK, "%s", string(body))
}

func sendMessage(c *gin.Context){
	var jsonAuth Auth
	if err := c.ShouldBindJSON(&jsonAuth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	result, err := json.Marshal(jsonAuth)
	if err != nil {
		log.Println(err)
	}
	url := "http://localhost:8060/sendsms"
	method := "POST"

	payload := strings.NewReader(string(result))

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	
	defer res.Body.Close()
	fmt.Println("res: " , res.Status)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	c.String(http.StatusOK, string(body))
}

func createAccount(c *gin.Context){

		url := "http://localhost:8060/createUser"
		method := "POST"
		var jsonUser User
		if err := c.ShouldBindJSON(&jsonUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// now Marshal it
		result, err := json.Marshal(jsonUser)
		if err != nil {
			log.Println(err)
		}
		// now result has targeted JSON structure
		fmt.Println(string(result))

		fmt.Println("bytes.NewBuffer(result): " ,string(result))

		payload := strings.NewReader(string(result))
	  
		client := &http.Client {
		}
		req, err := http.NewRequest(method, url, payload)
	  
		if err != nil {
		  fmt.Println(err)
		  return
		}
		req.Header.Add("Content-Type", "application/json")
	  
		res, err := client.Do(req)
		if err != nil {
		  fmt.Println(err)
		  return
		}
		defer res.Body.Close()
	  
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
		  fmt.Println(err)
		  return
		}
		fmt.Println(string(body))
		//här ska jag ta emot response från andra servern
		c.String(http.StatusOK,string(body))

}

type Auth struct {
	Username string `json: "username" binding:"required"`
	Password string `json: "password" binding:"required"`
	From string `json: "from"`
	To string `json: "to"`
	Text string `json: "text"`
}

type User struct {
	Username string `json: "username" binding:"required"`
	Password string `json: "password" binding:"required"`
	Name string `json: "name"`
	Age string `json: "age"`
}

type UserResponse struct {
	Collection []User
}

  //skriva till fil (ny inloggning)
  //när man skickar request ska man läsa från filen och kolla om man hittar användaren där
  //delete all

  //tT
  //skapa en till REST Server.  (som gör en autentisering)
