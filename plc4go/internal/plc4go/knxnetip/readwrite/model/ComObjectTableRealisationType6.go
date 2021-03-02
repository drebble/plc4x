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
	"errors"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"io"
)

// The data-structure of this message
type ComObjectTableRealisationType6 struct {
	ComObjectDescriptors *GroupObjectDescriptorRealisationType6
	Parent               *ComObjectTable
	IComObjectTableRealisationType6
}

// The corresponding interface
type IComObjectTableRealisationType6 interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ComObjectTableRealisationType6) FirmwareType() FirmwareType {
	return FirmwareType_SYSTEM_300
}

func (m *ComObjectTableRealisationType6) InitializeParent(parent *ComObjectTable) {
}

func NewComObjectTableRealisationType6(comObjectDescriptors *GroupObjectDescriptorRealisationType6) *ComObjectTable {
	child := &ComObjectTableRealisationType6{
		ComObjectDescriptors: comObjectDescriptors,
		Parent:               NewComObjectTable(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastComObjectTableRealisationType6(structType interface{}) *ComObjectTableRealisationType6 {
	castFunc := func(typ interface{}) *ComObjectTableRealisationType6 {
		if casted, ok := typ.(ComObjectTableRealisationType6); ok {
			return &casted
		}
		if casted, ok := typ.(*ComObjectTableRealisationType6); ok {
			return casted
		}
		if casted, ok := typ.(ComObjectTable); ok {
			return CastComObjectTableRealisationType6(casted.Child)
		}
		if casted, ok := typ.(*ComObjectTable); ok {
			return CastComObjectTableRealisationType6(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ComObjectTableRealisationType6) GetTypeName() string {
	return "ComObjectTableRealisationType6"
}

func (m *ComObjectTableRealisationType6) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (comObjectDescriptors)
	lengthInBits += m.ComObjectDescriptors.LengthInBits()

	return lengthInBits
}

func (m *ComObjectTableRealisationType6) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ComObjectTableRealisationType6Parse(io *utils.ReadBuffer) (*ComObjectTable, error) {

	// Simple Field (comObjectDescriptors)
	comObjectDescriptors, _comObjectDescriptorsErr := GroupObjectDescriptorRealisationType6Parse(io)
	if _comObjectDescriptorsErr != nil {
		return nil, errors.New("Error parsing 'comObjectDescriptors' field " + _comObjectDescriptorsErr.Error())
	}

	// Create a partially initialized instance
	_child := &ComObjectTableRealisationType6{
		ComObjectDescriptors: comObjectDescriptors,
		Parent:               &ComObjectTable{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ComObjectTableRealisationType6) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Simple Field (comObjectDescriptors)
		_comObjectDescriptorsErr := m.ComObjectDescriptors.Serialize(io)
		if _comObjectDescriptorsErr != nil {
			return errors.New("Error serializing 'comObjectDescriptors' field " + _comObjectDescriptorsErr.Error())
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *ComObjectTableRealisationType6) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "comObjectDescriptors":
				var data *GroupObjectDescriptorRealisationType6
				if err := d.DecodeElement(data, &tok); err != nil {
					return err
				}
				m.ComObjectDescriptors = data
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

func (m *ComObjectTableRealisationType6) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.ComObjectDescriptors, xml.StartElement{Name: xml.Name{Local: "comObjectDescriptors"}}); err != nil {
		return err
	}
	return nil
}