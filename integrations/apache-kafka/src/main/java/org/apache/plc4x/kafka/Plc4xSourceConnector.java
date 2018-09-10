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
package org.apache.plc4x.kafka;

import org.apache.kafka.common.config.ConfigDef;
import org.apache.kafka.connect.connector.Task;
import org.apache.kafka.connect.source.SourceConnector;
import org.apache.plc4x.kafka.source.Plc4xSourceTask;
import org.apache.plc4x.kafka.util.VersionUtil;

import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Plc4xSourceConnector extends SourceConnector {
    public static final String TOPIC_CONFIG = "topic";
    private static final String TOPIC_DOC = "Kafka topic to publish to";

    public static final String URL_CONFIG = "url";
    private static final String URL_DOC = "Connection string used by PLC4X to connect to the PLC";

    public static final String QUERY_CONFIG = "query";
    private static final String QUERY_DOC = "Field query to be sent to the PLC";

    private static final ConfigDef CONFIG_DEF = new ConfigDef()
        .define(TOPIC_CONFIG, ConfigDef.Type.STRING, ConfigDef.Importance.HIGH, TOPIC_DOC)
        .define(URL_CONFIG, ConfigDef.Type.STRING, ConfigDef.Importance.HIGH, URL_DOC)
        .define(QUERY_CONFIG, ConfigDef.Type.STRING, ConfigDef.Importance.HIGH, QUERY_DOC);

    private String topic;
    private String url;
    private String query;

    @Override
    public Class<? extends Task> taskClass() {
        return Plc4xSourceTask.class;
    }

    @Override
    public List<Map<String, String>> taskConfigs(int maxTasks) {
        Map<String, String> taskConfig = new HashMap<>();
        taskConfig.put(TOPIC_CONFIG, topic);
        taskConfig.put(URL_CONFIG, url);
        taskConfig.put(QUERY_CONFIG, query);

        // Only one task will be created; ignoring maxTasks for now
        return Collections.singletonList(taskConfig);
    }

    @Override
    public void start(Map<String, String> props) {
        topic = props.get(TOPIC_CONFIG);
        url = props.get(URL_CONFIG);
        query = props.get(QUERY_CONFIG);
    }

    @Override
    public void stop() {}

    @Override
    public ConfigDef config() {
        return CONFIG_DEF;
    }

    @Override
    public String version() {
        return VersionUtil.getVersion();
    }

}
