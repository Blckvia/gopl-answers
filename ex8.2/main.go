// ex8.2 is a minimal ftp server
// copied from https://github.com/xingdl2007/gopl-solutions/blob/master/ch8/8.2/main.go
package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

func mustCopy(w io.Writer, r io.Reader) {
	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}
}

func exeCmd(w io.Writer, e string, args ...string) {
	cmd := exec.Command(e, args...)
	cmd.Stdout = w
	if err := cmd.Run(); err != nil {
		log.Print(err)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)

	for input.Scan() {
		cmds := strings.Split(input.Text(), " ")
		switch cmds[0] {
		case "ls":
			exeCmd(c, cmds[0], cmds[1:]...)
		case "cd":
			if err := os.Chdir(cmds[1]); err != nil {
				log.Print(err)
			}
		case "get":
			file, err := os.Open(cmds[1])
			if err != nil {
				log.Printf("file %s: %v", cmds[1], err)
			}
			mustCopy(c, file)
		case "close":
			return
		default:
			help := "ls: list content\ncd: change directory\nget: get file content\n" +
				"close: close connection\n"
			mustCopy(c, strings.NewReader(help))
		}

	}
}

func main() {

}
