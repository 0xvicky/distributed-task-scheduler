# Distributed Task Scheduler (Go)

This project is a **learning-focused distributed systems project** where I am building a simplified **Distributed Task Scheduler** in Go.

The goal is to understand how workflow engines and task orchestration systems work internally by implementing a minimal version from scratch.

---

## Project Idea

The system will allow users to submit workflows composed of multiple tasks.

Each workflow will be stored as a **Directed Acyclic Graph (DAG)** where tasks depend on other tasks.

The scheduler will then:

- analyze task dependencies  
- schedule tasks in the correct order  
- distribute tasks to workers  
- monitor execution  
- handle failures and retries  

---

## Example Workflow

Simple workflow:

A → B → C

More complex workflow:

```
      A
     / \
    B   C
     \ /
      D
```

Tasks will only execute once their dependencies are completed.

---

## Planned System Components

- Workflow submission API  
- DAG workflow representation  
- Task scheduler  
- Worker nodes  
- Task queue  
- Task state tracking  
- Worker heartbeat monitoring  
- Basic monitoring dashboard  

---

## Learning Goals

This project is mainly to explore and learn:

- Go concurrency (goroutines, channels)  
- Distributed task scheduling  
- DAG execution and dependency tracking  
- Worker coordination  
- Failure detection and retry mechanisms  
- Task lifecycle management  
- Designing reliable backend systems  

---

## Status

Project starting — currently in the **design phase**.

The initial focus is to build a **minimal working scheduler (MVP)** and then gradually add more distributed system features.

---

## Author

Vivek Tyagi