/*
 * message.go
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
	"fmt"
)

type Message struct {
	Id      byte
	payload []byte
}

func NewMessage(id byte, payload []byte) *Message {
	return &Message{id, payload}
}

func (m Message) length() int {
	return len(m.payload) + MESG_FRAME_SIZE
}

func (m Message) String() string {
	return fmt.Sprint("AntMessage", "(", m.Id, "): ", m.payload)

}

func (m Message) Checksum() (checksum byte) {
	n := len(m.payload)
	checksum = MESG_TX_SYNC ^ byte(n) ^ m.Id
	for i := 0; i < n; i++ {
		checksum ^= m.payload[i]
	}
	return checksum
}

func (m Message) Encode() []byte {
	l := m.length()
	raw := make([]byte, l)

	raw[0] = MESG_TX_SYNC
	raw[1] = byte(l)
	raw[2] = m.Id
	copy(raw[MESG_DATA_OFFSET:], m.payload)

	checksum := MESG_TX_SYNC ^ byte(l) ^ m.Id

	for _, v := range m.payload {
		checksum ^= v
	}
	raw[l-1] = checksum

	return raw
}

func Decode(buffer []byte) (m *Message, err error) {

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
