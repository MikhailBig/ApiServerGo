package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", pong)
	r.GET("/user", getUser)
	r.POST("/user", postUser)
	r.POST("/getSymbols", getSymbols)
	r.POST("/getQuotes", getQuotes)
	r.Run(":9090")
}

var apiKey = "ILyuhLSC4DLFS9kSF2tvshJYMLQhCwbj"

func getQuotes(c *gin.Context) {

	s := struct {
		Symbol string `json:"symbol"`
	}{}

	c.BindJSON(&s)

	if len(s.Symbol) > 0 {
		symbols := []string{s.Symbol}
		quotes := GetQuotes(symbols, apiKey)
		c.JSON(200, gin.H{
			"quotes": quotes,
		})
	}
}
func getSymbols(c *gin.Context) {
	symbol_list := GetSymbols(apiKey)
	c.JSON(200, gin.H{
		"symbols": symbol_list,
	})
}

//pong проверка на пинг
func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

//getUser принимает get запрос с данными пользователя
func getUser(c *gin.Context) {
	id := c.Query("id")
	user := struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}{id, "Mikhail"}
	c.JSON(200, gin.H{
		"user": user,
	})
}

//postUser принимает post запрос с данными пользователя
func postUser(c *gin.Context) {
	user := struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}{}

	c.BindJSON(&user)

	user.Name = "Mikhail"

	c.JSON(200, gin.H{
		"user": user,
	})
}
