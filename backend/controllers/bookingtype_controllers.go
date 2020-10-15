package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/B6111427/app/ent"
	"github.com/B6111427/app/ent/bookingtype"
	"github.com/gin-gonic/gin"
)

// BookingTypeController defines the struct for the bookingtype controller
type BookingTypeController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateBookingType handles POST requests for adding bookingtype entities
// @Summary Create bookingtype
// @Description Create bookingtype
// @ID create-bookingtype
// @Accept   json
// @Produce  json
// @Param bookingtype body ent.Bookingtype true "Bookingtype entity"
// @Success 200 {object} ent.Bookingtype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookingtypes [post]
func (ctl *BookingTypeController) CreateBookingType(c *gin.Context) {
	obj := ent.Bookingtype{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "bookingtype binding failed",
		})
		return
	}

	bt, err := ctl.client.Bookingtype.
		Create().
		SetBOOKTYPENAME(obj.BOOKTYPENAME).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, bt)
}

// GetBookingType handles GET requests to retrieve a bookingtype entity
// @Summary Get a bookingtype entity by ID
// @Description get bookingtype by ID
// @ID get-bookingtype
// @Produce  json
// @Param id path int true "Bookingtype ID"
// @Success 200 {object} ent.Bookingtype
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookingtypes/{id} [get]
func (ctl *BookingTypeController) GetBookingType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	bt, err := ctl.client.Bookingtype.
		Query().
		Where(bookingtype.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, bt)
}

// ListBookingType handles request to get a list of bookingtype entities
// @Summary List bookingtype entities
// @Description list bookingtype entities
// @ID list-bookingtype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Bookingtype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookingtypes [get]
func (ctl *BookingTypeController) ListBookingType(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	bookingtypes, err := ctl.client.Bookingtype.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, bookingtypes)
}

// DeleteBookingType handles DELETE requests to delete a bookingtype entity
// @Summary Delete a bookingtype entity by ID
// @Description get bookingtype by ID
// @ID delete-bookingtype
// @Produce  json
// @Param id path int true "Bookingtype ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookingtypes/{id} [delete]
func (ctl *BookingTypeController) DeleteBookingType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Bookingtype.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateBookingType handles PUT requests to update a bookingtype entity
// @Summary Update a bookingtype entity by ID
// @Description update bookingtype by ID
// @ID update-bookingtype
// @Accept   json
// @Produce  json
// @Param id path int true "Bookingtype ID"
// @Param bookingtype body ent.Bookingtype true "Bookingtype entity"
// @Success 200 {object} ent.Bookingtype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookingtypes/{id} [put]
func (ctl *BookingTypeController) UpdateBookingType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Bookingtype{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "bookingtype binding failed",
		})
		return
	}
	obj.ID = int(id)
	bt, err := ctl.client.Bookingtype.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, bt)
}

// NewBookingTypeController creates and registers handles for the bookingtype controller
func NewBookingTypeController(router gin.IRouter, client *ent.Client) *BookingTypeController {
	uc := &BookingTypeController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitBookingTypeController registers routes to the main engine
func (ctl *BookingTypeController) register() {
	bookingtypes := ctl.router.Group("/bookingtypes")

	bookingtypes.GET("", ctl.ListBookingType)

	// CRUD
	bookingtypes.POST("", ctl.CreateBookingType)
	bookingtypes.GET(":id", ctl.GetBookingType)
	bookingtypes.PUT(":id", ctl.UpdateBookingType)
	bookingtypes.DELETE(":id", ctl.DeleteBookingType)
}
