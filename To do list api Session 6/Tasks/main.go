package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keploy/go-sdk/integrations/kgin/v1"
	"github.com/keploy/go-sdk/keploy"
)

type task struct {
	ID       string `json:"id"`
	Task     string `json:"task"`
	Deadline string `json:"deadline"`
	Status   string `json:"status"`
}

var tasks = []task{
	{ID: "1", Task: "Do the Keploy Homework", Deadline: "Friday", Status: "Pending"},
	{ID: "2", Task: "Upload files on Github", Deadline: "Friday", Status: "Pending"},
	{ID: "3", Task: "Watch the lectures", Deadline: "Wednesady", Status: "Completed"},
	{ID: "4", Task: "Attend the Meeting", Deadline: "Thursday", Status: "Completed"},
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ // H is a shortcut for map[string]interface{}
		"instructions": "Add '/tasks' to the link",
	})
}

func getAllTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

func addTask(c *gin.Context) {
	var newTask task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Bad Request",
		})
		return
	}

	tasks = append(tasks, newTask)
	c.JSON(http.StatusCreated, newTask)
}

func editTask(c *gin.Context) {
	id := c.Param("id")

	var editTask task

	if err := c.ShouldBindJSON(&editTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Bad Request",
		})
		return
	}
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Task = editTask.Task
			tasks[i].Deadline = editTask.Deadline
			tasks[i].Status = editTask.Status

			c.JSON(http.StatusOK, editTask)
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error":   true,
		"message": "Invalid",
	})
}

func removeTask(c *gin.Context) {
	id := c.Param("id")
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i],
				tasks[i+1:]...)

			c.JSON(http.StatusOK, gin.H{
				"message": "Item Deleted",
			})
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error":   true,
		"message": "Invalid",
	})
}

func main() {
	port := "8080"
	keploy := keploy.New(keploy.Config{
		App: keploy.AppConfig{
			Name: "Tasks",
			Port: port,
		},
		Server: keploy.ServerConfig{
			URL: "http://localhost:8081/api",
		},
	})
	router := gin.Default()
	kgin.GinV1(keploy, router)

	router.GET("/", home)
	router.GET("/tasks", getAllTasks)
	router.POST("/tasks", addTask)
	router.PUT("/tasks/:id", editTask)
	router.DELETE("/tasks/:id", removeTask)

	router.Run(":8080")
}
