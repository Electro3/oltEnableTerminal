package oltEnableTerminal

import (
	"github.com/google/goexpect"
	"log"
	"fmt"
	"regexp"
	"time"
)

type connectionModel struct {
	ip string
	adminPassword  string
	enablePassword string
	timeout time.Duration
}


func Create (connection connectionModel) (*expect.GExpect, error) {
	
	
	
	//prepare connection command 
	addr := fmt.Sprintf("%s%s", "-oKexAlgorithms=+diffie-hellman-group1-sha1 admin@", connection.ip)

	
	//create connection
	terminal, _, err := expect.Spawn(fmt.Sprintf("ssh %s", addr), -1)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	
	//send admin password
	terminal.Expect(regexp.MustCompile("password:"), connection.timeout)
	terminal.Send(connection.adminPassword + "\n")
	

	// verify admin login
	out2, ma, err := terminal.Expect(regexp.MustCompile(">"), connection.timeout)
	if err != nil {
		log.Println("Fail admin Login ", ma, err, out2)
		return nil, err
	}

	// send enable command
	terminal.Send("enable\n")
	
	
	//send enable password
	terminal.Expect(regexp.MustCompile("Password:"), connection.timeout)
	terminal.Send(connection.enablePassword + "\n")

	// verify enable login
	out2, ma, err := terminal.Expect(regexp.MustCompile("#"), connection.timeout)
	if err != nil {
		log.Println("Fail enable Login ", ma, err, out2)
		return nil, err
	}
	
	return terminal, nil 

}
