# Pipeline Pattern

## What is this pattern?
The **Pipeline Pattern** connects multiple stages of computation where the output of one stage is the input of the next.  
Each stage runs in its own goroutine and communicates through channels, enabling concurrent execution and a clean, linear data flow.

---

## Key Benefits of Pipeline Pattern
- **Concurrent Execution** → each stage runs independently in its own goroutine.  
- **Modularity** → stages can be reused and composed into larger workflows.  
- **Streaming Data Processing** → values are processed as they arrive, not stored all at once in memory.  
- **Backpressure Handling** → slow consumers naturally slow down producers via channel blocking, ensuring system stability.  
- **Scalability** → stages can be parallelized (fan-out/fan-in) for higher throughput.  


