package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
  
	// Query string parameters are parsed using the existing underlying request object.
	//sendsms?username=te&password=te&from=72120&to=4673121231&text=hej
	// The request responds to a url matching:  /sendsms?username=hej&password=test&from
	//sendsms?username=te&password=te&from=72120&to=4673121231&text=hej
	router.POST("/sendsms", func(c *gin.Context) {
		//username := c.DefaultQuery("username", "Guest")
		
		//fmt.Println("Password: %s %s", username, password);
		var json Auth
		fmt.Println("Password: %s %s", json.Username, json.Password);
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//läser från filen med accounts för att verifiera
		/*data, err := os.ReadFile("/Users/andols/code_camp22/cmd/rest/message.txt")
		if err != nil {
			log.Panicf("failed reading data from file: %s", err)
		}

		c.String(http.StatusOK, "%s", data)*/
		if json.Username !="tjena" || json.Password !="test" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		/*username := c.Query("username")
		password := c.Query("password") // shortcut for c.Request.URL.Query().Get("password")
		from := c.Query("from") // shortcut for c.Request.URL.Query().Get("from")
		to := c.Query("to") // shortcut for c.Request.URL.Query().Get("to")
		text := c.Query("text") // shortcut for c.Request.URL.Query().Get("text")*/
		//gör så att man skriver ut lösenordet men krypterat
		//lägg till authentication. Modifiera login funktionen
	
		// Use the following code if you need to write the logs to file and console at the same time.
		// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	
		//router := gin.Default()
		/*router.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})*/
		var test = ("username : " + json.Username+", password : " +  json.Password+", from : " +  json.From+", to: " +json.To +", text : " +json.Text +"\n")

		f, err := os.OpenFile("/Users/andols/code_camp22/cmd/rest/message.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := f.Write([]byte(test)); err != nil {
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}


		/*file, err := os.Create("messages.txt")
		if err != nil {
			return
		}
		defer file.Close()*/
		//file.WriteString(test)
	    c.String(http.StatusOK, "%s %s %s %s %s ", json.Username, json.Password, json.From, json.To, json.Text)
	})

	router.GET("/getAll", func(c *gin.Context) {
		data, err := os.ReadFile("/Users/andols/code_camp22/cmd/rest/message.txt")
		if err != nil {
			log.Panicf("failed reading data from file: %s", err)
		}
		c.String(http.StatusOK, "%s", data)
	})

	router.POST("/createUser", func(c *gin.Context) {
		data, err := os.ReadFile("/Users/andols/code_camp22/cmd/rest/accounts.json")
		if err != nil {
			log.Panicf("failed reading data from file: %s", err)
		}
		var jsonUser User
		if err := c.ShouldBindJSON(&jsonUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	// define slice of Identification
		var jsonUsers []User

			// Unmarshall it
		if err := json.Unmarshal([]byte(data), &jsonUsers); err != nil {
			log.Println(err)
		}

		jsonUsers = append(jsonUsers, User{Username: jsonUser.Username, Password: jsonUser.Password, Name: jsonUser.Name, Age: jsonUser.Age})
		fmt.Println("jsonUsers1: ", jsonUsers)
		// iterating it
		for _, v := range jsonUsers {
			fmt.Println(v)
		}
		fmt.Println("jsonUsers2: ", jsonUsers)

		// now Marshal it
		result, err := json.Marshal(jsonUsers)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("result: ",result)

	// now result has your targeted JSON structure
		fmt.Println(string(result))
		/*if json.Username != "tjaba" || json.Password != "tjaba" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}*/
		//var test = ("username : " + jsonUser.Username+", password : " +  jsonUser.Password+", name : " +  jsonUser.Name+", age: " +jsonUser.Age+"\n")
		//fmt.Printf("Username: " +jsonUser.Username)
		//jsonString := User{Username: jsonUser.Username, Password: jsonUser.Password, Name: jsonUser.Name, Age: jsonUser.Age}

		//jsonString := "{"username\":"+jsonUser.Username+", \"password\": "+jsonUser.Password+", \"name\": "+jsonUser.Name+", \"age\": "+jsonUser.Age+"}"
		/*j, err := json.Marshal(jsonString)
		if err != nil {
			log.Fatalln("Cannot marshal")
		}
		fmt.Println(string(j))*/
		/*s := User{}
	    json.Unmarshal([]byte(jsonString), &s)
		userJson, _ := json.Marshal(&s)

		fmt.Printf("Username: " + jsonString)*/

		f, err := os.OpenFile("/Users/andols/code_camp22/cmd/rest/accounts.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := f.Write([]byte(result)); err != nil {
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
		/*if _, err := f.Write([]byte("\n")); err != nil {
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}*/

		


		/*file, err := os.Create("messages.txt")
		if err != nil {
			return
		}
		defer file.Close()*/
		//file.WriteString(test)
	    //c.String(http.StatusOK, "%s %s %s %s %s ", json.Username, json.Password, json.From, json.To, json.Text)
	})
	router.Run(":8080")

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


  //skriva till fil (ny inloggning)
  //när man skickar request ska man läsa från filen och kolla om man hittar användaren där
  //delete all

  //tT
  //skapa en till REST Server.  (som gör en autentisering)
