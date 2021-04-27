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
	"encoding/hex"
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type ApduDataDeviceDescriptorResponse struct {
	DescriptorType uint8
	Data           []int8
	Parent         *ApduData
}

// The corresponding interface
type IApduDataDeviceDescriptorResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataDeviceDescriptorResponse) ApciType() uint8 {
	return 0xD
}

func (m *ApduDataDeviceDescriptorResponse) InitializeParent(parent *ApduData) {
}

func NewApduDataDeviceDescriptorResponse(descriptorType uint8, data []int8) *ApduData {
	child := &ApduDataDeviceDescriptorResponse{
		DescriptorType: descriptorType,
		Data:           data,
		Parent:         NewApduData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastApduDataDeviceDescriptorResponse(structType interface{}) *ApduDataDeviceDescriptorResponse {
	castFunc := func(typ interface{}) *ApduDataDeviceDescriptorResponse {
		if casted, ok := typ.(ApduDataDeviceDescriptorResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataDeviceDescriptorResponse); ok {
			return casted
		}
		if casted, ok := typ.(ApduData); ok {
			return CastApduDataDeviceDescriptorResponse(casted.Child)
		}
		if casted, ok := typ.(*ApduData); ok {
			return CastApduDataDeviceDescriptorResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataDeviceDescriptorResponse) GetTypeName() string {
	return "ApduDataDeviceDescriptorResponse"
}

func (m *ApduDataDeviceDescriptorResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduDataDeviceDescriptorResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (descriptorType)
	lengthInBits += 6

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *ApduDataDeviceDescriptorResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataDeviceDescriptorResponseParse(io utils.ReadBuffer, dataLength uint8) (*ApduData, error) {
	if pullErr := io.PullContext("ApduDataDeviceDescriptorResponse"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (descriptorType)
	descriptorType, _descriptorTypeErr := io.ReadUint8("descriptorType", 6)
	if _descriptorTypeErr != nil {
		return nil, errors.Wrap(_descriptorTypeErr, "Error parsing 'descriptorType' field")
	}

	// Array field (data)
	if pullErr := io.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	data := make([]int8, utils.InlineIf(bool(bool((dataLength) < (1))), func() uint16 { return uint16(uint16(0)) }, func() uint16 { return uint16(uint16(dataLength) - uint16(uint16(1))) }))
	for curItem := uint16(0); curItem < uint16(utils.InlineIf(bool(bool((dataLength) < (1))), func() uint16 { return uint16(uint16(0)) }, func() uint16 { return uint16(uint16(dataLength) - uint16(uint16(1))) })); curItem++ {
		_item, _err := io.ReadInt8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'data' field")
		}
		data[curItem] = _item
	}
	if closeErr := io.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := io.CloseContext("ApduDataDeviceDescriptorResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataDeviceDescriptorResponse{
		DescriptorType: descriptorType,
		Data:           data,
		Parent:         &ApduData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ApduDataDeviceDescriptorResponse) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := io.PushContext("ApduDataDeviceDescriptorResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (descriptorType)
		descriptorType := uint8(m.DescriptorType)
		_descriptorTypeErr := io.WriteUint8("descriptorType", 6, (descriptorType))
		if _descriptorTypeErr != nil {
			return errors.Wrap(_descriptorTypeErr, "Error serializing 'descriptorType' field")
		}

		// Array Field (data)
		if m.Data != nil {
			if pushErr := io.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Data {
				_elementErr := io.WriteInt8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'data' field")
				}
			}
			if popErr := io.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := io.PopContext("ApduDataDeviceDescriptorResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

// Deprecated: the utils.ReadBufferWriteBased should be used instead
func (m *ApduDataDeviceDescriptorResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "descriptorType":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.DescriptorType = data
			case "data":
				var _encoded string
				if err := d.DecodeElement(&_encoded, &tok); err != nil {
					return err
				}
				_decoded, err := hex.DecodeString(_encoded)
				_len := len(_decoded)
				if err != nil {
					return err
				}
				m.Data = utils.ByteArrayToInt8Array(_decoded[0:_len])
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
func (m *ApduDataDeviceDescriptorResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.DescriptorType, xml.StartElement{Name: xml.Name{Local: "descriptorType"}}); err != nil {
		return err
	}
	_encodedData := hex.EncodeToString(utils.Int8ArrayToByteArray(m.Data))
	_encodedData = strings.ToUpper(_encodedData)
	if err := e.EncodeElement(_encodedData, xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
		return err
	}
	return nil
}

func (m ApduDataDeviceDescriptorResponse) String() string {
	return string(m.Box("", 120))
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m ApduDataDeviceDescriptorResponse) Box(name string, width int) utils.AsciiBox {
	boxName := "ApduDataDeviceDescriptorResponse"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("DescriptorType", m.DescriptorType, -1))
		// Array Field (data)
		if m.Data != nil {
			// Simple array base type int8 will be rendered one by one
			arrayBoxes := make([]utils.AsciiBox, 0)
			for _, _element := range m.Data {
				arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
			}
			boxes = append(boxes, utils.BoxBox("Data", utils.AlignBoxes(arrayBoxes, width-4), 0))
		}
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
