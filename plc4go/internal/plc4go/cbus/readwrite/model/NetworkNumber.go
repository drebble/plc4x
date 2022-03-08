/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type NetworkNumber struct {
	Number uint8
}

// The corresponding interface
type INetworkNumber interface {
	// GetNumber returns Number (property field)
	GetNumber() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *NetworkNumber) GetNumber() uint8 {
	return m.Number
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNetworkNumber factory function for NetworkNumber
func NewNetworkNumber(number uint8) *NetworkNumber {
	return &NetworkNumber{Number: number}
}

func CastNetworkNumber(structType interface{}) *NetworkNumber {
	if casted, ok := structType.(NetworkNumber); ok {
		return &casted
	}
	if casted, ok := structType.(*NetworkNumber); ok {
		return casted
	}
	return nil
}

func (m *NetworkNumber) GetTypeName() string {
	return "NetworkNumber"
}

func (m *NetworkNumber) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *NetworkNumber) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (number)
	lengthInBits += 8

	return lengthInBits
}

func (m *NetworkNumber) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NetworkNumberParse(readBuffer utils.ReadBuffer) (*NetworkNumber, error) {
	if pullErr := readBuffer.PullContext("NetworkNumber"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (number)
	_number, _numberErr := readBuffer.ReadUint8("number", 8)
	if _numberErr != nil {
		return nil, errors.Wrap(_numberErr, "Error parsing 'number' field")
	}
	number := _number

	if closeErr := readBuffer.CloseContext("NetworkNumber"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewNetworkNumber(number), nil
}

func (m *NetworkNumber) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("NetworkNumber"); pushErr != nil {
		return pushErr
	}

	// Simple Field (number)
	number := uint8(m.Number)
	_numberErr := writeBuffer.WriteUint8("number", 8, (number))
	if _numberErr != nil {
		return errors.Wrap(_numberErr, "Error serializing 'number' field")
	}

	if popErr := writeBuffer.PopContext("NetworkNumber"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *NetworkNumber) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}