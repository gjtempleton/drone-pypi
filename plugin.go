package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
)

// Plugin defines the PyPi plugin parameters
type Plugin struct {
	Repository    string
	Username      string
	Password      string
	Distributions []string
}

func (p Plugin) createConfig() error {
	f, err := os.Create(path.Join(os.Getenv("HOME"), ".pypirc"))
	if err != nil {
		return err
	}
	defer f.Close()
	buff := bufio.NewWriter(f)
	_, err = io.WriteString(buff, fmt.Sprintf(`[distutils]
index-servers =
    pypi

[pypi]
repository: %s
username: %s
password: %s
`, p.Repository, p.Username, p.Password))
	if err != nil {
		return err
	}
	buff.Flush()
	return nil
}

// Exec runs the plugin - doing the necessary setup.py modifications
func (p Plugin) Exec() error {
	err := p.createConfig()

	file, err := os.Open(path.Join(os.Getenv("HOME"), ".pypirc"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(file)
	if err != nil {
		log.Fatalf("Unable to write .pypirc file due to: %s", err)
	}
	distributions := []string{"sdist"}
	if len(p.Distributions) > 0 {
		distributions = p.Distributions
	}
	args := []string{"setup.py"}
	for i := range distributions {
		args = append(args, distributions[i])
	}
	args = append(args, "upload")
	args = append(args, "-r")
	args = append(args, "pypi")
	out, err := exec.Command("python3", args...).CombinedOutput()

	if err != nil {
		log.Printf("Error enountered: %v", err)
	}
	log.Printf("Output: %s", out)
	return nil
}
