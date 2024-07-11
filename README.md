# Go Tutorial.
<b>Go</b> is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. Go is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency. The language is often referred to as Golang because of its domain name, golang.org, but the proper name is Go. The language was announced in November 2009 and is used in some of Google's production systems. 

## Installation
To install Go, you can download the binary distribution from the [official website](https://golang.org/dl/). The installation process is straightforward, and you can find the installation instructions for your operating system on the same page.

## Hello World
Here is a simple program that prints "Hello, World!" to the console:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

To run this program, save it to a file named `hello.go` and run the following command in the terminal:

```bash
    go mod init
    go run hello.go
```