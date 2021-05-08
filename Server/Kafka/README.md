# bitnami Kafka Setting

* client : /opt/bitnami/kafka/bin
* setting : ./kafka-topics.sh --create --bootstrap-server localhost:9092 --topic user-check-in
* docker-compose.yml setting : KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://"서버 ip":9092