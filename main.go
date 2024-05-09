package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacciHandler(c *gin.Context) {
	start := time.Now() // Record the start time

	numStr := c.Param("number")
	num, err := strconv.Atoi(numStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid number"})
		return
	}

	result := fibonacci(num)

	elapsed := time.Since(start) // Calculate the elapsed time

	// Prepare response with result and latency information
	c.JSON(200, gin.H{
		"fibonacci": result,
		"latency":   elapsed.String(), // Convert elapsed time to string for response
	})
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{"response": "pong"})
}

type Graph struct {
	Nodes []*Node
}

type Node struct {
	ID    int
	Edges []*Node
}

func generateGraph(c *gin.Context) {
	start := time.Now() // Record the start time

	numStr := c.Param("nodes")
	num, err := strconv.Atoi(numStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid number"})
		return
	}

	numNodes, numEdgesPerNode := num, num*10

	graph := &Graph{}
	for i := 0; i < numNodes; i++ {
		node := &Node{ID: i}
		graph.Nodes = append(graph.Nodes, node)
	}

	for _, node := range graph.Nodes {
		for j := 0; j < numEdgesPerNode; j++ {
			edgeIdx := rand.Intn(numNodes)
			node.Edges = append(node.Edges, graph.Nodes[edgeIdx])
		}
	}

	elapsed := time.Since(start) // Calculate the elapsed time

	// Prepare response with result and latency information
	c.JSON(200, gin.H{
		"result":  "done",
		"latency": elapsed.String(), // Convert elapsed time to string for response
	})
}

func main() {
	r := gin.Default()

	r.GET("/fibonacci/:number", fibonacciHandler)
	r.GET("/ping", ping)
	r.GET("/graph/:nodes", generateGraph)

	fmt.Println("Server listening on port 8080...")
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
