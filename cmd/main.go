package cmd

import (
	"fmt"
	"os"

	"github.com/mudler/edgevpn/internal"
	"github.com/mudler/edgevpn/pkg/edgevpn"
	"github.com/urfave/cli"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

const Copyright string = `	edgevpn  Copyright (C) 2021 Ettore Di Giacinto
This program comes with ABSOLUTELY NO WARRANTY.
This is free software, and you are welcome to redistribute it
under certain conditions.`

var CommonFlags []cli.Flag = []cli.Flag{
	&cli.StringFlag{
		Name:   "config",
		Usage:  "Specify a path to a edgevpn config file",
		EnvVar: "EDGEVPNCONFIG",
	},
	&cli.StringFlag{
		Name:   "token",
		Usage:  "Specify an edgevpn token in place of a config file",
		EnvVar: "EDGEVPNTOKEN",
	}}

func MainFlags() []cli.Flag {
	return append([]cli.Flag{
		&cli.BoolFlag{
			Name:  "g",
			Usage: "Generates a new configuration and prints it on screen",
		},
		&cli.StringFlag{
			Name:   "address",
			Usage:  "VPN virtual address",
			EnvVar: "ADDRESS",
			Value:  "10.1.0.1/24",
		},
		&cli.StringFlag{
			Name:   "interface",
			Usage:  "Interface name",
			Value:  "edgevpn0",
			EnvVar: "IFACE",
		}}, CommonFlags...)
}

func Main(l *zap.Logger) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		if c.Bool("g") {
			// Generates a new config and exit
			newData, err := edgevpn.GenerateNewConnectionData()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			bytesData, err := yaml.Marshal(newData)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println(string(bytesData))
			os.Exit(0)
		}

		e := edgevpn.New(cliToOpts(l, c)...)

		l.Sugar().Info(Copyright)

		l.Sugar().Infof("Version: %s commit: %s", internal.Version, internal.Commit)

		l.Sugar().Info("Start")

		if err := e.Start(); err != nil {
			l.Sugar().Fatal(err.Error())
		}

		return nil
	}
}
