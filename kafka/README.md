## 操作命令

### sasl
producer
```shell
./kafka_client -command producer   -sasl -username admin -password yunjikeji  -host 10.0.3.162:9092,10.0.3.162:9093,10.0.3.162:9094
```
consumer
```shell
./kafka_client -command consumer   -sasl -username admin -password yunjikeji  -host 10.0.3.162:9092,10.0.3.162:9093,10.0.3.162:9094
```

### tls
producer
```shell
./kafka_client -command producer   -tls -cert /tmp/act2/client.cer.pem -key /tmp/act2/client.key.pem -ca /tmp/act2/server.cer.pem   -host 10.0.3.109:9092,10.0.3.109:9093,10.0.3.109:9094
```

consumer
```shell
./kafka_client -command consumer   -tls -cert /tmp/act2/client.cer.pem -key /tmp/act2/client.key.pem -ca /tmp/act2/server.cer.pem   -host 10.0.3.109:9092,10.0.3.109:9093,10.0.3.109:9094
```
