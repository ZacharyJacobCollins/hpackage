I was a pre-med major 3 weeks before I began my first semester of college.  Last minute I decided that computers were pretty sweet, so I changed my major to computer science.

Never having programmed before, I grappled with my first programming course. Coming from a small highschool I had a tough time with the pacing, and I began to question whether or not I belonged in Computer Science.

I conveyed some of my concerns to a friend I had made in the class, and he sat me down, created a project for me to do, and walked me through it.  The project taught me how to use objects, methods, classes in Java.  Many of the fundamentals I was having difficulty wrapping my head around.

Because of my friend's time I stuck with computer science and aced the course.  I even ended up teaching a Java lab the following year.

Through his passion for technology, I found programming.  Without his willingness to help there is a possibility I would have changed majors and missed out on one of the things that has become dear to me.

Since then I have kept my eyes open for those who may need mentorship, and who would benefit from excitement



It was a pokemon-like game written in java (I actually still have it, https://goo.gl/g1fIQN).

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

type User struct {
	Name   string
	Output chan Message
}

type Message struct {
	Username string
	Text     string
}

type ChatServer struct {
	Users map[string]User
	Join  chan User
	Leave chan User
	Input chan Message
}

func (cs *ChatServer) Run() {
	for {
		select {
		case user := <-cs.Join:
			cs.Users[user.Name] = user
			go func() {
				cs.Input <- Message{
					Username: "SYSTEM",
					Text:     fmt.Sprintf("%s joined", user.Name),
				}
			}()
		case user := <-cs.Leave:
			delete(cs.Users, user.Name)
			go func() {
				cs.Input <- Message{
					Username: "SYSTEM",
					Text:     fmt.Sprintf("%s left", user.Name),
				}
			}()
		case msg := <-cs.Input:
			for _, user := range cs.Users {
				select {
				case user.Output <- msg:
				default:
				}
			}
		}
	}
}

func handleConn(chatServer *ChatServer, conn net.Conn) {
	defer conn.Close()

	io.WriteString(conn, "Enter your Username:")

	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	user := User{
		Name:   scanner.Text(),
		Output: make(chan Message, 10),
	}
	chatServer.Join <- user
	defer func() {
		chatServer.Leave <- user
	}()

	// Read from conn
	go func() {
		for scanner.Scan() {
			ln := scanner.Text()
			chatServer.Input <- Message{user.Name, ln}
		}
	}()

	// write to conn
	for msg := range user.Output {
		if msg.Username != user.Name {
			_, err := io.WriteString(conn, msg.Username+": "+msg.Text+"\n")
			if err != nil {
				break
			}
		}
	}
}

func main() {
	server, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer server.Close()

	chatServer := &ChatServer{
		Users: make(map[string]User),
		Join:  make(chan User),
		Leave: make(chan User),
		Input: make(chan Message),
	}
	go chatServer.Run()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handleConn(chatServer, conn)
	}
}
