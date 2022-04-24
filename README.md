
## 随便写写看。

----

## MySQL

```shell
docker pull mysql/mysql-server:latest
docker run --name mysql -p 3306:3306 -d mysql/mysql-server
docker exec -it xx mysql -uroot -p # docker logs xx


ALTER USER 'root'@'localhost' IDENTIFIED BY 'password'; 
CREATE USER 'root'@'%' IDENTIFIED BY 'root'; GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION;
ALTER USER 'root'@'%' IDENTIFIED BY '[newpassword]';

docker restart xx
```

## Redis

```shell
docker pull redis
docker run --name redis -p 16379:6379 -e -d redis
```