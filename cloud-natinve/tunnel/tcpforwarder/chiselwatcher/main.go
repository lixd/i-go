package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"

	"gopkg.in/yaml.v2"
)

type Tunnel struct {
	Local  int `yaml:"local"`
	Remote int `yaml:"remote"`
}

type Conf struct {
	Hosts map[string]map[string]Tunnel `yaml:"hosts"`
}

func main() {

	// bytes, err := ioutil.ReadFile(os.Args[1])
	bytes, err := ioutil.ReadFile("./config.yaml")

	if err != nil {
		log.Fatal(err)
	}

	var conf Conf
	err = yaml.Unmarshal(bytes, &conf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%-14s %-26s %-10s %-10s\n", "host", "service", "local", "remote")

	for host, tunnels := range conf.Hosts {
		var args []string
		for name, tunnel := range tunnels {
			args = append(args, "-L", fmt.Sprintf("%d:localhost:%d", tunnel.Local, tunnel.Remote))
			fmt.Printf("%-14s %-26s %-10d %-10d\n", host, name, tunnel.Local, tunnel.Remote)
		}
		args = append(args, "-N", host)
		go func() {
			fmt.Printf("ssh %v", strings.Join(args, " "))
			// cmd := exec.Command("ssh", args...)
			// err = cmd.Start()
			// if err != nil {
			// 	log.Fatal(err)
			// }
		}()
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	fmt.Println("exiting")
}
