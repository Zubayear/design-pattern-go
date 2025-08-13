# Go Design Patterns

This repository documents the implementation of common **Software Design Patterns** in the Go programming language. Each pattern includes:

- 🛠️ Go implementation
- 📘 Use cases

---

## 📐 Pattern Categories

### 🔨 Creational Patterns
These patterns provide object creation mechanisms that increase flexibility and reuse of existing code.

| Pattern    | Description |
|------------|-------------|
| Singleton  | Ensures a class has only one instance and provides a global point of access to it. |
| Factory    | Provides an interface for creating objects but lets subclasses alter the type. |
| Builder    | Constructs complex objects step by step. |
| Prototype  | Clones existing objects without depending on their classes. |

---

### 🏗️ Structural Patterns
These patterns explain how to assemble objects and classes into larger structures.

| Pattern     | Description |
|-------------|-------------|
| Adapter     | Converts the interface of a class into another interface clients expect. |
| Decorator   | Adds behavior to an object dynamically without affecting others. |
| Proxy       | Provides a surrogate or placeholder for another object. |
| Facade      | Provides a unified interface to a set of interfaces in a subsystem. |

---

### 🔁 Behavioral Patterns
These patterns are concerned with communication between objects.

| Pattern     | Description |
|-------------|-------------|
| Strategy    | Defines a family of algorithms and makes them interchangeable. |
| Observer    | Lets objects notify other objects about changes in state. |
| Command     | Encapsulates a request as an object. |
| Chain of Responsibility | Passes a request along a chain of handlers. |

---

## 🧪 Running the Examples

```bash
go test ./...
```

---
📚 Resources Used
* Head First Design Patterns
* Design Patterns: Elements of Reusable Object-Oriented Software (Gang of Four)
* Official Go Documentation

---
🏗️ Contribution
Feel free to fork the repo and add more examples or improvements. Contributions are welcome!
```bash
git clone https://github.com/your-username/go-design-patterns.git
cd go-design-patterns
```
