package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func check(err error, message string) {
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", message)
}

type ClientJob struct {
	name string
	conn net.Conn
}

func generateResponses(clientJobs chan ClientJob) {
	for {
		clientJob := <-clientJobs

		for start := time.Now(); time.Now().Sub(start) < time.Second; {

		}

		clientJob.conn.Write([]byte("Hello, " + clientJob.name))
	}
}

func main() {
	clientJobs := make(chan ClientJob)
	go generateResponses(clientJobs)

	ln, err := net.Listen("tcp", ":8080")
	check(err, "Server is ready")

	for {
		conn, err := ln.Accept()
		check(err, "Accepted connection.")

		go func() {
			buf := bufio.NewReader(conn)

			for {
				name, err := buf.ReadString('\n')

				if err != nil {
					fmt.Printf("Client disconnected.")
					break
				}

				clientJobs <- ClientJob{name, conn}
			}
		}()
	}
}