package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Results any    `json:"results"`
}

type Users struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var nextId = 1
var nextAccId = 1

var UserList []Users
var AccountList []Users

func main() {
	r := gin.Default()

	r.POST("/register", func(ctx *gin.Context) {
		var data Users
		err := ctx.ShouldBind(&data)

		if err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Register Failed",
			})
			return
		}

		for i := 0; i < len(AccountList); i++ {
			if AccountList[i].Email == data.Email {
				ctx.JSON(400, Response{
					Success: false,
					Message: "Email already exist",
				})
				return
			}
		}

		data.Id = nextAccId
		AccountList = append(AccountList, data)
		nextAccId++
		ctx.JSON(200, Response{
			Success: true,
			Message: "Register Success",
		})
	})

	r.POST("/login", func(ctx *gin.Context) {
		var data Users
		err := ctx.ShouldBind(&data)

		if err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Login Failed",
			})
			return
		}
		if data.Email == "" || data.Password == "" {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Email and Password cannot blank",
			})
			return
		}

		for i := 0; i < len(AccountList); i++ {
			if AccountList[i].Email == data.Email {
				if AccountList[i].Password == data.Password {
					ctx.JSON(200, Response{
						Success: true,
						Message: "Login Successfuly",
					})
				} else {
					ctx.JSON(400, Response{
						Success: false,
						Message: "Password Incorrect",
					})
				}
			} else {
				ctx.JSON(400, Response{
					Success: false,
					Message: "Email Incorrect",
				})
			}
		}
	})

	r.POST("/users", func(ctx *gin.Context) {
		var data Users
		err := ctx.ShouldBind(&data)

		if err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "User create Failed",
			})
			return
		}
		for i := 0; i < len(UserList); i++ {
			if UserList[i].Email == data.Email {
				ctx.JSON(400, Response{
					Success: false,
					Message: "Email already exist",
				})
				return
			}
		}
		data.Id = nextId
		UserList = append(UserList, data)
		nextId++
		ctx.JSON(200, Response{
			Success: true,
			Message: "User created successfully",
		})
	})

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, Response{
			Success: true,
			Message: "List of users",
			Results: UserList,
		})
	})

	r.GET("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		for i := range UserList {
			if fmt.Sprint(UserList[i].Id) == id {
				ctx.JSON(200, Response{
					Success: true,
					Message: fmt.Sprintf("Hello User %s", id),
					Results: UserList[i],
				})
				return
			}
		}
		ctx.JSON(404, Response{
			Success: false,
			Message: fmt.Sprintf("User %s not found", id),
		})
	})

	r.PATCH("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		var newData Users
		err := ctx.ShouldBind(&newData)

		if err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "Input error",
			})
			return
		}
		for i := range UserList {
			if fmt.Sprint(UserList[i].Id) == id {
				if newData.Email != "" {
					for j := range UserList {
						if UserList[j].Email == newData.Email && UserList[j].Id != newData.Id {
							ctx.JSON(400, Response{
								Success: false,
								Message: "Email already registered",
							})
							return
						}
					}
					UserList[i].Email = newData.Email
				}
				if newData.Password != "" {
					UserList[i].Password = newData.Password
				}
				ctx.JSON(200, Response{
					Success: true,
					Message: "User data updated",
				})
				return
			}
			ctx.JSON(400, Response{
				Success: false,
				Message: "Input error",
			})
		}
		ctx.JSON(400, Response{
			Success: false,
			Message: "User not found",
		})
	})

	r.DELETE("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		for i := range UserList {
			if fmt.Sprint(UserList[i].Id) == id {
				UserList = append(UserList[:i], UserList[i+1:]...)
				ctx.JSON(200, Response{
					Success: true,
					Message: "User deleted",
				})
			}
		}
	})

	r.Run("localhost:8989")
}
