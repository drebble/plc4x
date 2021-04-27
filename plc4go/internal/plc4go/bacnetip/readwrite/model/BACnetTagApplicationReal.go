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
type BACnetTagApplicationReal struct {
	Value  float32
	Parent *BACnetTag
}

// The corresponding interface
type IBACnetTagApplicationReal interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetTagApplicationReal) ContextSpecificTag() uint8 {
	return 0
}

func (m *BACnetTagApplicationReal) InitializeParent(parent *BACnetTag, typeOrTagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8) {
	m.Parent.TypeOrTagNumber = typeOrTagNumber
	m.Parent.LengthValueType = lengthValueType
	m.Parent.ExtTagNumber = extTagNumber
	m.Parent.ExtLength = extLength
}

func NewBACnetTagApplicationReal(value float32, typeOrTagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8) *BACnetTag {
	child := &BACnetTagApplicationReal{
		Value:  value,
		Parent: NewBACnetTag(typeOrTagNumber, lengthValueType, extTagNumber, extLength),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetTagApplicationReal(structType interface{}) *BACnetTagApplicationReal {
	castFunc := func(typ interface{}) *BACnetTagApplicationReal {
		if casted, ok := typ.(BACnetTagApplicationReal); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetTagApplicationReal); ok {
			return casted
		}
		if casted, ok := typ.(BACnetTag); ok {
			return CastBACnetTagApplicationReal(casted.Child)
		}
		if casted, ok := typ.(*BACnetTag); ok {
			return CastBACnetTagApplicationReal(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetTagApplicationReal) GetTypeName() string {
	return "BACnetTagApplicationReal"
}

func (m *BACnetTagApplicationReal) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetTagApplicationReal) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (value)
	lengthInBits += 32

	return lengthInBits
}

func (m *BACnetTagApplicationReal) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetTagApplicationRealParse(io utils.ReadBuffer, lengthValueType uint8, extLength uint8) (*BACnetTag, error) {
	if pullErr := io.PullContext("BACnetTagApplicationReal"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (value)
	value, _valueErr := io.ReadFloat32("value", true, 8, 23)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}

	if closeErr := io.CloseContext("BACnetTagApplicationReal"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetTagApplicationReal{
		Value:  value,
		Parent: &BACnetTag{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetTagApplicationReal) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := io.PushContext("BACnetTagApplicationReal"); pushErr != nil {
			return pushErr
		}

		// Simple Field (value)
		value := float32(m.Value)
		_valueErr := io.WriteFloat32("value", 32, (value))
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := io.PopContext("BACnetTagApplicationReal"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

// Deprecated: the utils.ReadBufferWriteBased should be used instead
func (m *BACnetTagApplicationReal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "value":
				var data float32
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Value = data
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
func (m *BACnetTagApplicationReal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.Value, xml.StartElement{Name: xml.Name{Local: "value"}}); err != nil {
		return err
	}
	return nil
}

func (m BACnetTagApplicationReal) String() string {
	return string(m.Box("", 120))
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m BACnetTagApplicationReal) Box(name string, width int) utils.AsciiBox {
	boxName := "BACnetTagApplicationReal"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// float32 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("Value", m.Value, -1))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
