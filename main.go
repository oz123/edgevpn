// Copyright © 2021 Ettore Di Giacinto <mudler@mocaccino.org>
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 2 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program; if not, see <http://www.gnu.org/licenses/>.

package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/mudler/edgevpn/cmd"
	internal "github.com/mudler/edgevpn/internal"
	"go.uber.org/zap"
)

func main() {
	l, _ := zap.NewProduction()
	defer l.Sync() // flushes buffer, if any

	app := &cli.App{
		Name:        "edgevpn",
		Version:     internal.Version,
		Author:      "Ettore Di Giacinto",
		Usage:       "edgevpn --config /etc/edgevpn/config.yaml",
		Description: "edgevpn uses libp2p to build an immutable trusted blockchain addressable p2p network",
		Copyright:   cmd.Copyright,
		Flags:       cmd.MainFlags(),
		Commands: []cli.Command{
			cmd.Join(l),
			cmd.API(l),
			cmd.ServiceAdd(l),
			cmd.ServiceConnect(l),
			cmd.FileReceive(l),
			cmd.FileSend(l),
		},

		Action: cmd.Main(l),
	}

	err := app.Run(os.Args)
	if err != nil {
		l.Sugar().Fatal(err)
	}
}
