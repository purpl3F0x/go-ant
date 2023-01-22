/*
 * ant.go
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

package ant

import (
	"encoding/binary"
	"fmt"
	"log"
	"time"
)

type Driver interface {
	Open() error
	Close()
	Read(b []byte) (int, error)
	Write(b []byte) (int, error)
	BufferSize() int
}

type Ant struct {
	driver          Driver
	buffer          Packet
	read            chan *Message
	write           chan *Message
	writeInTimeslot chan *Message
	stopper         chan struct{}
	decoder         chan byte
	done            chan struct{}
}

func MakeAnt(dev Driver, read chan *Message) (ant *Ant) {
	ant = &Ant{
		driver: dev,
		read:   read,
		write:  make(chan *Message),
		// writeInTimeslot: make(chan *Message),
		stopper: make(chan struct{}),
		decoder: make(chan byte),
		done:    make(chan struct{}),
	}

	return ant
}

func (dev *Ant) Start() (e error) {
	log.Println("Starting Device")
	e = dev.driver.Open()

	if e != nil {
		return e
	}

	dev.buffer = make(Packet, dev.driver.BufferSize())

	go dev.loop()
	go dev.decodeLoop()
	go dev.readLoop()
	return e
}

func (dev *Ant) Stop() {
	dev.stopper <- struct{}{}
	dev.buffer = nil

	// Wait for loops to finish
	<-dev.done
	<-dev.done
}

func (dev *Ant) loop() {

	// ticker := time.NewTicker(time.Millisecond)
	defer func() { dev.done <- struct{}{} }()
	defer dev.driver.Close()
	defer close(dev.decoder)
	defer close(dev.write)
	// defer ticker.Stop()
	// defer close(dev.writeInTimeslot)
	defer log.Println("Loop stopped!")

	log.Println("Loop Started")

	for {
		select {
		case <-dev.stopper:
			return

		case d := <-dev.write:
			m := d.Encode()

			log.Println("Writing: ", m)
			_, err := dev.driver.Write(m)
			if err != nil {
				log.Println(err)
			}
			time.Sleep(time.Nanosecond)
		}
	}
}

func (dev *Ant) readLoop() {
	ticker := time.NewTicker(time.Millisecond)

	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if i, err := dev.driver.Read(dev.buffer); err == nil {
				for _, v := range dev.buffer[:i] {
					dev.decoder <- v

				}
			}
		}
	}

}

func (dev *Ant) decodeLoop() {
	defer func() { dev.done <- struct{}{} }()
	defer close(dev.read)

	for {
		// Wait for TX Sync
		if sync, ok := <-dev.decoder; !ok {
			return
		} else if sync != MESG_TX_SYNC {
			continue
		}

		// Get content length (+1byte type + 1byte checksum)
		length, ok := <-dev.decoder
		if !ok {
			continue
		}

		buf := make(Packet, length+4)
		buf[0] = MESG_TX_SYNC
		buf[1] = length
		for i := 2; i < int(length+4); i++ {
			if buf[i], ok = <-dev.decoder; !ok {
				return
			}
		}

		// Check message integrity
		msg, err := Decode(buf)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Println("Read:", msg)
		// if dev.read != nil {
		// 	dev.read <- msg
		//
		// }

		// select {
		// case dev.read <- msg:
		// default:
		// }

	}
}

// //////////////////////////////////////////////////////////////////////////////////////
// Config Messages
// //////////////////////////////////////////////////////////////////////////////////////

func (dev *Ant) UnAssignChannel(channel uint8) {
	message := NewMessage(MESG_UNASSIGN_CHANNEL_ID, Packet{channel})
	dev.write <- message
}

func (dev *Ant) AssignChannel(channel uint8, channelType uint8, networkNumber uint8) {
	message := NewMessage(MESG_ASSIGN_CHANNEL_ID, Packet{channel, channelType, networkNumber})
	dev.write <- message
}

func (dev *Ant) AssignChannelExt(channel uint8, channelType uint8, networkNumber uint8, ExtFlags uint8) {
	message := NewMessage(MESG_ASSIGN_CHANNEL_ID, Packet{channel, channelType, networkNumber, ExtFlags})
	dev.write <- message
}

func (dev *Ant) SetChannelId(channel uint8, deviceNum uint16, deviceType uint8, transmissionType uint8) {
	payload := [5]byte{channel, 0, 0, deviceType, transmissionType}
	binary.LittleEndian.PutUint16(payload[1:], uint16(deviceNum))

	message := NewMessage(MESG_CHANNEL_ID_ID, payload[:])
	dev.write <- message
}

func (dev *Ant) SetChannelPeriod(channel uint8, messagePeriod uint16) {
	payload := [3]byte{channel, 0, 0}
	binary.LittleEndian.PutUint16(payload[1:], uint16(messagePeriod))

	message := NewMessage(MESG_CHANNEL_MESG_PERIOD_ID, payload[:])
	dev.write <- message
}

func (dev *Ant) SetChannelSearchTimeout(channel uint8, messagePeriod uint8) {
	message := NewMessage(MESG_CHANNEL_SEARCH_TIMEOUT_ID, Packet{channel, messagePeriod})
	dev.write <- message
}

func (dev *Ant) SetChannelRFFreq(channel uint8, rfFreq uint8) {
	message := NewMessage(MESG_CHANNEL_RADIO_FREQ_ID, Packet{channel, rfFreq})
	dev.write <- message
}

func (dev *Ant) SetNetworkKey(channel uint8, key [8]uint8) {
	payload := [9]byte{channel}
	copy(payload[1:], key[:])
	message := NewMessage(MESG_NETWORK_KEY_ID, payload[:])
	dev.write <- message
}

func (dev *Ant) SetTransmitPower(power uint8) {
	message := NewMessage(MESG_CHANNEL_RADIO_TX_POWER_ID, Packet{0, power & RADIO_TX_POWER_LVL_MASK})
	dev.write <- message
}

func (dev *Ant) SetSearchWaveform(channel uint8, searchWaveform uint16) {
	if searchWaveform != 316 && searchWaveform != 97 {
		panic("The search waveform to be set. One of these two values only. (316 or 97)")
	}
	payload := [3]byte{channel}
	binary.LittleEndian.PutUint16(payload[1:], uint16(searchWaveform))
	message := NewMessage(MESG_RADIO_TX_POWER_ID, payload[:])
	dev.write <- message
}

// //////////////////////////////////////////////////////////////////////////////////////
// ANT Control messages
// //////////////////////////////////////////////////////////////////////////////////////

func (dev *Ant) ResetSystem() {
	message := NewMessage(MESG_SYSTEM_RESET_ID, Packet{0})
	dev.write <- message
}

func (dev *Ant) OpenChannel(channel uint8) {
	message := NewMessage(MESG_OPEN_CHANNEL_ID, Packet{channel})
	dev.write <- message
}

func (dev *Ant) CloseChannel(channel uint8) {
	message := NewMessage(MESG_CLOSE_CHANNEL_ID, Packet{channel})
	dev.write <- message
}

func (dev *Ant) RequestMessage(channel uint8, messageId uint8) {
	message := NewMessage(MESG_REQUEST_SIZE, Packet{channel, messageId})
	dev.write <- message
}

func (dev *Ant) WriteMessage(messageID uint8, data Packet) {
	message := NewMessage(messageID, data)
	dev.write <- message
}

// //////////////////////////////////////////////////////////////////////////////////////
// The following are the synchronous RF event functions used to update the synchronous data sent over a channel
// //////////////////////////////////////////////////////////////////////////////////////

func (dev *Ant) SendBroadcastData(channel uint8, data Packet) {
	if len(data) != 8 {
		panic(fmt.Sprint("Data length should be 8 not ", len(data)))
	}

	payload := [9]byte{channel}
	copy(payload[1:], data)
	message := NewMessage(MESG_BROADCAST_DATA_ID, payload[:])

	dev.write <- message
}

func (dev *Ant) SendAcknowledgedData(channel uint8, data Packet) {
	if len(data) != 8 {
		panic(fmt.Sprint("Data length should be 8 not ", len(data)))
	}
	payload := [9]byte{channel}
	copy(payload[1:], data)
	message := NewMessage(MESG_ACKNOWLEDGED_DATA_ID, payload[:])
	dev.writeInTimeslot <- message
}

func (dev *Ant) SendBurstTransferPacket(channelSeq uint8, data Packet) {
	if len(data) != 8 {
		panic(fmt.Sprint("Data length should be 8 not ", len(data)))
	}

	payload := [9]byte{channelSeq}
	copy(payload[1:], data)
	message := NewMessage(MESG_BURST_DATA_ID, payload[:])
	dev.writeInTimeslot <- message
}

func (dev *Ant) SendBurstTransfer(channel uint8, data Packet) {
	if len(data)%8 != 0 {
		panic("Data length should be multiple of 8 not ")
	}

	packets := uint8(len(data) / 8)

	for i := uint8(0); i < packets; i++ {
		sequence := ((i - 1) % 3) + 1

		if i == 0 {
			sequence = 0
		} else if i == packets-1 {
			sequence |= 0b100
		}

		channelSeq := channel | sequence<<5

		dev.SendBurstTransferPacket(channelSeq, data[i*8:i*8+8])
	}

}

// //////////////////////////////////////////////////////////////////////////////////////
// The following functions are used with version 2 modules
// //////////////////////////////////////////////////////////////////////////////////////

func (dev *Ant) AddChannelID(channel uint8, deviceNum uint16, deviceType uint8, transmissionType uint8, index uint8) {
	payload := [6]byte{channel, 0, 0, deviceType, transmissionType, index}
	binary.LittleEndian.PutUint16(payload[1:], uint16(deviceNum))
	message := NewMessage(MESG_ID_LIST_ADD_ID, payload[:])
	dev.write <- message
}

func (dev *Ant) ConfigList(channel uint8, listSize uint8, exclude uint8) {
	message := NewMessage(MESG_ID_LIST_ADD_ID, Packet{channel, listSize, exclude})
	dev.write <- message
}

func (dev *Ant) OpenRxScanMode() {
	message := NewMessage(MESG_OPEN_RX_SCAN_ID, Packet{0, 1}) // [0-Channel, 1-Enable]
	dev.write <- message
}

// //////////////////////////////////////////////////////////////////////////////////////
// The following functions are used with AP2 modules (not AP1 or AT3)
// //////////////////////////////////////////////////////////////////////////////////////
