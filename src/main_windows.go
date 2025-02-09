//go:build windows
// +build windows

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

type McpServer struct {
	Command string            `json:"command"`
	Args    []string          `json:"args"`
	Env     map[string]string `json:"env"`
}

type Config struct {
	McpServers map[string]McpServer `json:"mcpServers"`
}

func executeServer(server McpServer) {
	cmd := exec.Command(server.Command, server.Args...)
	
	// Set up environment variables
	cmd.Env = os.Environ()
	for k, v := range server.Env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
	}

	// Windows-specific: Hide the command prompt window
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	// Connect stdio
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(fmt.Sprintf("Failed to run command: %v", err))
	}
} 