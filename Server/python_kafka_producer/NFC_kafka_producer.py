from kafka import KafkaProducer
from json import dumps
import time

producer = KafkaProducer(acks=0, compression_type='gzip', bootstrap_servers=['59.27.40.116:9092'], value_serializer=lambda x: dumps(x).encode('utf-8'))

start = time.time()

for i in range(10000):
    data = {'str' : 'result'+str(i)}
    producer.send('user-check-in', value=data)
    producer.flush()
print("elapsed : ", time.time() - start)