//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type CipRRData struct {
	Exchange *CipExchange
	Parent   *EipPacket
}

// The corresponding interface
type ICipRRData interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *CipRRData) Command() uint16 {
	return 0x006F
}

func (m *CipRRData) InitializeParent(parent *EipPacket, sessionHandle uint32, status uint32, senderContext []uint8, options uint32) {
	m.Parent.SessionHandle = sessionHandle
	m.Parent.Status = status
	m.Parent.SenderContext = senderContext
	m.Parent.Options = options
}

func NewCipRRData(exchange *CipExchange, sessionHandle uint32, status uint32, senderContext []uint8, options uint32) *EipPacket {
	child := &CipRRData{
		Exchange: exchange,
		Parent:   NewEipPacket(sessionHandle, status, senderContext, options),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastCipRRData(structType interface{}) *CipRRData {
	castFunc := func(typ interface{}) *CipRRData {
		if casted, ok := typ.(CipRRData); ok {
			return &casted
		}
		if casted, ok := typ.(*CipRRData); ok {
			return casted
		}
		if casted, ok := typ.(EipPacket); ok {
			return CastCipRRData(casted.Child)
		}
		if casted, ok := typ.(*EipPacket); ok {
			return CastCipRRData(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *CipRRData) GetTypeName() string {
	return "CipRRData"
}

func (m *CipRRData) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *CipRRData) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 32

	// Reserved Field (reserved)
	lengthInBits += 16

	// Simple field (exchange)
	lengthInBits += m.Exchange.LengthInBits()

	return lengthInBits
}

func (m *CipRRData) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func CipRRDataParse(io utils.ReadBuffer, len uint16) (*EipPacket, error) {
	if pullErr := io.PullContext("CipRRData"); pullErr != nil {
		return nil, pullErr
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := io.ReadUint32("reserved", 32)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint32(0x00000000) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint32(0x00000000),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := io.ReadUint16("reserved", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint16(0x0000) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint16(0x0000),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	if pullErr := io.PullContext("exchange"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (exchange)
	exchange, _exchangeErr := CipExchangeParse(io, uint16(len)-uint16(uint16(6)))
	if _exchangeErr != nil {
		return nil, errors.Wrap(_exchangeErr, "Error parsing 'exchange' field")
	}
	if closeErr := io.CloseContext("exchange"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := io.CloseContext("CipRRData"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &CipRRData{
		Exchange: exchange,
		Parent:   &EipPacket{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *CipRRData) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := io.PushContext("CipRRData"); pushErr != nil {
			return pushErr
		}

		// Reserved Field (reserved)
		{
			_err := io.WriteUint32("reserved", 32, uint32(0x00000000))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Reserved Field (reserved)
		{
			_err := io.WriteUint16("reserved", 16, uint16(0x0000))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (exchange)
		if pushErr := io.PushContext("exchange"); pushErr != nil {
			return pushErr
		}
		_exchangeErr := m.Exchange.Serialize(io)
		if popErr := io.PopContext("exchange"); popErr != nil {
			return popErr
		}
		if _exchangeErr != nil {
			return errors.Wrap(_exchangeErr, "Error serializing 'exchange' field")
		}

		if popErr := io.PopContext("CipRRData"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

// Deprecated: the utils.ReadBufferWriteBased should be used instead
func (m *CipRRData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "exchange":
				var data CipExchange
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Exchange = &data
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
	}
}

// Deprecated: the utils.WriteBufferReadBased should be used instead
func (m *CipRRData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.Exchange, xml.StartElement{Name: xml.Name{Local: "exchange"}}); err != nil {
		return err
	}
	return nil
}

func (m CipRRData) String() string {
	return string(m.Box("", 120))
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m CipRRData) Box(name string, width int) utils.AsciiBox {
	boxName := "CipRRData"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Reserved Field (reserved)
		// reserved field can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("reserved", uint32(0x00000000), -1))
		// Reserved Field (reserved)
		// reserved field can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("reserved", uint16(0x0000), -1))
		// Complex field (case complex)
		boxes = append(boxes, m.Exchange.Box("exchange", width-2))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}