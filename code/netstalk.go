package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli"
)

type iparea []struct {
	IP    string `json:"ip"`
	Ports []struct {
		Port int `json:"port"`
	} `json:"ports"`
}

const (
	ColorGreen string = "\033[32m"
	ColorReset string = "\033[0m"
)

var ip iparea

var (
	file_db, _    = os.OpenFile("log/database.txt", os.O_APPEND|os.O_WRONLY, 0666)
	file_ip, _    = os.OpenFile("log/ip.txt", os.O_APPEND|os.O_WRONLY, 0666)
	file_read, _  = os.OpenFile("out.json", os.O_RDONLY, 0777)
	file_write, _ = os.OpenFile("result.txt", os.O_APPEND|os.O_WRONLY, 0666)
)

func generate_ip() string {
	var octets []string
	for i := 0; i < 4; i++ {
		octets = append(octets, strconv.Itoa(rand.Intn(255)))
	}

	result := strings.Join(octets, ".")
	return result
}

func writeFile(count int) {
	defer file_db.Close()
	defer file_ip.Close()

	data, _ := ioutil.ReadAll(file_db)

	n := 0
	for i := 1; i <= count; i++ {
		addr := generate_ip()
		file_db.WriteString(addr + "\n")

		if !strings.Contains(string(data), addr) {
			n += 1
			file_ip.WriteString(addr + "\n")
		}
	}
	fmt.Printf("%vCreated %v IP address%v\n", ColorGreen, n, ColorReset)
}

func format_json() {
	defer file_read.Close()
	defer file_write.Close()

	byteValue, _ := ioutil.ReadAll(file_read)
	json.Unmarshal(byteValue, &ip)

	for index := range ip {
		ports, _ := json.Marshal(ip[index].Ports)
		file_write.WriteString(ip[index].IP + strings.Trim(string(ports), `[{"port}]`) + "\n")
	}

	log.Printf("%vFormatted %v IP address%v\n", ColorGreen, len(ip), ColorReset)
}

func main() {
	rand.Seed(time.Now().UTC().Unix())

	app := cli.NewApp()
	app.Name = "rand-stalk"
	app.Usage = "generate random IP-address"
	app.Commands = []cli.Command{
		{
			Name:    "gen",
			Aliases: []string{"g"},
			Usage:   "gen [count INT] random IP address generation like",
			Action: func(c *cli.Context) {
				count, _ := strconv.Atoi(os.Args[2])
				writeFile(count)
			},
		},
		{
			Name:    "format",
			Aliases: []string{"f"},
			Usage:   "formatting masscan json output to text format like 0.0.0.0:80",
			Action: func(c *cli.Context) {
				format_json()
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
