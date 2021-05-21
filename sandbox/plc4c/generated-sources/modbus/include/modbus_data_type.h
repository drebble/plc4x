/*
  Licensed to the Apache Software Foundation (ASF) under one
  or more contributor license agreements.  See the NOTICE file
  distributed with this work for additional information
  regarding copyright ownership.  The ASF licenses this file
  to you under the Apache License, Version 2.0 (the
  "License"); you may not use this file except in compliance
  with the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing,
  software distributed under the License is distributed on an
  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
  KIND, either express or implied.  See the License for the
  specific language governing permissions and limitations
  under the License.
*/

#ifndef PLC4C_MODBUS_READ_WRITE_MODBUS_DATA_TYPE_H_
#define PLC4C_MODBUS_READ_WRITE_MODBUS_DATA_TYPE_H_

#include <stdbool.h>
#include <stdint.h>

// Code generated by code-generation. DO NOT EDIT.


#ifdef __cplusplus
extern "C" {
#endif

enum plc4c_modbus_read_write_modbus_data_type {
  plc4c_modbus_read_write_modbus_data_type_BOOL = 1,
  plc4c_modbus_read_write_modbus_data_type_BYTE = 2,
  plc4c_modbus_read_write_modbus_data_type_WORD = 3,
  plc4c_modbus_read_write_modbus_data_type_DWORD = 4,
  plc4c_modbus_read_write_modbus_data_type_LWORD = 5,
  plc4c_modbus_read_write_modbus_data_type_SINT = 6,
  plc4c_modbus_read_write_modbus_data_type_INT = 7,
  plc4c_modbus_read_write_modbus_data_type_DINT = 8,
  plc4c_modbus_read_write_modbus_data_type_LINT = 9,
  plc4c_modbus_read_write_modbus_data_type_USINT = 10,
  plc4c_modbus_read_write_modbus_data_type_UINT = 11,
  plc4c_modbus_read_write_modbus_data_type_UDINT = 12,
  plc4c_modbus_read_write_modbus_data_type_ULINT = 13,
  plc4c_modbus_read_write_modbus_data_type_REAL = 14,
  plc4c_modbus_read_write_modbus_data_type_LREAL = 15,
  plc4c_modbus_read_write_modbus_data_type_TIME = 16,
  plc4c_modbus_read_write_modbus_data_type_LTIME = 17,
  plc4c_modbus_read_write_modbus_data_type_DATE = 18,
  plc4c_modbus_read_write_modbus_data_type_LDATE = 19,
  plc4c_modbus_read_write_modbus_data_type_TIME_OF_DAY = 20,
  plc4c_modbus_read_write_modbus_data_type_LTIME_OF_DAY = 21,
  plc4c_modbus_read_write_modbus_data_type_DATE_AND_TIME = 22,
  plc4c_modbus_read_write_modbus_data_type_LDATE_AND_TIME = 23,
  plc4c_modbus_read_write_modbus_data_type_CHAR = 24,
  plc4c_modbus_read_write_modbus_data_type_WCHAR = 25,
  plc4c_modbus_read_write_modbus_data_type_STRING = 26,
  plc4c_modbus_read_write_modbus_data_type_WSTRING = 27
};
typedef enum plc4c_modbus_read_write_modbus_data_type plc4c_modbus_read_write_modbus_data_type;

// Get an empty NULL-struct
plc4c_modbus_read_write_modbus_data_type plc4c_modbus_read_write_modbus_data_type_null();

plc4c_modbus_read_write_modbus_data_type plc4c_modbus_read_write_modbus_data_type_value_of(char* value_string);

int plc4c_modbus_read_write_modbus_data_type_num_values();

plc4c_modbus_read_write_modbus_data_type plc4c_modbus_read_write_modbus_data_type_value_for_index(int index);

uint8_t plc4c_modbus_read_write_modbus_data_type_get_data_type_size(plc4c_modbus_read_write_modbus_data_type value);
plc4c_modbus_read_write_modbus_data_type plc4c_modbus_read_write_modbus_data_type_get_first_enum_for_field_data_type_size(uint8_t value);

#ifdef __cplusplus
}
#endif

#endif  // PLC4C_MODBUS_READ_WRITE_MODBUS_DATA_TYPE_H_
