package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

type Greeting struct {
	Body struct {
		Message string `json:"message"`
	}
}

type UserRequest struct {
	Body struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
}

type UserResponse struct {
	Body struct {
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		ID        string    `json:"id"`
		CreatedAt time.Time `json:"created_at"`
	}
}

func main() {
	router := http.NewServeMux()
	api := humago.New(router, huma.DefaultConfig("Greeting API", "v1.0.0"))

	huma.Register(api, huma.Operation{
		OperationID: "greeting",
		Method:      http.MethodGet,
		Path:        "/greeting/{name}",
		Summary:     "Greet someone",
		Description: "This endpoint greets someone by name.",
		Tags:        []string{"greeting"},
	}, func(ctx context.Context, i *struct {
		Name string `path:"name" maxLength:"30" example:"leon" doc:"Name to greet"`
	}) (*Greeting, error) {
		resp := &Greeting{}
		resp.Body.Message = "Hello, " + i.Name + "!"
		return resp, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "user-create",
		Method:      http.MethodPost,
		Path:        "/users",
		Summary:     "create user",
		Description: "This endpoint creates a user.",
		Tags:        []string{"user"},
	}, func(ctx context.Context, i *UserRequest) (*UserResponse, error) {
		resp := &UserResponse{}
		resp.Body.Name = i.Body.Name
		resp.Body.Email = i.Body.Email
		resp.Body.ID = "123"
		resp.Body.CreatedAt = time.Now()
		return resp, nil
	})

	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
