package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

// Plugin defines the PyPi plugin parameters
type Plugin struct {
	Repository    string
	Username      string
	Password      string
	Distributions []string
}

func (p Plugin) createConfig(w io.Writer) error {

	_, err := io.WriteString(w, fmt.Sprintf(`[distutils]
		index-servers =
			pypi
		[pypi]
		repository: %s
		username: %s
		password: %s
		`, p.Repository, p.Username, p.Password))
	return err
}

// Exec runs the plugin - doing the necessary setup.py modifications
func (p Plugin) Exec() error {
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
