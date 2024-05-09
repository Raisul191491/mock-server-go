package main

import (
	"fmt"
	"time"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacciHandler(val int) {
	start := time.Now() // Record the start time

	result := fibonacci(val)

	elapsed := time.Since(start) // Calculate the elapsed time

	fmt.Println("The ans is for %d is : ", result)
	fmt.Println("latency :%v", elapsed)
}

// func ping() {
// 	c.JSON(200, gin.H{"response": "pong"})
// }

// type Graph struct {
// 	Nodes []*Node
// }

// type Node struct {
// 	ID    int
// 	Edges []*Node
// }

// func generateGraph() {
// 	start := time.Now() // Record the start time

// 	numStr := c.Param("nodes")
// 	num, err := strconv.Atoi(numStr)
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": "Invalid number"})
// 		return
// 	}

// 	numNodes, numEdgesPerNode := num, num*10

// 	graph := &Graph{}
// 	for i := 0; i < numNodes; i++ {
// 		node := &Node{ID: i}
// 		graph.Nodes = append(graph.Nodes, node)
// 	}

// 	for _, node := range graph.Nodes {
// 		for j := 0; j < numEdgesPerNode; j++ {
// 			edgeIdx := rand.Intn(numNodes)
// 			node.Edges = append(node.Edges, graph.Nodes[edgeIdx])
// 		}
// 	}

// 	elapsed := time.Since(start) // Calculate the elapsed time

// 	// Prepare response with result and latency information
// 	c.JSON(200, gin.H{
// 		"result":  "done",
// 		"latency": elapsed.String(), // Convert elapsed time to string for response
// 	})
// }

func main() {
	testArr := []int{30, 45, 47, 50}
	for _, val := range testArr {
		fibonacci(val)
	}
	// r := gin.Default()

	// r.GET("/fibonacci/:number", fibonacciHandler)
	// r.GET("/ping", ping)
	// r.GET("/graph/:nodes", generateGraph)

	// fmt.Println("Server listening on port 8080...")
	// err := r.Run(":8080")
	// if err != nil {
	// 	fmt.Println("Error starting server:", err)
	// }
}
