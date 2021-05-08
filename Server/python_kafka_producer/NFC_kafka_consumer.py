from kafka import KafkaConsumer
from json import loads
consumer = KafkaConsumer(
    'user-check-in',
    bootstrap_servers=['59.27.40.116:9092'],
    auto_offset_reset='latest',
    enable_auto_commit=True,
    value_deserializer=lambda x: loads(x.decode('utf-8')),
)

print('[begin] get consumer list')
for message in consumer:
    print("Topic: %s, Partition: %d, Offset: %d, Key: %s, Value: %s" % ( message.topic, message.partition, message.offset, message.key, message.value ))

print('[end] get consumer list')