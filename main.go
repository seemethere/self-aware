package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	hostname := exec.Command("hostname")
	var hostnameOut bytes.Buffer
	hostname.Stdout = &hostnameOut
	if err := hostname.Run(); err != nil {
		panic(err)
	}

	containerInfo, err := cli.ContainerInspect(context.Background(), strings.TrimSpace(hostnameOut.String()))
	if err != nil {
		panic(err)
	}

	fmt.Printf("[*] You're running this with: %s", containerInfo.Config.Image)
}
