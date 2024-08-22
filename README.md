Pingmesh: A Large-Scale System for Data Center Network Latency Measurement and Analysis
Pingmesh：用于数据中心网络延迟测量和分析的大规模系统
====

最终效果图
-
![Image text](https://github.com/aprilmadaha/pingmesh/blob/master/pingmesh-image/effect%20picture.png)<br>

![Image text](https://github.com/aprilmadaha/pingmesh/blob/master/pingmesh-image/pingmesh-architecture.png)<br>

----后端配置（客户端和服务端时间必须准确）----<br>
-
一、服务端（Centos7）<br>
-
1.配置步骤<br>
yum install -y golang<br>

sudo yum install mariadb-servergo

sudo systemctl start mariadb<br>
sudo systemctl enable mariadb

sudo systemctl status mariadb


mysql_secure_installation //初始化mariadb paswd:123456<br>
sudo systemctl stop firewalld.service<br>
setenforce 0<br>

2.导入数据库表(pingmesh.sql)

3.编译（pingmesh-s-v1.1-GetResult.go/pingmesh-s-v1.1-GetHostIp.go）<br>
go build pingmesh-s-v1.1-GetResult.go<br>
go build pingmesh-s-v1.1-GetHostIp.go<br>

4.运行<br>
nohup ./pingmesh-s-v1.1-GetResult > output.log 2>&1 &<br>
nohup ./pingmesh-s-v1.1-GetHostIp > output.log 2>&1 &<br>

二、客户端<br>
-
1.基础安装<br>
yum install epel-release -y<br>
yum install fping -y<br>

2.编译后端（pingmesh-c-v1.1.go）<br>
go build pingmesh-c-v1.1.go<br>

3.运行<br>
nohup ./pingmesh-c-v1.1 > output.log 2>&1 &<br>

----前端配置----<br>
-
修改参数pingmesh.py<br>
conn = pymysql.connect(<br>
    host='172.19.129.11',<br>
    user='root',<br>
    password='123456',<br>
    db='ping',<br>
    charset='utf8'<br>
)<br>
部分<br>


1.运行项目<br>
python pingmesh.py<br>
后台运行：nohup  python pingmesh.py > pingmeshpy.log 2>&1 &

2.结果显示<br>
访问http://127.0.0.1:9000<br>

3.报错

ImportError: No module named flask
yum install python-pip -y<br>
pip install flask

ImportError: No module named pymysql<br>
 pip install pymysql

