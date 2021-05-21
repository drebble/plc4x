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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type S7PayloadWriteVarRequest struct {
	Items  []*S7VarPayloadDataItem
	Parent *S7Payload
}

// The corresponding interface
type IS7PayloadWriteVarRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *S7PayloadWriteVarRequest) ParameterParameterType() uint8 {
	return 0x05
}

func (m *S7PayloadWriteVarRequest) MessageType() uint8 {
	return 0x01
}

func (m *S7PayloadWriteVarRequest) InitializeParent(parent *S7Payload) {
}

func NewS7PayloadWriteVarRequest(items []*S7VarPayloadDataItem) *S7Payload {
	child := &S7PayloadWriteVarRequest{
		Items:  items,
		Parent: NewS7Payload(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastS7PayloadWriteVarRequest(structType interface{}) *S7PayloadWriteVarRequest {
	castFunc := func(typ interface{}) *S7PayloadWriteVarRequest {
		if casted, ok := typ.(S7PayloadWriteVarRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*S7PayloadWriteVarRequest); ok {
			return casted
		}
		if casted, ok := typ.(S7Payload); ok {
			return CastS7PayloadWriteVarRequest(casted.Child)
		}
		if casted, ok := typ.(*S7Payload); ok {
			return CastS7PayloadWriteVarRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7PayloadWriteVarRequest) GetTypeName() string {
	return "S7PayloadWriteVarRequest"
}

func (m *S7PayloadWriteVarRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *S7PayloadWriteVarRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Array field
	if len(m.Items) > 0 {
		for i, element := range m.Items {
			last := i == len(m.Items)-1
			lengthInBits += element.LengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *S7PayloadWriteVarRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7PayloadWriteVarRequestParse(readBuffer utils.ReadBuffer, parameter *S7Parameter) (*S7Payload, error) {
	if pullErr := readBuffer.PullContext("S7PayloadWriteVarRequest"); pullErr != nil {
		return nil, pullErr
	}

	// Array field (items)
	if pullErr := readBuffer.PullContext("items", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	items := make([]*S7VarPayloadDataItem, uint16(len(CastS7ParameterWriteVarRequest(parameter).Items)))
	for curItem := uint16(0); curItem < uint16(uint16(len(CastS7ParameterWriteVarRequest(parameter).Items))); curItem++ {
		lastItem := curItem == uint16((len(CastS7ParameterWriteVarRequest(parameter).Items))-1)
		_item, _err := S7VarPayloadDataItemParse(readBuffer, lastItem)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'items' field")
		}
		items[curItem] = _item
	}
	if closeErr := readBuffer.CloseContext("items", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("S7PayloadWriteVarRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7PayloadWriteVarRequest{
		Items:  items,
		Parent: &S7Payload{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *S7PayloadWriteVarRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadWriteVarRequest"); pushErr != nil {
			return pushErr
		}

		// Array Field (items)
		if m.Items != nil {
			if pushErr := writeBuffer.PushContext("items", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			itemCount := uint16(len(m.Items))
			var curItem uint16 = 0
			for _, _element := range m.Items {
				var lastItem bool = curItem == (itemCount - 1)
				_elementErr := _element.Serialize(writeBuffer, lastItem)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'items' field")
				}
				curItem++
			}
			if popErr := writeBuffer.PopContext("items", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("S7PayloadWriteVarRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *S7PayloadWriteVarRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
