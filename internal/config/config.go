package config

import "os"

func Host() string {
	host := os.Getenv("HOST_EXC") // export HOST_EXC=ваш_хост_здесь
	if host == "" {
		host = "localhost:"
	}
	return host
}

func Port() string {
	port := os.Getenv("TODO_PORT") //export TODO_PORT=ваш_порт_здесь
	if port == "" {
		port = "7540"
	}

	return port
}

// TODO uber logger
