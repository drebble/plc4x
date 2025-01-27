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

public abstract class OpenProtocolMessage implements Message {

  // Abstract accessors for discriminator values.
  public abstract Mid getMid();

  // Constant values.
  public static final Short END = 0x00;

  // Properties.
  protected final OpenProtocolRevision revision;
  protected final short noAckFlag;
  protected final int stationId;
  protected final int spindleId;
  protected final int sequenceNumber;
  protected final short numberOfMessageParts;
  protected final short messagePartNumber;

  public OpenProtocolMessage(
      OpenProtocolRevision revision,
      short noAckFlag,
      int stationId,
      int spindleId,
      int sequenceNumber,
      short numberOfMessageParts,
      short messagePartNumber) {
    super();
    this.revision = revision;
    this.noAckFlag = noAckFlag;
    this.stationId = stationId;
    this.spindleId = spindleId;
    this.sequenceNumber = sequenceNumber;
    this.numberOfMessageParts = numberOfMessageParts;
    this.messagePartNumber = messagePartNumber;
  }

  public OpenProtocolRevision getRevision() {
    return revision;
  }

  public short getNoAckFlag() {
    return noAckFlag;
  }

  public int getStationId() {
    return stationId;
  }

  public int getSpindleId() {
    return spindleId;
  }

  public int getSequenceNumber() {
    return sequenceNumber;
  }

  public short getNumberOfMessageParts() {
    return numberOfMessageParts;
  }

  public short getMessagePartNumber() {
    return messagePartNumber;
  }

  public short getEnd() {
    return END;
  }

  protected abstract void serializeOpenProtocolMessageChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("OpenProtocolMessage");

    // Implicit Field (length) (Used for parsing, but its value is not stored as it's implicitly
    // given by the objects content)
    long length = (long) (getLengthInBytes());
    writeImplicitField("length", length, writeUnsignedLong(writeBuffer, 32));

    // Discriminator Field (mid) (Used as input to a switch field)
    writeDiscriminatorEnumField(
        "mid",
        "Mid",
        getMid(),
        new DataWriterEnumDefault<>(Mid::getValue, Mid::name, writeUnsignedLong(writeBuffer, 32)));

    // Simple Field (revision)
    writeSimpleEnumField(
        "revision",
        "OpenProtocolRevision",
        revision,
        new DataWriterEnumDefault<>(
            OpenProtocolRevision::getValue,
            OpenProtocolRevision::name,
            writeUnsignedLong(writeBuffer, 24)));

    // Simple Field (noAckFlag)
    writeSimpleField("noAckFlag", noAckFlag, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (stationId)
    writeSimpleField(
        "stationId",
        stationId,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("AsciiUint"));

    // Simple Field (spindleId)
    writeSimpleField(
        "spindleId",
        spindleId,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("AsciiUint"));

    // Simple Field (sequenceNumber)
    writeSimpleField(
        "sequenceNumber",
        sequenceNumber,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("AsciiUint"));

    // Simple Field (numberOfMessageParts)
    writeSimpleField(
        "numberOfMessageParts",
        numberOfMessageParts,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithEncoding("AsciiUint"));

    // Simple Field (messagePartNumber)
    writeSimpleField(
        "messagePartNumber",
        messagePartNumber,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithEncoding("AsciiUint"));

    // Switch field (Serialize the sub-type)
    serializeOpenProtocolMessageChild(writeBuffer);

    // Const Field (end)
    writeConstField("end", END, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("OpenProtocolMessage");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    OpenProtocolMessage _value = this;

    // Implicit Field (length)
    lengthInBits += 32;

    // Discriminator Field (mid)
    lengthInBits += 32;

    // Simple field (revision)
    lengthInBits += 24;

    // Simple field (noAckFlag)
    lengthInBits += 8;

    // Simple field (stationId)
    lengthInBits += 16;

    // Simple field (spindleId)
    lengthInBits += 16;

    // Simple field (sequenceNumber)
    lengthInBits += 16;

    // Simple field (numberOfMessageParts)
    lengthInBits += 8;

    // Simple field (messagePartNumber)
    lengthInBits += 8;

    // Length of sub-type elements will be added by sub-type...

    // Const Field (end)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static OpenProtocolMessage staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static OpenProtocolMessage staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("OpenProtocolMessage");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    long length =
        readImplicitField(
            "length", readUnsignedLong(readBuffer, 32), WithOption.WithEncoding("AsciiUint"));

    Mid mid =
        readDiscriminatorField(
            "mid",
            new DataReaderEnumDefault<>(Mid::enumForValue, readUnsignedLong(readBuffer, 32)));

    OpenProtocolRevision revision =
        readEnumField(
            "revision",
            "OpenProtocolRevision",
            new DataReaderEnumDefault<>(
                OpenProtocolRevision::enumForValue, readUnsignedLong(readBuffer, 24)));

    short noAckFlag = readSimpleField("noAckFlag", readUnsignedShort(readBuffer, 8));

    int stationId =
        readSimpleField(
            "stationId", readUnsignedInt(readBuffer, 16), WithOption.WithEncoding("AsciiUint"));

    int spindleId =
        readSimpleField(
            "spindleId", readUnsignedInt(readBuffer, 16), WithOption.WithEncoding("AsciiUint"));

    int sequenceNumber =
        readSimpleField(
            "sequenceNumber",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithEncoding("AsciiUint"));

    short numberOfMessageParts =
        readSimpleField(
            "numberOfMessageParts",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithEncoding("AsciiUint"));

    short messagePartNumber =
        readSimpleField(
            "messagePartNumber",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithEncoding("AsciiUint"));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    OpenProtocolMessageBuilder builder = null;
    if (EvaluationHelper.equals(mid, Mid.ApplicationCommunicationStart)) {
      builder = OpenProtocolMessageApplicationCommunicationStart.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(mid, Mid.ApplicationCommunicationStartAcknowledge)) {
      builder =
          OpenProtocolMessageApplicationCommunicationStartAcknowledge.staticParseBuilder(
              readBuffer, revision);
    } else if (EvaluationHelper.equals(mid, Mid.ApplicationCommunicationStop)) {
      builder = OpenProtocolMessageApplicationCommunicationStop.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(mid, Mid.ApplicationCommandError)) {
      builder = OpenProtocolMessageApplicationCommandError.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(mid, Mid.ApplicationCommandAccepted)) {
      builder = OpenProtocolMessageApplicationCommandAccepted.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(mid, Mid.ApplicationGenericDataRequest)) {
      builder = OpenProtocolMessageApplicationGenericDataRequest.staticParseBuilder(readBuffer);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type" + " parameters [" + "mid=" + mid + "]");
    }

    short end = readConstField("end", readUnsignedShort(readBuffer, 8), OpenProtocolMessage.END);

    readBuffer.closeContext("OpenProtocolMessage");
    // Create the instance
    OpenProtocolMessage _openProtocolMessage =
        builder.build(
            revision,
            noAckFlag,
            stationId,
            spindleId,
            sequenceNumber,
            numberOfMessageParts,
            messagePartNumber);
    return _openProtocolMessage;
  }

  public static interface OpenProtocolMessageBuilder {
    OpenProtocolMessage build(
        OpenProtocolRevision revision,
        short noAckFlag,
        int stationId,
        int spindleId,
        int sequenceNumber,
        short numberOfMessageParts,
        short messagePartNumber);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof OpenProtocolMessage)) {
      return false;
    }
    OpenProtocolMessage that = (OpenProtocolMessage) o;
    return (getRevision() == that.getRevision())
        && (getNoAckFlag() == that.getNoAckFlag())
        && (getStationId() == that.getStationId())
        && (getSpindleId() == that.getSpindleId())
        && (getSequenceNumber() == that.getSequenceNumber())
        && (getNumberOfMessageParts() == that.getNumberOfMessageParts())
        && (getMessagePartNumber() == that.getMessagePartNumber())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getRevision(),
        getNoAckFlag(),
        getStationId(),
        getSpindleId(),
        getSequenceNumber(),
        getNumberOfMessageParts(),
        getMessagePartNumber());
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
