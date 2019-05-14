#coding=utf-8
from flask import render_template,jsonify,request
from flask import Flask
#from config import hosts, password, username
from threading import Timer
import time
import random
import pymysql
import re
import os
import threading
app = Flask(__name__)
# app.debug = True

conn = pymysql.connect(
    host='172.19.129.11',
    user='root',
    password='123456',
    db='ping',
    charset='utf8'
)

hosts =[]
# conn1 = pymysql.connect(
#     host='172.19.129.11',
#     user='root',
#     password='finup2019',
#     db='ping',
#     charset='utf8'
# )

cur = conn.cursor()
sql = 'select * from host'
cur.execute(sql)
hosts1 = cur.fetchall()
hosts = [h[0] for h in hosts1]
print("hosts",hosts)

@app.route('/')
def index():
    hosts_ = []
    # cur = conn.cursor()
    # sql = 'select * from host'
    # cur.execute(sql)
    # hosts = cur.fetchall()
    # print(hosts)
    for i,host in enumerate(hosts):
        index = str(i)
        name = 'host'+index.rjust(2, '0')
        hosts_.append(name)
   # cur.close()
  #  conn.close()
    return render_template("index.html", **locals())


@app.route('/update_mesh/')
def update_mesh():
    # 线程数视具体情况而定
    mesh_data = {}
    avg_time100 = 0
    cur = conn.cursor()
    # sql = 'select * from host'
    # cur.execute(sql)
    # hosts1 = cur.fetchall()
    # hosts = [h[0] for h in hosts1]
    for i, host in enumerate(hosts):
        s_ip = hosts[i]

        for j in range(len(hosts)):
            r_ip = hosts[j]
            key = 'host' + str(i).rjust(2, '0') + '-' + 'host' + str(j).rjust(2, '0')
            sql1 = "select rttavg from valu where tss > UNIX_TIMESTAMP()-120 and tss < UNIX_TIMESTAMP()-60 and src='%s' and dst = '%s'" % (s_ip, r_ip)
            print(sql1)
            conn.ping(True)
            cur.execute(sql1)
            avg_time = cur.fetchall()
            print("avg_time", avg_time)
            if avg_time:
               avg_time100 = avg_time[0][0]
               print("avg_time[0])",avg_time[0])
               print("avg_time[0][0])",avg_time[0][0])
               value = int(float(avg_time100) * 100)
               print("value", value)
               mesh_data[key] = [value, "aprilmadaha"]
               print("mesh_data", mesh_data)
    cur.close()
    conn.close()
    return jsonify(mesh_data)



if __name__ == '__main__':
    #start_thread()
    app.run(host='0.0.0.0',port=9000)
