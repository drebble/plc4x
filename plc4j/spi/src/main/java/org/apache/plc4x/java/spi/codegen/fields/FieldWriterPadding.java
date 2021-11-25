/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.spi.codegen.fields;

import org.apache.commons.lang3.NotImplementedException;
import org.apache.plc4x.java.spi.codegen.FieldCommons;
import org.apache.plc4x.java.spi.codegen.io.DataWriter;
import org.apache.plc4x.java.spi.generation.SerializationException;
import org.apache.plc4x.java.spi.generation.WithReaderWriterArgs;
import org.apache.plc4x.java.spi.generation.WithWriterArgs;

public class FieldWriterPadding<T> implements FieldCommons {

    public void writePaddingField(String logicalName, int timesPadding, T value, DataWriter<T> dataWriter, WithWriterArgs... writerArgs) throws SerializationException {
        switchSerializeByteOrderIfNecessary(() -> {
            dataWriter.pushContext(logicalName, WithReaderWriterArgs.WithRenderAsList(true));
            for (int i = 0; i < timesPadding; i++) {
                dataWriter.write("value", value, writerArgs);
            }
            dataWriter.popContext(logicalName, WithReaderWriterArgs.WithRenderAsList(true));
        }, dataWriter, extractByteOder(writerArgs).orElse(null));
    }

}
