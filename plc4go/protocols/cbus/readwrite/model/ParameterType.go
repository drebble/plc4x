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
	"encoding/binary"

	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ParameterType is an enum
type ParameterType uint8

type IParameterType interface {
	utils.Serializable
}

const (
	ParameterType_UNKNOWN                               ParameterType = 0
	ParameterType_APPLICATION_ADDRESS_1                 ParameterType = 1
	ParameterType_APPLICATION_ADDRESS_2                 ParameterType = 2
	ParameterType_INTERFACE_OPTIONS_1                   ParameterType = 3
	ParameterType_INTERFACE_OPTIONS_2                   ParameterType = 4
	ParameterType_INTERFACE_OPTIONS_3                   ParameterType = 5
	ParameterType_BAUD_RATE_SELECTOR                    ParameterType = 6
	ParameterType_INTERFACE_OPTIONS_1_POWER_UP_SETTINGS ParameterType = 7
	ParameterType_CUSTOM_MANUFACTURER                   ParameterType = 8
	ParameterType_SERIAL_NUMBER                         ParameterType = 9
	ParameterType_CUSTOM_TYPE                           ParameterType = 10
)

var ParameterTypeValues []ParameterType

func init() {
	_ = errors.New
	ParameterTypeValues = []ParameterType{
		ParameterType_UNKNOWN,
		ParameterType_APPLICATION_ADDRESS_1,
		ParameterType_APPLICATION_ADDRESS_2,
		ParameterType_INTERFACE_OPTIONS_1,
		ParameterType_INTERFACE_OPTIONS_2,
		ParameterType_INTERFACE_OPTIONS_3,
		ParameterType_BAUD_RATE_SELECTOR,
		ParameterType_INTERFACE_OPTIONS_1_POWER_UP_SETTINGS,
		ParameterType_CUSTOM_MANUFACTURER,
		ParameterType_SERIAL_NUMBER,
		ParameterType_CUSTOM_TYPE,
	}
}

func ParameterTypeByValue(value uint8) (enum ParameterType, ok bool) {
	switch value {
	case 0:
		return ParameterType_UNKNOWN, true
	case 1:
		return ParameterType_APPLICATION_ADDRESS_1, true
	case 10:
		return ParameterType_CUSTOM_TYPE, true
	case 2:
		return ParameterType_APPLICATION_ADDRESS_2, true
	case 3:
		return ParameterType_INTERFACE_OPTIONS_1, true
	case 4:
		return ParameterType_INTERFACE_OPTIONS_2, true
	case 5:
		return ParameterType_INTERFACE_OPTIONS_3, true
	case 6:
		return ParameterType_BAUD_RATE_SELECTOR, true
	case 7:
		return ParameterType_INTERFACE_OPTIONS_1_POWER_UP_SETTINGS, true
	case 8:
		return ParameterType_CUSTOM_MANUFACTURER, true
	case 9:
		return ParameterType_SERIAL_NUMBER, true
	}
	return 0, false
}

func ParameterTypeByName(value string) (enum ParameterType, ok bool) {
	switch value {
	case "UNKNOWN":
		return ParameterType_UNKNOWN, true
	case "APPLICATION_ADDRESS_1":
		return ParameterType_APPLICATION_ADDRESS_1, true
	case "CUSTOM_TYPE":
		return ParameterType_CUSTOM_TYPE, true
	case "APPLICATION_ADDRESS_2":
		return ParameterType_APPLICATION_ADDRESS_2, true
	case "INTERFACE_OPTIONS_1":
		return ParameterType_INTERFACE_OPTIONS_1, true
	case "INTERFACE_OPTIONS_2":
		return ParameterType_INTERFACE_OPTIONS_2, true
	case "INTERFACE_OPTIONS_3":
		return ParameterType_INTERFACE_OPTIONS_3, true
	case "BAUD_RATE_SELECTOR":
		return ParameterType_BAUD_RATE_SELECTOR, true
	case "INTERFACE_OPTIONS_1_POWER_UP_SETTINGS":
		return ParameterType_INTERFACE_OPTIONS_1_POWER_UP_SETTINGS, true
	case "CUSTOM_MANUFACTURER":
		return ParameterType_CUSTOM_MANUFACTURER, true
	case "SERIAL_NUMBER":
		return ParameterType_SERIAL_NUMBER, true
	}
	return 0, false
}

func ParameterTypeKnows(value uint8) bool {
	for _, typeValue := range ParameterTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastParameterType(structType interface{}) ParameterType {
	castFunc := func(typ interface{}) ParameterType {
		if sParameterType, ok := typ.(ParameterType); ok {
			return sParameterType
		}
		return 0
	}
	return castFunc(structType)
}

func (m ParameterType) GetLengthInBits() uint16 {
	return 8
}

func (m ParameterType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ParameterTypeParse(theBytes []byte) (ParameterType, error) {
	return ParameterTypeParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian))) // TODO: get endianness from mspec
}

func ParameterTypeParseWithBuffer(readBuffer utils.ReadBuffer) (ParameterType, error) {
	val, err := readBuffer.ReadUint8("ParameterType", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ParameterType")
	}
	if enum, ok := ParameterTypeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return ParameterType(val), nil
	} else {
		return enum, nil
	}
}

func (e ParameterType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ParameterType) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("ParameterType", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ParameterType) PLC4XEnumName() string {
	switch e {
	case ParameterType_UNKNOWN:
		return "UNKNOWN"
	case ParameterType_APPLICATION_ADDRESS_1:
		return "APPLICATION_ADDRESS_1"
	case ParameterType_CUSTOM_TYPE:
		return "CUSTOM_TYPE"
	case ParameterType_APPLICATION_ADDRESS_2:
		return "APPLICATION_ADDRESS_2"
	case ParameterType_INTERFACE_OPTIONS_1:
		return "INTERFACE_OPTIONS_1"
	case ParameterType_INTERFACE_OPTIONS_2:
		return "INTERFACE_OPTIONS_2"
	case ParameterType_INTERFACE_OPTIONS_3:
		return "INTERFACE_OPTIONS_3"
	case ParameterType_BAUD_RATE_SELECTOR:
		return "BAUD_RATE_SELECTOR"
	case ParameterType_INTERFACE_OPTIONS_1_POWER_UP_SETTINGS:
		return "INTERFACE_OPTIONS_1_POWER_UP_SETTINGS"
	case ParameterType_CUSTOM_MANUFACTURER:
		return "CUSTOM_MANUFACTURER"
	case ParameterType_SERIAL_NUMBER:
		return "SERIAL_NUMBER"
	}
	return ""
}

func (e ParameterType) String() string {
	return e.PLC4XEnumName()
}