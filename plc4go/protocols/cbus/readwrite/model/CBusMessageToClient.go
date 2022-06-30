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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// CBusMessageToClient is the corresponding interface of CBusMessageToClient
type CBusMessageToClient interface {
	utils.LengthAware
	utils.Serializable
	CBusMessage
	// GetConfirmation returns Confirmation (property field)
	GetConfirmation() Confirmation
}

// CBusMessageToClientExactly can be used when we want exactly this type and not a type which fulfills CBusMessageToClient.
// This is useful for switch cases.
type CBusMessageToClientExactly interface {
	CBusMessageToClient
	isCBusMessageToClient() bool
}

// _CBusMessageToClient is the data-structure of this message
type _CBusMessageToClient struct {
	*_CBusMessage
	Confirmation Confirmation
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CBusMessageToClient) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CBusMessageToClient) InitializeParent(parent CBusMessage) {}

func (m *_CBusMessageToClient) GetParent() CBusMessage {
	return m._CBusMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusMessageToClient) GetConfirmation() Confirmation {
	return m.Confirmation
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCBusMessageToClient factory function for _CBusMessageToClient
func NewCBusMessageToClient(confirmation Confirmation, srchk bool) *_CBusMessageToClient {
	_result := &_CBusMessageToClient{
		Confirmation: confirmation,
		_CBusMessage: NewCBusMessage(srchk),
	}
	_result._CBusMessage._CBusMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCBusMessageToClient(structType interface{}) CBusMessageToClient {
	if casted, ok := structType.(CBusMessageToClient); ok {
		return casted
	}
	if casted, ok := structType.(*CBusMessageToClient); ok {
		return *casted
	}
	return nil
}

func (m *_CBusMessageToClient) GetTypeName() string {
	return "CBusMessageToClient"
}

func (m *_CBusMessageToClient) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CBusMessageToClient) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (confirmation)
	lengthInBits += m.Confirmation.GetLengthInBits()

	return lengthInBits
}

func (m *_CBusMessageToClient) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CBusMessageToClientParse(readBuffer utils.ReadBuffer, response bool, srchk bool) (CBusMessageToClient, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusMessageToClient"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusMessageToClient")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (confirmation)
	if pullErr := readBuffer.PullContext("confirmation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for confirmation")
	}
	_confirmation, _confirmationErr := ConfirmationParse(readBuffer)
	if _confirmationErr != nil {
		return nil, errors.Wrap(_confirmationErr, "Error parsing 'confirmation' field")
	}
	confirmation := _confirmation.(Confirmation)
	if closeErr := readBuffer.CloseContext("confirmation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for confirmation")
	}

	if closeErr := readBuffer.CloseContext("CBusMessageToClient"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusMessageToClient")
	}

	// Create a partially initialized instance
	_child := &_CBusMessageToClient{
		Confirmation: confirmation,
		_CBusMessage: &_CBusMessage{
			Srchk: srchk,
		},
	}
	_child._CBusMessage._CBusMessageChildRequirements = _child
	return _child, nil
}

func (m *_CBusMessageToClient) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusMessageToClient"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CBusMessageToClient")
		}

		// Simple Field (confirmation)
		if pushErr := writeBuffer.PushContext("confirmation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for confirmation")
		}
		_confirmationErr := writeBuffer.WriteSerializable(m.GetConfirmation())
		if popErr := writeBuffer.PopContext("confirmation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for confirmation")
		}
		if _confirmationErr != nil {
			return errors.Wrap(_confirmationErr, "Error serializing 'confirmation' field")
		}

		if popErr := writeBuffer.PopContext("CBusMessageToClient"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CBusMessageToClient")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CBusMessageToClient) isCBusMessageToClient() bool {
	return true
}

func (m *_CBusMessageToClient) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
