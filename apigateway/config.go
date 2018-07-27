package main

import (
	"io/ioutil"
	"log"

	"fmt"

	"github.com/jessevdk/go-flags"
	"gopkg.in/yaml.v2"
)

// settings is configure settings
type settings struct {
	// ServerAddress is address of server.
	ServerAddress string `yaml:"server" short:"s" long:"server" description:"Server address" required:"true" default:":8080"`
	//PlayerAddress is address of player service.
	PlayerAddress string `yaml:"players" short:"p" long:"players" description:"Player Service address" required:"true"default:":8081"`
	// ConfigFile is file with configs.
	ConfigFile string `short:"f" long:"configfile" description:"File with config"`
}

// Parse parses command line parameters. If there is ConfigFile, then override params by values from file.
func (s *settings) Parse() error {
	parser := flags.NewParser(s, flags.Default|flags.IgnoreUnknown)
	_, err := parser.Parse()
	if err != nil {
		return fmt.Errorf("parse error: %v", err)
	}
	if s.ConfigFile != "" {
		err = s.LoadOptionsFromFile()
		if err != nil {
			log.Printf("cannot read settings from file: %v", err)
		}
	}
	return nil
}

// LoadOptionsFromFile tries to read settings from file.
func (s *settings) LoadOptionsFromFile() error {
	data, err := ioutil.ReadFile(s.ConfigFile)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, s)
}
