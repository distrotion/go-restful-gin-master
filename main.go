package main

import (
	_ "encoding/json"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// cors "github.com/itsjamie/gin-cors"
)

type Gqlquery struct {
	Query string `json:"query"`
}

type AddScore struct {
	Address string `json:"address"`
	Score   string `json:"score"`
	Egg     string `json:"egg"`
}
type ReturnGetTop100 struct {
	Address string `json:"address"`
	Score   string `json:"score"`
	Egg     string `json:"egg"`
}
type GetRank struct {
	Address string `json:"address"`
}

type ReturnGetRank struct {
	Address string `json:"address"`
	Score   string `json:"score"`
	Egg     string `json:"egg"`
	Rank    int    `json:"rank"`
}

//----------------------------------------------- var
var (
	Collection = `MAIN`
)

//-----------------------------------------------

func main() {
	r := gin.Default()
	// r := gin.New()

	// r.Use(cors.Default())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {

		c.JSON(200, "topsc V0-11-1 revm add")
	})

	r.POST("/testinput01", func(c *gin.Context) {
		fmt.Println(c)
		var input interface{}
		c.BindJSON(&input)
		//-----------------------------------------
		fmt.Println(input)
		//=========================================
		// b, _ := json.Marshal(input)
		// s := string(b)

		c.JSON(200, input)

	})

	r.Run(":11001")
	// r.RunTLS(":11001", "./testdata/server.pem", "./testdata/server.key")
}
