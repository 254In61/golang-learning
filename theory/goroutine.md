# overview
- A goroutine is a lightweight thread of execution in the Go programming language. 
- It is a key feature of Go that enables concurrency, allowing a program to perform multiple tasks simultaneously.

# Characteristics
1. Lightweight : 
   - Goroutines are much lighter than traditional operating system threads.
   - They start with a small amount of stack space (around 2 KB), which grows and shrinks as needed.

2. Managed by Go Runtime:
   - The Go runtime manages goroutines, including scheduling and execution, using a mechanism called the Go scheduler.
   - Unlike OS threads, goroutines are multiplexed onto a smaller number of OS threads, making them more efficient.

3. Concurrency:
   - Goroutines allow for concurrent execution of functions, which means multiple tasks can run independently at the same time.

4. Ease of Use:
   - You can create a goroutine by simply prefixing a function call with the go keyword.

# Goroutines and Channels
- Goroutines often use channels to communicate and synchronize with each other. 
- Channels are a way to safely pass data between goroutines.

