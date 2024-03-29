/*
 * main.go
 *
 * Copyright (c) 2021-2023 Stavros Avramidis (@purpl3F0x). All rights reserved.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 *
 *
 */

package main

import (
	"github.com/purpl3F0x/go-ant"
	"time"
)

func eval(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	dev := ant.GetUsbDevice(0x0FCF, 0x1008)

	readChan := make(chan *ant.Message)
	Ant := ant.MakeAnt(dev, readChan)

	eval(Ant.Start())

	// defer Ant.ResetSystem()
	defer Ant.Stop()

	Ant.ResetSystem()

	time.Sleep(1000 * time.Millisecond)

	Ant.SetNetworkKey(0, ant.AntPlusNetworkKey())
	Ant.AssignChannel(0, ant.PARAMETER_RX_NOT_TX, 0)
	Ant.SetChannelId(0, 0, 120, 0)
	Ant.SetChannelPeriod(0, 8070)
	Ant.SetChannelRFFreq(0, 57)
	Ant.OpenRxScanMode()

	select {}
	// time.Sleep(100 * time.Second)

	// interrupt := make(chan os.Signal, 1)
	// signal.Notify(interrupt, os.Interrupt)
	// <-interrupt
}
