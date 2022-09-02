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
	router.POST("/sendsms", func(c *gin.Context) {
		var jsonAuth Auth
		if err := c.ShouldBindJSON(&jsonAuth); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		file, err := os.ReadFile("/Users/andols/code_camp22/cmd/rest/accounts.json")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		newUsers := make([]User, 0)
		err = json.Unmarshal([]byte(file), &newUsers)
		if err == nil {
			fmt.Printf("%+v\n", newUsers)
		} else {
			fmt.Println(err)
			fmt.Printf("%+v\n", newUsers)
		}

		fmt.Println("testar: " ,newUsers)
		authorized :=false
		for i := 0; i<len(newUsers); i++{
			if jsonAuth.Username == newUsers[i].Username && jsonAuth.Password ==newUsers[i].Password {
				authorized = true
				break
			}
		}
		if !authorized {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		var message = ("username : " + jsonAuth.Username+", password : " +  jsonAuth.Password+", from : " +  jsonAuth.From+", to: " +jsonAuth.To +", text : " +jsonAuth.Text +"\n")

		f, err := os.OpenFile("/Users/andols/code_camp22/cmd/rest/message.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := f.Write([]byte(message)); err != nil {
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	    c.String(http.StatusOK, "%s has sent a message from %s to %s saying: '%s' ", jsonAuth.Username, jsonAuth.From, jsonAuth.To, jsonAuth.Text)
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

		// now Marshal it
		result, err := json.Marshal(jsonUsers)
		if err != nil {
			log.Println(err)
		}
		// now result has targeted JSON structure
		fmt.Println(string(result))

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
	    c.String(http.StatusOK, "%s has been created with password: %s, name: %s, Age: %s", jsonUser.Username, jsonUser.Password, jsonUser.Name, jsonUser.Age)
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
type UserResponse struct {
	Collection []User
}

  //skriva till fil (ny inloggning)
  //när man skickar request ska man läsa från filen och kolla om man hittar användaren där
  //delete all

  //tT
  //skapa en till REST Server.  (som gör en autentisering)
