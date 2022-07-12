package webapi

import (
	"github.com/raflesngln/MyProjectGo/internal/app/myshop/webhandler/shoppinghandler"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Route(r *gin.Engine) *gin.Engine {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/Profiles", Profiles)


	v1ShoppingList := r.Group("/v1/shoppinglist")
	{
		v1ShoppingList.POST("/", shoppinghandler.CreateHandler)
		v1ShoppingList.GET("/:id", shoppinghandler.ShowHandler)
		//v1ShoppingList.GET("/:id", Profiles)
		v1ShoppingList.PUT("/:id", shoppinghandler.PutHandler)
		v1ShoppingList.DELETE("/:id", shoppinghandler.DeleteHandler)
	}

	return r
}

func Profiles(c *gin.Context) {
	prod := []Product{{
		NamaDepan:    "rafles",
		NamaBelakang: "nainggolan",
	},
		{
			NamaDepan:    "budi",
			NamaBelakang: "Houten",
		},
	}
	c.JSON(200, gin.H{
		"data": prod,
	})
}


type Product struct {
	NamaDepan    string `json:"nama_depan"`
	NamaBelakang string `json:"nama_belakang"`
}