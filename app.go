package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    // Create a new Gin router
    router := gin.Default()

    // Define a route for the root endpoint
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, World!",
        })
    })

    // Define a route for a GET request to /users
    router.GET("/users", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "users": []string{"Alice", "Bob", "Charlie"},
        })
    })

    // Define a route for a POST request to /users
    router.POST("/users", func(c *gin.Context) {
        // Parse the request body into a struct
        var newUser struct {
            Name string `json:"name"`
        }
        if err := c.BindJSON(&newUser); err != nil {
            c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Add the new user to the database
        // ...

        // Return a success response
        c.JSON(http.StatusOK, gin.H{
            "message": "User created successfully",
        })
    })

    // Start the server on port 8080
    router.Run(":8080")
}
