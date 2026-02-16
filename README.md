# ✍️ SyncScribe: Real-Time Collaborative Engine

SyncScribe is a high-concurrency, distributed text editor built with **Go** and **WebSockets**. It solves the "Consistency Problem" in collaborative editing by using **Conflict-free Replicated Data Types (CRDTs)** and a **Fractional Indexing** strategy, ensuring that all users see the same document state regardless of network latency or message order.

## 🚀 Features

* **Real-Time Synchronization**: Bi-directional event streaming via WebSockets for sub-100ms perceived latency.
* **Fractional Indexing**: A mathematical approach to positioning that allows for infinite insertions between characters without index shifting bugs.
* **Conflict Resolution**: Built-in logic to handle concurrent edits using unique identifiers and deterministic sorting.
* **Modern Workspace UI**: A clean, distraction-free interface designed for SDE2-level portfolios.
* **Consistency Guarantee**: Uses the LWW (Last-Write-Wins) principle to ensure eventual consistency across all replicas.



## 🛠️ Tech Stack

| Layer | Technology | Role |
| :--- | :--- | :--- |
| **Backend** | **Go (Golang)** | High-concurrency WebSocket Hub and Message Dispatcher |
| **Frontend** | **Vanilla JS / CSS3** | Custom CRDT implementation and state-managed UI |
| **Real-time** | **WebSockets** | Low-latency bi-directional communication |
| **Data Model** | **LWW-Element-Set** | Conflict-free Replicated Data Type for eventual consistency |

---

## 🧠 The Engineering Behind SyncScribe

### 1. Fractional Indexing Logic
Instead of using standard array indices (0, 1, 2...), SyncScribe assigns a `float64` position to every character. This allows the system to calculate a new position for any insertion:
* **Append**: `Last Position + 1000.0`
* **Insert Middle**: `(Before Position + After Position) / 2`
* **Prepend**: `First Position - 1000.0`

This ensures that even if User A and User B type at the same time, their characters are inserted into a deterministic order based on their unique fractional value.



### 2. The WebSocket Hub (Concurrency)
The Go backend utilizes a **Hub-and-Spoke** model. It manages a thread-safe map of clients and uses Go Channels to broadcast operations. This ensures that the server remains non-blocking even as the number of concurrent collaborators scales.

### 3. Cursor Synchronization & UI Polish
To prevent the "backward typing" bug, SyncScribe implements a **Local Cursor Tracker**. It manually offsets the selection range after every remote update, ensuring a natural typing experience where the cursor always follows the text.

---

## 📂 Project Structure

```text
SyncScribe/
├── cmd/
│   └── server/
│       └── main.go          # Entry point for the Go server
├── internal/
│   ├── crdt/
│   │   └── lww_set.go       # CRDT Data Structures (ID, Position, Value)
│   └── socket/
│       ├── client.go        # WebSocket Client logic
│       └── hub.go           # Broadcast & Connection management
├── static/
│   └── index.html           # Modern UI & Frontend CRDT logic
└── go.mod                   # Project dependencies
```

---

## 🚥 Getting Started

Follow these steps to set up the project locally.

### 1. Clone and Initialize
Clone the repository to your local machine and navigate into the project directory:

```bash
git clone https://github.com/rohannnguptaaa/SyncScribe.git
cd SyncScribe
```

### 2. Initialize the module and download dependencies

```bash
go mod init
go mod tidy
```

### 3. Launch the Go backend hub

```bash
go run cmd/server/main.go
```

### 4. Open multiple tabs on localhost:8080 to test the editor