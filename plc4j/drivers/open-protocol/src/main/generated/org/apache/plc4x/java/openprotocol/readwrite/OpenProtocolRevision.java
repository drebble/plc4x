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
package org.apache.plc4x.java.openprotocol.readwrite;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum OpenProtocolRevision {
  Revision1((long) 1, (short) 3),
  Revision2((long) 2, (short) 4),
  Revision3((long) 3, (short) 7),
  Revision4((long) 4, (short) 9),
  Revision5((long) 5, (short) 11),
  Revision6((long) 6, (short) 16);
  private static final Map<Long, OpenProtocolRevision> map;

  static {
    map = new HashMap<>();
    for (OpenProtocolRevision value : OpenProtocolRevision.values()) {
      map.put((long) value.getValue(), value);
    }
  }

  private long value;
  private short numCommunicationStartAcknowledgeBlocks;

  OpenProtocolRevision(long value, short numCommunicationStartAcknowledgeBlocks) {
    this.value = value;
    this.numCommunicationStartAcknowledgeBlocks = numCommunicationStartAcknowledgeBlocks;
  }

  public long getValue() {
    return value;
  }

  public short getNumCommunicationStartAcknowledgeBlocks() {
    return numCommunicationStartAcknowledgeBlocks;
  }

  public static OpenProtocolRevision firstEnumForFieldNumCommunicationStartAcknowledgeBlocks(
      short fieldValue) {
    for (OpenProtocolRevision _val : OpenProtocolRevision.values()) {
      if (_val.getNumCommunicationStartAcknowledgeBlocks() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<OpenProtocolRevision> enumsForFieldNumCommunicationStartAcknowledgeBlocks(
      short fieldValue) {
    List<OpenProtocolRevision> _values = new ArrayList();
    for (OpenProtocolRevision _val : OpenProtocolRevision.values()) {
      if (_val.getNumCommunicationStartAcknowledgeBlocks() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public static OpenProtocolRevision enumForValue(long value) {
    return map.get(value);
  }

  public static Boolean isDefined(long value) {
    return map.containsKey(value);
  }
}
