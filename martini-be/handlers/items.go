package handlers

import (
	"martini-be/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Item = models.Item

type ItemsHandler struct {
	db *gorm.DB
}

func CreateItemsHandler(db *gorm.DB) *ItemsHandler {
	return &ItemsHandler{
		db: db,
	}
}

func (h *ItemsHandler) GetItems(c *gin.Context) {
	var items []Item
	h.db.Order("id ASC").Find(&items)

	c.JSON(http.StatusOK, items)
}

func (h *ItemsHandler) GetItem(c *gin.Context) {
	id := c.Param("id")

	var item Item
	h.db.First(&item, "id = ?", id)

	c.JSON(http.StatusOK, item)
}

func (h *ItemsHandler) CreateItem(c *gin.Context) {
	var item Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item.ID = 0
	if err := h.db.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func (h *ItemsHandler) UpdateItem(c *gin.Context) {
	id := c.Param("id")

	var item Item
	h.db.First(&item, "id = ?", id)

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.db.Save(&item)

	c.JSON(http.StatusOK, item)
}

func (h *ItemsHandler) DeleteItem(c *gin.Context) {
	id := c.Param("id")

	var item Item
	h.db.First(&item, "id = ?", id)

	h.db.Delete(&item)

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
}
