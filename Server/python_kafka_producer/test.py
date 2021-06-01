from json import loads, load
import pymysql

with open("Consumer-init.json","r") as init_file:
    init_val = load(init_file)

conn = pymysql.connect(host=init_val["mariadb"]["host"], user=init_val["mariadb"]["user"], password=init_val["mariadb"]["password"], db=init_val["mariadb"]["db"], charset=init_val["mariadb"]["charset"])