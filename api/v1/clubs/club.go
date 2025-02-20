package clubs

import (
	"context"

	"github.com/gin-gonic/gin"
	"noan.dev/uniklub/constants"
	"noan.dev/uniklub/models/club"
)

func Add(ctx context.Context) func(*gin.Context) {
	return func(c *gin.Context) {
		type clubRequestPost struct {
			Name string `json:"name" binding:"required"`
		}

		var request clubRequestPost
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, constants.CreateErrorMessage(err.Error()))
			return
		}

		result := club.Create(ctx, club.ClubCreationFields{
			Name: request.Name,
		})
		if result == nil {
			c.JSON(500, constants.CreateErrorMessage(constants.InternalError))
			return
		}

		c.JSON(200, result)
	}
}

func Get(ctx context.Context) func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, club.FindAll(ctx))
	}
}

func GetById(ctx context.Context) func(*gin.Context) {
	return func(c *gin.Context) {
		type clubUri struct {
			Id int `uri:"id" binding:"required"`
		}

		var uri clubUri
		if err := c.ShouldBindUri(&uri); err != nil {
			c.JSON(400, constants.CreateErrorMessage(constants.WrongFormat))
			return
		}

		result := club.Find(ctx, club.ClubPrimaryKey{
			Id: uri.Id,
		})
		if result == nil {
			c.JSON(404, constants.CreateErrorMessage("club not found"))
			return
		}

		c.JSON(200, result)
	}
}

func Update(ctx context.Context) func(*gin.Context) {
	return func(c *gin.Context) {
		type clubUri struct {
			Id int `uri:"id" binding:"required"`
		}
		type clubRequestPost struct {
			Name string `json:"name" binding:"required"`
		}

		var uri clubUri
		if err := c.ShouldBindUri(&uri); err != nil {
			c.JSON(400, constants.CreateErrorMessage(constants.WrongFormat))
			return
		}

		var request clubRequestPost
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, constants.CreateErrorMessage(err.Error()))
			return
		}

		result := club.Find(ctx, club.ClubPrimaryKey{
			Id: uri.Id,
		})
		if result == nil {
			c.JSON(404, constants.CreateErrorMessage("club not found"))
			return
		}

		c.JSON(200, club.Update(ctx, club.ClubPrimaryKey{
			Id: result.Id,
		}, club.ClubUpdateFields{
			Name: request.Name,
		}))
	}
}

func Remove(ctx context.Context) func(*gin.Context) {
	return func(c *gin.Context) {
		type clubUri struct {
			Id int `uri:"id" binding:"required"`
		}

		var uri clubUri
		if err := c.ShouldBindUri(&uri); err != nil {
			c.JSON(400, constants.CreateErrorMessage(constants.WrongFormat))
			return
		}

		result := club.Find(ctx, club.ClubPrimaryKey{
			Id: uri.Id,
		})
		if result == nil {
			c.JSON(404, constants.CreateErrorMessage("club not found"))
			return
		}

		if err := club.Delete(ctx, club.ClubPrimaryKey{
			Id: result.Id,
		}); err != nil {
			c.JSON(500, constants.CreateErrorMessage(constants.InternalError))
		}
		c.Status(200)
	}
}
