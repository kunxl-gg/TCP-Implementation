# WebSocket Go Project

This Go project consists of a WebSocket server and a WebSocket client.

## Project Structure

- `websocket-server`: Contains the WebSocket server implementation.
- `websocket-client`: Contains the WebSocket client implementation.

## Prerequisites

- Go installed on your machine.

## Getting Started

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/your-repo.git
    cd your-repo
    ```

2. Build and run both the WebSocket server and client using `make`:

    ```bash
    make
    ```

3. Open your browser and navigate to [http://localhost:8080](http://localhost:8080) (or relevant URL) to interact with the WebSocket server.

## Makefile Targets

- `make`: Build and run both WebSocket server and client.
- `make websocket-server`: Build and run only the WebSocket server.
- `make websocket-client`: Build and run only the WebSocket client.
- `make clean`: Clean up artifacts and binaries.

## Contribut

