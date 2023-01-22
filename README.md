# go-ant

Garmin's ANT, ANT+, ANT-FS protocol stack and ANT-USB driver written in Go

## Installation

1. Install with:
    ```go 
    go get -u github.com/purpl3F0x/go-ant
    ```
2. Import to project
    ```go
    import "github.com/purpl3F0x/go-ant"
    ```

## Example 

```go
	dev := ant.GetUsbDevice(0x0FCF, 0x1008)
	Ant := ant.MakeAnt(dev)
	Ant.Start()
	Ant.ResetSystem()
	Ant.SetNetworkKey(0, ant.AntPlusNetworkKey())
	Ant.AssignChannel(0, ant.PARAMETER_RX_NOT_TX, 0)
	Ant.SetChannelId(0, 0, 0, 0)
	Ant.SetChannelRFFreq(0, 57)
	Ant.SetChannelPeriod(0, 8192)
	Ant.OpenRxScanMode()
```



## License
```
Copyright (c) 2021-3 Stavros Avramidis (@purpl3F0x). All rights reserved.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
```