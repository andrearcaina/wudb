# Key-Value Store

A key value store is basically like a hashmap or a dictionary.

There are two main types of key-value stores:
1. In-Memory Key-Value Store
2. Persistent Key-Value Store

This can be done in many different ways, such as distributed or non-distributed.

## In-Memory Key-Value Store
An in-memory key-value store keeps all data in the system's main memory (RAM). This allows for very fast read and write operations, making it suitable for applications that require low latency and high throughput. However, since the data is stored in volatile memory, it is lost when the system is powered off or crashes.

## Persistent Key-Value Store
A persistent key-value store saves data to non-volatile storage, such as a hard drive or SSD. This ensures that data is retained even after a system reboot or failure. Persistent key-value stores are typically slower than in-memory stores due to the overhead of disk I/O operations, but they provide durability and reliability for long-term data storage.

## Distributed Key-Value Store
A distributed key-value store spreads data across multiple nodes or servers in a network. This makes it highly scalable as data is distributed across many machines (known as horizontal scaling). A similar concept is microservices architecture, where different services handle different parts of the application.

## Non-Distributed Key-Value Store
A non-distributed key-value store operates on a single machine. This is simpler to set up and manage but is limited by resources and often harder to scale. A similar concept is a monolithic architecture, where all components of the application are tightly integrated and run on a single server.

## The Goal
The goal of creating a key-value store is to understand how distributed systems work in practice in relation to databases. I also want to learn how to implement a key-value store from scratch, including handling data storage, retrieval, and management. Thus, I'll be planning to make it fully distributed across multiple nodes (which is like multiple computers working together) and a combination of both in-memory and persistent storage to balance speed and durability (like Redis).