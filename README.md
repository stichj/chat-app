# 💬 TCP Chat App in Go
Welcome to my first Go project! I have implemented a real-time chat app with a terminal client interface on top of TCP sockets. Feel free to try it out yourselves 😁. 

## Features 
- Real-time messaging over TCP
- Supports multiple clients concurrently
- Username prompt for personalization
- Graceful client disconnection via `/quit`, `/q`, or `Ctrl+C`
- Pub/Sub-pattern to manage client registration, broadcasting, and deregistration
- Dockerized server and client containers for local development

## 🚀 Getting Started 

1. Clone the repository 
```bash
git clone https://github.com/your-username/tcp-chat-go.git
cd tcp-chat-go
```
2. Build and run with Docker
    
    Make sure Docker is installed and running

    ```bash
    docker-compose up --build
    ````
    If you are running v2, use ```docker compose up --build```.

    Docker will start both, a server and a client. If you want to run another client, just open another Terminal window and run
    ```bash
    docker-compose run chat-client
    ```
    
## 💻 Usage

After launching:
- You'll be prompted to enter your username.
- Start chatting with others!
- Type `/quit` or `/q` or press `Ctrl+C` to leave the chat gracefully.

## 📚 Future Improvements

- Add authentication or user presence
- Persist chat history
- WebSocket or HTTP front-end

## 📄 License

MIT License – see LICENSE file for details.