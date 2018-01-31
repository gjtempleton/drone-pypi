package main

// Plugin defines the PyPi plugin parameters
type Plugin struct {
	Repository    string
	Username      string
	Password      string
	Distributions []string
}

// Exec runs the plugin - doing the necessary setup.py modifications
func (p Plugin) Exec() error {
	// plugin logic goes here
	return nil
}
