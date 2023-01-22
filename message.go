/*
 * message.go
 *
 * Copyright (c) 2021-2022 Stavros Avramidis (@purpl3F0x). All rights reserved.
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
	"fmt"
	"strings"
)

type Packet []byte

func (p Packet) String() string {
	l := len(p)

	if l == 0 {
		return "[]"
	}
	l--
	var b strings.Builder
	b.Grow((l)*6 + 2 + 4)
	b.Write(Packet(fmt.Sprint("[")))
	for i := 0; i < l; i++ {
		b.Write(Packet(fmt.Sprintf("0x%02X, ", p[i])))
	}
	b.Write(Packet(fmt.Sprintf("0x%02X]", p[l])))
	return b.String()
}

type Message struct {
	Id   byte
	Data Packet
}

func NewMessage(id byte, data Packet) *Message {
	return &Message{id, data}
}

func (m Message) length() int {
	return len(m.Data) + MESG_FRAME_SIZE
}

func (m Message) String() string {
	return fmt.Sprintf("AntMessage (0x%02X) %s", m.Id, m.Data)

}

func (m Message) Checksum() (checksum byte) {
	n := len(m.Data)
	checksum = MESG_TX_SYNC ^ byte(n) ^ m.Id
	for i := 0; i < n; i++ {
		checksum ^= m.Data[i]
	}
	return checksum
}

func (m Message) Encode() Packet {
	rawLen := m.length()
	msgLen := len(m.Data)
	raw := make(Packet, rawLen)

	raw[0] = MESG_TX_SYNC
	raw[1] = byte(msgLen)
	raw[2] = m.Id
	copy(raw[MESG_DATA_OFFSET:], m.Data)

	checksum := MESG_TX_SYNC ^ byte(msgLen) ^ m.Id

	for n := 2; n < msgLen; n++ {
		checksum ^= m.Data[n]
	}
	raw[rawLen-1] = checksum
	// raw[rawLen] = 0
	// raw[rawLen+1] = 0

	return raw
}

func Decode(buffer Packet) (m *Message, err error) {

	sync := buffer[0]
	length := int(buffer[MESG_SIZE_OFFSET])
	id := buffer[MESG_ID_OFFSET]
	data := buffer[MESG_DATA_OFFSET : len(buffer)-1]
	checksum := buffer[len(buffer)-1]

	if sync != MESG_TX_SYNC {
		return nil, errors.New(fmt.Sprintf("Could not decode. Expected TX sync but got 0x%.2x.", sync))
	}

	if len(buffer) != length+MESG_FRAME_SIZE {
		return nil, errors.New(fmt.Sprintf("Could not decode. Message length should be %d but was %d.", length+MESG_FRAME_SIZE, len(data)))
	}

	m = NewMessage(id, data)

	if checksum != m.Checksum() {
		return nil, errors.New(fmt.Sprintf("Could not decode. Checksum should be 0x%.2x but was 0x%.2x.", checksum, m.Checksum()))
	}

	return m, nil
}
