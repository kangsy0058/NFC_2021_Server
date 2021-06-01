from os import fpathconf
from random import lognormvariate
from kafka import KafkaConsumer
from json import loads, load
import pymysql

#Read init file
with open("Consumer-init.json","r") as init_file:
    init_val = load(init_file)

# init Consumer object
consumer = KafkaConsumer(
    init_val["kafka"]["topic"],
    bootstrap_servers= init_val["kafka"]["bootstrap_servers"],
    auto_offset_reset= init_val["kafka"]["auto_offset_reset"],
    enable_auto_commit= init_val["kafka"]["enable_auto_commit"],
    #group_id= init_val["kafka"]["group_id"],
    value_deserializer=lambda x: loads(x.decode('utf-8')),
)

#init mariadb
conn = pymysql.connect(host=init_val["mariadb"]["host"], user=init_val["mariadb"]["user"], password=init_val["mariadb"]["password"], db=init_val["mariadb"]["db"], charset=init_val["mariadb"]["charset"])
#make db cursor
cur = conn.cursor()


# Running consumer 
print('[begin] get consumer list')
for message in consumer:
    print("Topic: %s, Partition: %d, Offset: %d, Key: %s, Value: %s" % ( message.topic, message.partition, message.offset, message.key, message.value ))
    rec_val = message.value

    #get kiosk detail
    kiosk_detail_sql = f'SELECT * FROM kiosk_set ks WHERE kiosk_SN = "{rec_val["kiosk_SN"]}"'
    cur.execute(kiosk_detail_sql)
    # kiosk_SN, Group_code, detail_position, building_namem latitude, longitude
    _, G_code, detail_P, building_N, lat, lg = cur.fetchall()[0]

    log_send_sql = f'INSERT INTO user_log (wearable_SN, kiosk_SN, time, date, temp, Group_code, detail_position, building_name, latitude, longitude) VALUES ("{rec_val["wearable_SN"]}", "{rec_val["kiosk_SN"]}", "{rec_val["time"]}", "{rec_val["date"]}", "{rec_val["temp"]}", "{G_code}", "{detail_P}", "{building_N}", "{lat}", "{lg}")'
    print(log_send_sql)
    cur.execute("INSERT INTO user_log (wearable_SN, kiosk_SN, time, date, temp, Group_code, detail_position, building_name, latitude, longitude) VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", (rec_val["wearable_SN"], rec_val["kiosk_SN"], rec_val["time"], rec_val["date"], rec_val["temp"], G_code, detail_P, building_N, lat, lg))
    conn.commit()
    


print('[end] get consumer list')