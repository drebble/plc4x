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

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class ApplicationCommunicationStartAcknowledgeBlockSequenceNumberSupport
    extends ApplicationCommunicationStartAcknowledgeBlock implements Message {

  // Accessors for discriminator values.
  public Integer getBlockType() {
    return (int) 12;
  }

  // Properties.
  protected final boolean sequenceNumberSupport;
  // Reserved Fields
  private Short reservedField0;

  public ApplicationCommunicationStartAcknowledgeBlockSequenceNumberSupport(
      boolean sequenceNumberSupport) {
    super();
    this.sequenceNumberSupport = sequenceNumberSupport;
  }

  public boolean getSequenceNumberSupport() {
    return sequenceNumberSupport;
  }

  @Override
  protected void serializeApplicationCommunicationStartAcknowledgeBlockChild(
      WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ApplicationCommunicationStartAcknowledgeBlockSequenceNumberSupport");

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (short) 0x00,
        writeUnsignedShort(writeBuffer, 7));

    // Simple Field (sequenceNumberSupport)
    writeSimpleField("sequenceNumberSupport", sequenceNumberSupport, writeBoolean(writeBuffer));

    writeBuffer.popContext("ApplicationCommunicationStartAcknowledgeBlockSequenceNumberSupport");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ApplicationCommunicationStartAcknowledgeBlockSequenceNumberSupport _value = this;

    // Reserved Field (reserved)
    lengthInBits += 7;

    // Simple field (sequenceNumberSupport)
    lengthInBits += 1;

    return lengthInBits;
  }

  public static ApplicationCommunicationStartAcknowledgeBlockSequenceNumberSupportBuilder
      staticParseBuilder(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("ApplicationCommunicationStartAcknowledgeBlockSequenceNumberSupport");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    Short reservedField0 =
        readReservedField("reserved", readUnsignedShort(readBuffer, 7), (short) 0x00);

    boolean sequenceNumberSupport =
        readSimpleField("sequenceNumberSupport", readBoolean(readBuffer));

    readBuffer.closeContext("ApplicationCommunicationStartAcknowledgeBlockSequenceNumberSupport");
    // Create the instance
    return new ApplicationCommunicationStartAcknowledgeBlockSequenceNumberSupportBuilder(
        sequenceNumberSupport, reservedField0);
  }

  public static class ApplicationCommunicationStartAcknowledgeBlockSequenceNumberSupportBuilder
      implements ApplicationCommunicationStartAcknowledgeBlock
          .ApplicationCommunicationStartAcknowledgeBlockBuilder {
    private final boolean sequenceNumberSupport;
    private final Short reservedField0;

    public ApplicationCommunicationStartAcknowledgeBlockSequenceNumberSupportBuilder(
        boolean sequenceNumberSupport, Short reservedField0) {
      this.sequenceNumberSupport = sequenceNumberSupport;
      this.reservedField0 = reservedField0;
    }

    public ApplicationCommunicationStartAcknowledgeBlockSequenceNumberSupport build() {
      ApplicationCommunicationStartAcknowledgeBlockSequenceNumberSupport
          applicationCommunicationStartAcknowledgeBlockSequenceNumberSupport =
              new ApplicationCommunicationStartAcknowledgeBlockSequenceNumberSupport(
                  sequenceNumberSupport);
      applicationCommunicationStartAcknowledgeBlockSequenceNumberSupport.reservedField0 =
          reservedField0;
      return applicationCommunicationStartAcknowledgeBlockSequenceNumberSupport;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ApplicationCommunicationStartAcknowledgeBlockSequenceNumberSupport)) {
      return false;
    }
    ApplicationCommunicationStartAcknowledgeBlockSequenceNumberSupport that =
        (ApplicationCommunicationStartAcknowledgeBlockSequenceNumberSupport) o;
    return (getSequenceNumberSupport() == that.getSequenceNumberSupport())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getSequenceNumberSupport());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
