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
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type SysexCommandSysexRealtime struct {
	Parent *SysexCommand
}

// The corresponding interface
type ISysexCommandSysexRealtime interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *SysexCommandSysexRealtime) CommandType() uint8 {
	return 0x7F
}

func (m *SysexCommandSysexRealtime) Response() bool {
	return false
}

func (m *SysexCommandSysexRealtime) InitializeParent(parent *SysexCommand) {
}

func NewSysexCommandSysexRealtime() *SysexCommand {
	child := &SysexCommandSysexRealtime{
		Parent: NewSysexCommand(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastSysexCommandSysexRealtime(structType interface{}) *SysexCommandSysexRealtime {
	castFunc := func(typ interface{}) *SysexCommandSysexRealtime {
		if casted, ok := typ.(SysexCommandSysexRealtime); ok {
			return &casted
		}
		if casted, ok := typ.(*SysexCommandSysexRealtime); ok {
			return casted
		}
		if casted, ok := typ.(SysexCommand); ok {
			return CastSysexCommandSysexRealtime(casted.Child)
		}
		if casted, ok := typ.(*SysexCommand); ok {
			return CastSysexCommandSysexRealtime(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *SysexCommandSysexRealtime) GetTypeName() string {
	return "SysexCommandSysexRealtime"
}

func (m *SysexCommandSysexRealtime) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *SysexCommandSysexRealtime) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *SysexCommandSysexRealtime) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func SysexCommandSysexRealtimeParse(readBuffer utils.ReadBuffer) (*SysexCommand, error) {
	if pullErr := readBuffer.PullContext("SysexCommandSysexRealtime"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("SysexCommandSysexRealtime"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &SysexCommandSysexRealtime{
		Parent: &SysexCommand{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *SysexCommandSysexRealtime) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandSysexRealtime"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("SysexCommandSysexRealtime"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *SysexCommandSysexRealtime) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
