package rest

// import (
// 	"github.com/hadihammurabi/belajar-go-rest-api/api/rest/dto"
// 	"github.com/hadihammurabi/belajar-go-rest-api/api/rest/response"
// 	"github.com/hadihammurabi/belajar-go-rest-api/internal/entity"

// 	"github.com/gofiber/fiber/v2"
// )

// type auth struct {
// 	*APIRest
// }

// // Auth func
// func Auth(api *APIRest) {
// 	router := api.HTTP.Group("/auth")
// 	router.Post("/login", api.Login)
// 	router.Get("/me", api.Middlewares.Auth, api.Me)
// }

// // Login func
// func (api APIRest) Login(c *fiber.Ctx) error {
// 	input := &dto.UserLoginRequest{}
// 	if err := c.BodyParser(input); err != nil {
// 		return response.Fail(c, err)
// 	}

// 	user := &entity.User{
// 		Email:    input.Email,
// 		Password: input.Password,
// 	}

// 	token, err := api.Service.Auth.Login(c.Context(), user)
// 	if err != nil {
// 		return response.Fail(c, "invalid credentials")
// 	}

// 	return response.Ok(c, &dto.UserLoginResponse{
// 		Token: token,
// 		Type:  "Bearer",
// 	})
// }

// // Me func
// func (api APIRest) Me(c *fiber.Ctx) error {
// 	fromLocals := c.Locals("user").(*entity.User)
// 	return response.Ok(c, fromLocals)
// }
