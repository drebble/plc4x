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
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type TDataConnectedReq struct {
	Parent *CEMI
}

// The corresponding interface
type ITDataConnectedReq interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *TDataConnectedReq) MessageCode() uint8 {
	return 0x41
}

func (m *TDataConnectedReq) InitializeParent(parent *CEMI) {
}

func NewTDataConnectedReq() *CEMI {
	child := &TDataConnectedReq{
		Parent: NewCEMI(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastTDataConnectedReq(structType interface{}) *TDataConnectedReq {
	castFunc := func(typ interface{}) *TDataConnectedReq {
		if casted, ok := typ.(TDataConnectedReq); ok {
			return &casted
		}
		if casted, ok := typ.(*TDataConnectedReq); ok {
			return casted
		}
		if casted, ok := typ.(CEMI); ok {
			return CastTDataConnectedReq(casted.Child)
		}
		if casted, ok := typ.(*CEMI); ok {
			return CastTDataConnectedReq(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *TDataConnectedReq) GetTypeName() string {
	return "TDataConnectedReq"
}

func (m *TDataConnectedReq) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *TDataConnectedReq) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *TDataConnectedReq) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func TDataConnectedReqParse(io utils.ReadBuffer) (*CEMI, error) {
	if pullErr := io.PullContext("TDataConnectedReq"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := io.CloseContext("TDataConnectedReq"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &TDataConnectedReq{
		Parent: &CEMI{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *TDataConnectedReq) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := io.PushContext("TDataConnectedReq"); pushErr != nil {
			return pushErr
		}

		if popErr := io.PopContext("TDataConnectedReq"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

// Deprecated: the utils.ReadBufferWriteBased should be used instead
func (m *TDataConnectedReq) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
func (m *TDataConnectedReq) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func (m TDataConnectedReq) String() string {
	return string(m.Box("", 120))
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m TDataConnectedReq) Box(name string, width int) utils.AsciiBox {
	boxName := "TDataConnectedReq"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
