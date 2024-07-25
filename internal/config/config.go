package config

import "os"

func Host() string {
	host := os.Getenv("HOST_EXC")
	if host == "" {
		host = "localhost:"
	}
	return host
}

func Port() string {
	port := os.Getenv("TODO_PORT") // TODO uber logger
	if port == "" {
		port = "7540"
	}

	return port
}
