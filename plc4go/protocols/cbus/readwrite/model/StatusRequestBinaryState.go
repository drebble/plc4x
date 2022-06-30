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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// StatusRequestBinaryState is the corresponding interface of StatusRequestBinaryState
type StatusRequestBinaryState interface {
	utils.LengthAware
	utils.Serializable
	StatusRequest
	// GetApplication returns Application (property field)
	GetApplication() byte
}

// StatusRequestBinaryStateExactly can be used when we want exactly this type and not a type which fulfills StatusRequestBinaryState.
// This is useful for switch cases.
type StatusRequestBinaryStateExactly interface {
	StatusRequestBinaryState
	isStatusRequestBinaryState() bool
}

// _StatusRequestBinaryState is the data-structure of this message
type _StatusRequestBinaryState struct {
	*_StatusRequest
	Application byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_StatusRequestBinaryState) InitializeParent(parent StatusRequest, statusType byte) {
	m.StatusType = statusType
}

func (m *_StatusRequestBinaryState) GetParent() StatusRequest {
	return m._StatusRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_StatusRequestBinaryState) GetApplication() byte {
	return m.Application
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewStatusRequestBinaryState factory function for _StatusRequestBinaryState
func NewStatusRequestBinaryState(application byte, statusType byte) *_StatusRequestBinaryState {
	_result := &_StatusRequestBinaryState{
		Application:    application,
		_StatusRequest: NewStatusRequest(statusType),
	}
	_result._StatusRequest._StatusRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastStatusRequestBinaryState(structType interface{}) StatusRequestBinaryState {
	if casted, ok := structType.(StatusRequestBinaryState); ok {
		return casted
	}
	if casted, ok := structType.(*StatusRequestBinaryState); ok {
		return *casted
	}
	return nil
}

func (m *_StatusRequestBinaryState) GetTypeName() string {
	return "StatusRequestBinaryState"
}

func (m *_StatusRequestBinaryState) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_StatusRequestBinaryState) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (application)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *_StatusRequestBinaryState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func StatusRequestBinaryStateParse(readBuffer utils.ReadBuffer) (StatusRequestBinaryState, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("StatusRequestBinaryState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for StatusRequestBinaryState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadByte("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != byte(0x7A) {
			log.Info().Fields(map[string]interface{}{
				"expected value": byte(0x7A),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
		}
	}

	// Simple Field (application)
	_application, _applicationErr := readBuffer.ReadByte("application")
	if _applicationErr != nil {
		return nil, errors.Wrap(_applicationErr, "Error parsing 'application' field")
	}
	application := _application

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadByte("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != byte(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": byte(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
		}
	}

	if closeErr := readBuffer.CloseContext("StatusRequestBinaryState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for StatusRequestBinaryState")
	}

	// Create a partially initialized instance
	_child := &_StatusRequestBinaryState{
		Application:    application,
		_StatusRequest: &_StatusRequest{},
	}
	_child._StatusRequest._StatusRequestChildRequirements = _child
	return _child, nil
}

func (m *_StatusRequestBinaryState) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("StatusRequestBinaryState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for StatusRequestBinaryState")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteByte("reserved", byte(0x7A))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (application)
		application := byte(m.GetApplication())
		_applicationErr := writeBuffer.WriteByte("application", (application))
		if _applicationErr != nil {
			return errors.Wrap(_applicationErr, "Error serializing 'application' field")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteByte("reserved", byte(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		if popErr := writeBuffer.PopContext("StatusRequestBinaryState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for StatusRequestBinaryState")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_StatusRequestBinaryState) isStatusRequestBinaryState() bool {
	return true
}

func (m *_StatusRequestBinaryState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
