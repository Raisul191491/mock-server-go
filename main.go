package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacciHandler(val int, file *os.File) {
	start := time.Now()

	result := fibonacci(val)

	elapsed := time.Since(start)

	output := fmt.Sprintf("The ans for %d is: %d\n", val, result)
	output += fmt.Sprintf("Latency: %v\n\n", elapsed)

	file.WriteString(output)
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
// 	start := time.Now()

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
	file, err := os.Create("fibonacci_performance_output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	testArr := []int{46, 47, 48, 50}

	var wg sync.WaitGroup

	for _, val := range testArr {
		wg.Add(1)
		go func(num int, file *os.File) {
			defer wg.Done()
			fibonacciHandler(num, file)
			fmt.Println(" ")
		}(val, file)
	}

	wg.Wait()
	fmt.Println("All goroutines finished. Output written to fibonacci_output.txt.")
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
