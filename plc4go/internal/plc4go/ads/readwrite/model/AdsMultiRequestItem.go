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
type AdsMultiRequestItem struct {
	Child IAdsMultiRequestItemChild
}

// The corresponding interface
type IAdsMultiRequestItem interface {
	IndexGroup() uint32
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IAdsMultiRequestItemParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IAdsMultiRequestItem, serializeChildFunction func() error) error
	GetTypeName() string
}

type IAdsMultiRequestItemChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *AdsMultiRequestItem)
	GetTypeName() string
	IAdsMultiRequestItem
}

func NewAdsMultiRequestItem() *AdsMultiRequestItem {
	return &AdsMultiRequestItem{}
}

func CastAdsMultiRequestItem(structType interface{}) *AdsMultiRequestItem {
	castFunc := func(typ interface{}) *AdsMultiRequestItem {
		if casted, ok := typ.(AdsMultiRequestItem); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsMultiRequestItem); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsMultiRequestItem) GetTypeName() string {
	return "AdsMultiRequestItem"
}

func (m *AdsMultiRequestItem) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AdsMultiRequestItem) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *AdsMultiRequestItem) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *AdsMultiRequestItem) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsMultiRequestItemParse(readBuffer utils.ReadBuffer, indexGroup uint32) (*AdsMultiRequestItem, error) {
	if pullErr := readBuffer.PullContext("AdsMultiRequestItem"); pullErr != nil {
		return nil, pullErr
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *AdsMultiRequestItem
	var typeSwitchError error
	switch {
	case indexGroup == 61568: // AdsMultiRequestItemRead
		_parent, typeSwitchError = AdsMultiRequestItemReadParse(readBuffer)
	case indexGroup == 61569: // AdsMultiRequestItemWrite
		_parent, typeSwitchError = AdsMultiRequestItemWriteParse(readBuffer)
	case indexGroup == 61570: // AdsMultiRequestItemReadWrite
		_parent, typeSwitchError = AdsMultiRequestItemReadWriteParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("AdsMultiRequestItem"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *AdsMultiRequestItem) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *AdsMultiRequestItem) SerializeParent(writeBuffer utils.WriteBuffer, child IAdsMultiRequestItem, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("AdsMultiRequestItem"); pushErr != nil {
		return pushErr
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("AdsMultiRequestItem"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *AdsMultiRequestItem) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
