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

// CBusMessage is the corresponding interface of CBusMessage
type CBusMessage interface {
	utils.LengthAware
	utils.Serializable
	// GetResponse returns Response (discriminator field)
	GetResponse() bool
}

// CBusMessageExactly can be used when we want exactly this type and not a type which fulfills CBusMessage.
// This is useful for switch cases.
type CBusMessageExactly interface {
	CBusMessage
	isCBusMessage() bool
}

// _CBusMessage is the data-structure of this message
type _CBusMessage struct {
	_CBusMessageChildRequirements

	// Arguments.
	Srchk bool
}

type _CBusMessageChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	GetResponse() bool
}

type CBusMessageParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child CBusMessage, serializeChildFunction func() error) error
	GetTypeName() string
}

type CBusMessageChild interface {
	utils.Serializable
	InitializeParent(parent CBusMessage)
	GetParent() *CBusMessage

	GetTypeName() string
	CBusMessage
}

// NewCBusMessage factory function for _CBusMessage
func NewCBusMessage(srchk bool) *_CBusMessage {
	return &_CBusMessage{Srchk: srchk}
}

// Deprecated: use the interface for direct cast
func CastCBusMessage(structType interface{}) CBusMessage {
	if casted, ok := structType.(CBusMessage); ok {
		return casted
	}
	if casted, ok := structType.(*CBusMessage); ok {
		return *casted
	}
	return nil
}

func (m *_CBusMessage) GetTypeName() string {
	return "CBusMessage"
}

func (m *_CBusMessage) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_CBusMessage) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CBusMessageParse(readBuffer utils.ReadBuffer, response bool, srchk bool) (CBusMessage, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type CBusMessageChildSerializeRequirement interface {
		CBusMessage
		InitializeParent(CBusMessage)
		GetParent() CBusMessage
	}
	var _childTemp interface{}
	var _child CBusMessageChildSerializeRequirement
	var typeSwitchError error
	switch {
	case response == bool(false): // CBusMessageToServer
		_childTemp, typeSwitchError = CBusMessageToServerParse(readBuffer, response, srchk)
	case response == bool(true): // CBusMessageToClient
		_childTemp, typeSwitchError = CBusMessageToClientParse(readBuffer, response, srchk)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}
	_child = _childTemp.(CBusMessageChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("CBusMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusMessage")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_CBusMessage) SerializeParent(writeBuffer utils.WriteBuffer, child CBusMessage, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("CBusMessage"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CBusMessage")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("CBusMessage"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CBusMessage")
	}
	return nil
}

func (m *_CBusMessage) isCBusMessage() bool {
	return true
}

func (m *_CBusMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
