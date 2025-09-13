# Fan-Out / Fan-In Pattern

## What is this pattern?
The **Fan-Out / Fan-In Pattern** is a concurrency design pattern used to process data in parallel and then aggregate the results.  

- **Fan-Out**: Distributes tasks across multiple worker goroutines to be processed concurrently.  
- **Fan-In**: Merges results from multiple workers into a single channel for further processing.  

This pattern allows Go programs to efficiently utilize multiple cores while maintaining controlled concurrency.

---

## Key Benefits of Fan-Out / Fan-In Pattern
- **Parallel Processing** → multiple workers process jobs simultaneously.  
- **Resource Efficiency** → limits the number of active goroutines while handling multiple tasks.  
- **Synchronization** → results are aggregated in a controlled way using channels.  
- **Scalability** → easy to scale up by increasing the number of workers.  