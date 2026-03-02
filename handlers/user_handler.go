package handlers

import (
	"fmt"
	"koda-b6-backend1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var data models.Users
	err := ctx.ShouldBind(&data)

	if err != nil {
		ctx.JSON(400, models.Response{
			Success: false,
			Message: "Register Failed",
		})
		return
	}
	if data.Email == "" || data.Password == "" {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Email and Password cannot blank",
		})
		return
	}

	for i := 0; i < len(models.AccountList); i++ {
		if models.AccountList[i].Email == data.Email {
			ctx.JSON(400, models.Response{
				Success: false,
				Message: "Email already exist",
			})
			return
		}
	}

	hashedPassword, err := HashPassword(data.Password)
	if err != nil {
		ctx.JSON(500, models.Response{
			Success: false,
			Message: "Failed to hash password",
		})
		return
	}

	data.Password = hashedPassword
	data.Id = models.NextAccId
	models.AccountList = append(models.AccountList, data)
	models.NextAccId++
	ctx.JSON(200, models.Response{
		Success: true,
		Message: "Register Success",
	})
}

func Login(ctx *gin.Context) {
	var data models.Users
	err := ctx.ShouldBind(&data)

	if err != nil {
		ctx.JSON(400, models.Response{
			Success: false,
			Message: "Login Failed",
		})
		return
	}
	if data.Email == "" || data.Password == "" {
		ctx.JSON(400, models.Response{
			Success: false,
			Message: "Email and Password cannot blank",
		})
		return
	}

	for i := 0; i < len(models.AccountList); i++ {
		if models.AccountList[i].Email == data.Email {
			if VerifyPassword(models.AccountList[i].Password, data.Password) {
				ctx.JSON(http.StatusOK, models.Response{
					Success: true,
					Message: "Login successful",
				})
			} else {
				ctx.JSON(400, models.Response{
					Success: false,
					Message: "Password Incorrect",
				})
			}
			return
		}
		ctx.JSON(400, models.Response{
			Success: false,
			Message: "Email Incorrect",
		})

	}
}

func CreateUser(ctx *gin.Context) {
	var data models.Users
	err := ctx.ShouldBind(&data)

	if err != nil {
		ctx.JSON(400, models.Response{
			Success: false,
			Message: "User create Failed",
		})
		return
	}
	for i := 0; i < len(models.UserList); i++ {
		if models.UserList[i].Email == data.Email {
			ctx.JSON(400, models.Response{
				Success: false,
				Message: "Email already exist",
			})
			return
		}
	}
	data.Id = models.NextId
	models.UserList = append(models.UserList, data)
	models.NextId++
	ctx.JSON(200, models.Response{
		Success: true,
		Message: "User created successfully",
	})
}

func GetUser(ctx *gin.Context) {
	ctx.JSON(200, models.Response{
		Success: true,
		Message: "List of users",
		Results: models.UserList,
	})
}

func GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	for i := range models.UserList {
		if fmt.Sprint(models.UserList[i].Id) == id {
			ctx.JSON(200, models.Response{
				Success: true,
				Message: fmt.Sprintf("Hello User %s", id),
				Results: models.UserList[i],
			})
			return
		}
	}
	ctx.JSON(404, models.Response{
		Success: false,
		Message: fmt.Sprintf("User %s not found", id),
	})
}

func EditUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var newData models.Users
	err := ctx.ShouldBind(&newData)

	if err != nil {
		ctx.JSON(400, models.Response{
			Success: false,
			Message: "Input error",
		})
		return
	}
	for i := range models.UserList {
		if fmt.Sprint(models.UserList[i].Id) == id {
			if newData.Email != "" {
				for j := range models.UserList {
					if models.UserList[j].Email == newData.Email && models.UserList[j].Id != newData.Id {
						ctx.JSON(400, models.Response{
							Success: false,
							Message: "Email already registered",
						})
						return
					}
				}
				models.UserList[i].Email = newData.Email
			}
			if newData.Password != "" {
				models.UserList[i].Password = newData.Password
			}
			ctx.JSON(200, models.Response{
				Success: true,
				Message: "User data updated",
			})
			return
		}
		ctx.JSON(400, models.Response{
			Success: false,
			Message: "Input error",
		})
	}
	ctx.JSON(400, models.Response{
		Success: false,
		Message: "User not found",
	})
}

func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	for i := range models.UserList {
		if fmt.Sprint(models.UserList[i].Id) == id {
			models.UserList = append(models.UserList[:i], models.UserList[i+1:]...)
			ctx.JSON(200, models.Response{
				Success: true,
				Message: "User deleted",
			})
		}
	}
}
