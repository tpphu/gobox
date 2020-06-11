package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tpphu/gobox/example/model"
)

type Product struct {
	DB *gorm.DB `inject:"db"`
}


func (n *Product) Get(c *gin.Context) {
	id := c.Param("id")
	p := model.Product{}
	n.DB.Find(&p, id)
	c.JSON(200, p)
}