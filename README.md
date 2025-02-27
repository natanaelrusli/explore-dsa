# Algorithm and Data Structure Implementations

A comprehensive collection of algorithms and data structures implemented in Go. This repository serves as both a learning resource and a reference for common computer science concepts.

## 📚 Table of Contents

1. [Sorting Algorithms](#sorting-algorithms)
2. [Searching Algorithms](#searching-algorithms)
3. [Data Structures](#data-structures)
4. [Dynamic Programming](#dynamic-programming)
5. [Graph Algorithms](#graph-algorithms)

## 🔄 Sorting Algorithms

### Merge Sort
A stable, divide-and-conquer sorting algorithm.

**Complexity:**
- Time: O(n log n)
- Space: O(n)

### Bubble Sort
A simple comparison-based sorting algorithm.

**Complexity:**
- Time: O(n²)
- Space: O(1)

## 🔍 Searching Algorithms

### Binary Search
An efficient algorithm for finding elements in a sorted array.

**Complexity:**
- Time: O(log n)
- Space: O(1)

### Linear Search
A simple search that checks each element sequentially.

**Complexity:**
- Time: O(n)
- Space: O(1)

## 📊 Data Structures

### Arrays and Slices
Basic operations and manipulations on arrays and slices.

### Linked Lists
Implementation of singly and doubly linked lists.

### Trees
- Binary Search Trees
- AVL Trees
- Red-Black Trees

### Hash Tables
Implementation of hash tables with collision resolution.

### Heaps
- Min Heap
- Max Heap

## 🧮 Dynamic Programming

### Classic Problems
- Fibonacci Sequence
- Longest Common Subsequence
- Knapsack Problem
- Matrix Chain Multiplication

## 🕸️ Graph Algorithms

### Graph Traversal
- Depth-First Search (DFS)
- Breadth-First Search (BFS)

### Shortest Path
- Dijkstra's Algorithm
- Bellman-Ford Algorithm

### Minimum Spanning Tree
- Kruskal's Algorithm
- Prim's Algorithm

## 🚀 Usage

Each algorithm and data structure is implemented in its own package. Here's an example of using the sorting algorithms:

```go
package main

import (
    "fmt"
    "github.com/yourusername/algo-ds/sorting"
    "github.com/yourusername/algo-ds/searching"
)

func main() {
    // Sorting example
    arr := []int{3, 5, 1, 3, 2, 8, 7}
    sorted := sorting.MergeSort(arr)
    fmt.Println("Sorted array:", sorted)

    // Binary search example
    target := 5
    index := searching.BinarySearch(sorted, target)
    fmt.Printf("Found %d at index: %d\n", target, index)
}
```

## 📊 Performance Comparison

| Algorithm      | Time Complexity (Best) | Time Complexity (Average) | Time Complexity (Worst) | Space Complexity |
|---------------|----------------------|-------------------------|----------------------|-----------------|
| Merge Sort    | O(n log n)          | O(n log n)             | O(n log n)          | O(n)           |
| Bubble Sort   | O(n)                | O(n²)                  | O(n²)               | O(1)           |
| Binary Search | O(1)                | O(log n)               | O(log n)            | O(1)           |
| Linear Search | O(1)                | O(n)                   | O(n)                | O(1)           |

## 🛠️ Project Structure

```
.
├── sorting/
│   ├── bubble_sort.go
│   ├── merge_sort.go
│   └── sort_test.go
├── searching/
│   ├── binary_search.go
│   ├── linear_search.go
│   └── search_test.go
├── datastructures/
│   ├── linkedlist/
│   ├── trees/
│   ├── heap/
│   └── hashtable/
├── dynamic/
│   ├── fibonacci.go
│   └── knapsack.go
├── graphs/
│   ├── dfs.go
│   ├── bfs.go
│   └── dijkstra.go
└── README.md
```

## 🧪 Testing

Each package includes comprehensive test cases. Run the tests using:

```bash
go test ./...
```

## 📈 Benchmarks

To run benchmarks:

```bash
go test -bench=. ./...
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

### Contributing Guidelines
1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📚 Additional Resources

- [Introduction to Algorithms (CLRS)](https://mitpress.mit.edu/books/introduction-algorithms-fourth-edition)
- [Visualgo - Algorithm Visualization](https://visualgo.net/en)
- [Big O Cheat Sheet](https://www.bigocheatsheet.com/)
- [GeeksforGeeks](https://www.geeksforgeeks.org/)

## ✨ Acknowledgments

- Thanks to all contributors who have helped with implementations
- Special thanks to the Go community for their excellent documentation and resources
