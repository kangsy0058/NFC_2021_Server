from kafka import KafkaProducer
from json import dumps
import time
import random

#kiosk SN 이미 가지고있다 가정 (값은 추후 변경가능)
kiosk_sn = ["KSN1111", "KSN1112", "KSN1113", "KSN1114", "KSN1115", "KSN1116", "KSN1117", "KSN1118", "KSN1119", "KSN1120"]
# wearable SN 받았다 가정 (값은 추후 변경가능)
wearable_sn = ["WSN1111", "WSN1112", "WSN1113", "WSN1114", "WSN1115", "WSN1116", "WSN1117", "WSN1118", "WSN1119", "WSN1120"]

# kiosk_SN 과 wearable_SN으로 랜덤한 로그를 생성한다.
def make_rand_log():

    #sn 결정위한 랜덤 넘버
    ki_rand = random.randrange(10)
    wr_rand = random.randrange(10)

    #log time = 현재시간, l_time 시간만 따로 저장, l_data 날짜만 따로 저장
    logtime = time.localtime()
    l_time = time.strftime('%I:%M:%S', logtime)
    l_date = time.strftime('%Y-%m-%d', logtime)

    temp = random.randint(36, 37)

    print(l_time)
    print(l_date)
    #--------------(중요) 전송할 데이터 포멧 이대로 전송해야함---------------------------------------------
    data = {'wearable_SN' : wearable_sn[wr_rand], 'kiosk_SN' : kiosk_sn[ki_rand], "time" : l_time, "date" : l_date, "temp": temp}

    return data


#카프카프로듀서 연결 bootstrap server만 kafka 서버 ip로 변경하여 사용
producer = KafkaProducer(acks=0, compression_type='gzip', bootstrap_servers=['59.27.40.116:9092'], value_serializer=lambda x: dumps(x).encode('utf-8'))

print("start send log")

for i in range(1000):
    
    data = make_rand_log()
    producer.send('user-check-in', value=data)
    producer.flush()


print("finish send log")