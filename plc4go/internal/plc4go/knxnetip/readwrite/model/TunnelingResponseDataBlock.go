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
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type TunnelingResponseDataBlock struct {
	CommunicationChannelId uint8
	SequenceCounter        uint8
	Status                 Status
}

// The corresponding interface
type ITunnelingResponseDataBlock interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

func NewTunnelingResponseDataBlock(communicationChannelId uint8, sequenceCounter uint8, status Status) *TunnelingResponseDataBlock {
	return &TunnelingResponseDataBlock{CommunicationChannelId: communicationChannelId, SequenceCounter: sequenceCounter, Status: status}
}

func CastTunnelingResponseDataBlock(structType interface{}) *TunnelingResponseDataBlock {
	castFunc := func(typ interface{}) *TunnelingResponseDataBlock {
		if casted, ok := typ.(TunnelingResponseDataBlock); ok {
			return &casted
		}
		if casted, ok := typ.(*TunnelingResponseDataBlock); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *TunnelingResponseDataBlock) GetTypeName() string {
	return "TunnelingResponseDataBlock"
}

func (m *TunnelingResponseDataBlock) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *TunnelingResponseDataBlock) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8

	// Simple field (communicationChannelId)
	lengthInBits += 8

	// Simple field (sequenceCounter)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	return lengthInBits
}

func (m *TunnelingResponseDataBlock) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func TunnelingResponseDataBlockParse(io utils.ReadBuffer) (*TunnelingResponseDataBlock, error) {
	if pullErr := io.PullContext("TunnelingResponseDataBlock"); pullErr != nil {
		return nil, pullErr
	}

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength, _structureLengthErr := io.ReadUint8("structureLength", 8)
	_ = structureLength
	if _structureLengthErr != nil {
		return nil, errors.Wrap(_structureLengthErr, "Error parsing 'structureLength' field")
	}

	// Simple Field (communicationChannelId)
	communicationChannelId, _communicationChannelIdErr := io.ReadUint8("communicationChannelId", 8)
	if _communicationChannelIdErr != nil {
		return nil, errors.Wrap(_communicationChannelIdErr, "Error parsing 'communicationChannelId' field")
	}

	// Simple Field (sequenceCounter)
	sequenceCounter, _sequenceCounterErr := io.ReadUint8("sequenceCounter", 8)
	if _sequenceCounterErr != nil {
		return nil, errors.Wrap(_sequenceCounterErr, "Error parsing 'sequenceCounter' field")
	}

	if pullErr := io.PullContext("status"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (status)
	status, _statusErr := StatusParse(io)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field")
	}
	if closeErr := io.CloseContext("status"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := io.CloseContext("TunnelingResponseDataBlock"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewTunnelingResponseDataBlock(communicationChannelId, sequenceCounter, status), nil
}

func (m *TunnelingResponseDataBlock) Serialize(io utils.WriteBuffer) error {
	if pushErr := io.PushContext("TunnelingResponseDataBlock"); pushErr != nil {
		return pushErr
	}

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength := uint8(uint8(m.LengthInBytes()))
	_structureLengthErr := io.WriteUint8("structureLength", 8, (structureLength))
	if _structureLengthErr != nil {
		return errors.Wrap(_structureLengthErr, "Error serializing 'structureLength' field")
	}

	// Simple Field (communicationChannelId)
	communicationChannelId := uint8(m.CommunicationChannelId)
	_communicationChannelIdErr := io.WriteUint8("communicationChannelId", 8, (communicationChannelId))
	if _communicationChannelIdErr != nil {
		return errors.Wrap(_communicationChannelIdErr, "Error serializing 'communicationChannelId' field")
	}

	// Simple Field (sequenceCounter)
	sequenceCounter := uint8(m.SequenceCounter)
	_sequenceCounterErr := io.WriteUint8("sequenceCounter", 8, (sequenceCounter))
	if _sequenceCounterErr != nil {
		return errors.Wrap(_sequenceCounterErr, "Error serializing 'sequenceCounter' field")
	}

	// Simple Field (status)
	if pushErr := io.PushContext("status"); pushErr != nil {
		return pushErr
	}
	_statusErr := m.Status.Serialize(io)
	if popErr := io.PopContext("status"); popErr != nil {
		return popErr
	}
	if _statusErr != nil {
		return errors.Wrap(_statusErr, "Error serializing 'status' field")
	}

	if popErr := io.PopContext("TunnelingResponseDataBlock"); popErr != nil {
		return popErr
	}
	return nil
}

// Deprecated: the utils.ReadBufferWriteBased should be used instead
func (m *TunnelingResponseDataBlock) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	for {
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "communicationChannelId":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.CommunicationChannelId = data
			case "sequenceCounter":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.SequenceCounter = data
			case "status":
				var data Status
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Status = data
			}
		}
	}
}

// Deprecated: the utils.WriteBufferReadBased should be used instead
func (m *TunnelingResponseDataBlock) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := "org.apache.plc4x.java.knxnetip.readwrite.TunnelingResponseDataBlock"
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.CommunicationChannelId, xml.StartElement{Name: xml.Name{Local: "communicationChannelId"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.SequenceCounter, xml.StartElement{Name: xml.Name{Local: "sequenceCounter"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Status, xml.StartElement{Name: xml.Name{Local: "status"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func (m TunnelingResponseDataBlock) String() string {
	return string(m.Box("", 120))
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m TunnelingResponseDataBlock) Box(name string, width int) utils.AsciiBox {
	boxName := "TunnelingResponseDataBlock"
	if name != "" {
		boxName += "/" + name
	}
	boxes := make([]utils.AsciiBox, 0)
	// Implicit Field (structureLength)
	structureLength := uint8(uint8(m.LengthInBytes()))
	// uint8 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("StructureLength", structureLength, -1))
	// Simple field (case simple)
	// uint8 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("CommunicationChannelId", m.CommunicationChannelId, -1))
	// Simple field (case simple)
	// uint8 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("SequenceCounter", m.SequenceCounter, -1))
	// Complex field (case complex)
	boxes = append(boxes, m.Status.Box("status", width-2))
	return utils.BoxBox(boxName, utils.AlignBoxes(boxes, width-2), 0)
}
