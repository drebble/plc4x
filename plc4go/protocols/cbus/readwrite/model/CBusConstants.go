/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const CBusConstants_CBUSTCPDEFAULTPORT uint16 = uint16(10001)

// CBusConstants is the corresponding interface of CBusConstants
type CBusConstants interface {
	utils.LengthAware
	utils.Serializable
}

// CBusConstantsExactly can be used when we want exactly this type and not a type which fulfills CBusConstants.
// This is useful for switch cases.
type CBusConstantsExactly interface {
	CBusConstants
	isCBusConstants() bool
}

// _CBusConstants is the data-structure of this message
type _CBusConstants struct {
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_CBusConstants) GetCbusTcpDefaultPort() uint16 {
	return CBusConstants_CBUSTCPDEFAULTPORT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCBusConstants factory function for _CBusConstants
func NewCBusConstants() *_CBusConstants {
	return &_CBusConstants{}
}

// Deprecated: use the interface for direct cast
func CastCBusConstants(structType interface{}) CBusConstants {
	if casted, ok := structType.(CBusConstants); ok {
		return casted
	}
	if casted, ok := structType.(*CBusConstants); ok {
		return *casted
	}
	return nil
}

func (m *_CBusConstants) GetTypeName() string {
	return "CBusConstants"
}

func (m *_CBusConstants) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CBusConstants) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Const Field (cbusTcpDefaultPort)
	lengthInBits += 16

	return lengthInBits
}

func (m *_CBusConstants) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CBusConstantsParse(readBuffer utils.ReadBuffer) (CBusConstants, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusConstants"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusConstants")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (cbusTcpDefaultPort)
	cbusTcpDefaultPort, _cbusTcpDefaultPortErr := readBuffer.ReadUint16("cbusTcpDefaultPort", 16)
	if _cbusTcpDefaultPortErr != nil {
		return nil, errors.Wrap(_cbusTcpDefaultPortErr, "Error parsing 'cbusTcpDefaultPort' field")
	}
	if cbusTcpDefaultPort != CBusConstants_CBUSTCPDEFAULTPORT {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", CBusConstants_CBUSTCPDEFAULTPORT) + " but got " + fmt.Sprintf("%d", cbusTcpDefaultPort))
	}

	if closeErr := readBuffer.CloseContext("CBusConstants"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusConstants")
	}

	// Create the instance
	return NewCBusConstants(), nil
}

func (m *_CBusConstants) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("CBusConstants"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CBusConstants")
	}

	// Const Field (cbusTcpDefaultPort)
	_cbusTcpDefaultPortErr := writeBuffer.WriteUint16("cbusTcpDefaultPort", 16, 10001)
	if _cbusTcpDefaultPortErr != nil {
		return errors.Wrap(_cbusTcpDefaultPortErr, "Error serializing 'cbusTcpDefaultPort' field")
	}

	if popErr := writeBuffer.PopContext("CBusConstants"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CBusConstants")
	}
	return nil
}

func (m *_CBusConstants) isCBusConstants() bool {
	return true
}

func (m *_CBusConstants) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
