#!/bin/bash

startServer() {
    go run ./server/server.go &
    SERVER_PID=$!
    echo "Server application started with PID: $SERVER_PID"
}

startClient() {
    go run ./client/client.go
}

killProcesses() {
    pkill -x server
    pkill -x client
}

checkServerRunning() {
    if pgrep -x "server" >/dev/null; then
        echo "Server application is already running."
        return 0
    else
        return 1
    fi
}

killProcesses

if checkServerRunning; then
    startClient
else
    startServer
    echo "Waiting for the server application to start..."
    while ! checkServerRunning; do
        sleep 1
    done
    echo "Server application is now running."
    startClient
fi
