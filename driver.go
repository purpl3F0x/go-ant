/*
 * driver.go
 *
 * Copyright (c) 2021 Stavros Avramidis (@purpl3F0x). All rights reserved.
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

package ant

import (
	"errors"
	"github.com/google/gousb"
	"log"
)

type UsbDevice struct {
	vid, pid   gousb.ID
	context    *gousb.Context
	device     *gousb.Device
	closeIface func()
	intf       *gousb.Interface
	in         *gousb.InEndpoint
	out        *gousb.OutEndpoint
}

func (dev *UsbDevice) Open() (e error) {
	log.Println("Opening USB device")

	dev.context = gousb.NewContext()

	dev.device, e = dev.context.OpenDeviceWithVIDPID(dev.vid, dev.pid)

	if e != nil {
		return
	}

	if dev.device == nil {
		e = errors.New("USB Device not found")
		return
	}

	if dev.device.SetAutoDetach(true) != nil {
		return
	}

	// Claim default interface
	dev.intf, dev.closeIface, e = dev.device.DefaultInterface()
	if e != nil {
		return
	}

	// Open an OUT endpoint.
	dev.out, e = dev.intf.OutEndpoint(1)
	if e != nil {
		return
	}

	// Open an IN endpoint.
	dev.in, e = dev.intf.InEndpoint(1)
	if e != nil {
		return
	}

	log.Println("USB Device opened")

	return
}

func (dev *UsbDevice) Close() {
	log.Println("Closing USB device")

	if dev.closeIface != nil {
		dev.closeIface()
	}

	if dev.device != nil {
		_ = dev.device.Close()
	}

	if dev.context != nil {
		_ = dev.context.Close()
	}
	log.Println("USB Device closed")
}

func (dev *UsbDevice) Read(b []byte) (int, error) {
	return dev.in.Read(b)
}

func (dev *UsbDevice) Write(b []byte) (int, error) {
	return dev.out.Write(b)
}

func (dev *UsbDevice) BufferSize() int {
	return 64
}

func GetUsbDevice(vid, pid gousb.ID) *UsbDevice {
	return &UsbDevice{
		vid: vid,
		pid: pid,
	}
}

func AntUsb1() *UsbDevice {
	return GetUsbDevice(0x0FCF, 0x1004)
}

func AntDevBoardUSB() *UsbDevice {
	return GetUsbDevice(0x0FCF, 0x1006)
}

func AntUsb2() *UsbDevice {
	return GetUsbDevice(0x0FCF, 0x1008)
}

func AntUsbM() *UsbDevice {
	return GetUsbDevice(0x0FCF, 0x1009)
}
