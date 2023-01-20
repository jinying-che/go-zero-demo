# shorturl service
this demo consists of api and DB layer which is simple version removed the rpc module from the [offical demo](https://github.com/zeromicro/zero-doc/blob/main/doc/shorturl-en.md)
### Service Setup
refer to https://github.com/zeromicro/zero-doc/blob/main/doc/shorturl-en.md

#### MySQL Docker
set up MySQL Docker basing on https://hub.docker.com/_/mysql: 
```sh
# run mysql container
docker run -p 13306:3306 -e MYSQL_ROOT_PASSWORD=123 -d mysql:latest

# access mysql container
docker exec -it container_id bash

# create database gozero
# create table...
```

### Run 
```sh
# start api server
 go run ./api/shorturl.go -f api/etc/shorturl-api.yaml

# make a request via curl
curl -i http://localhost:8888/shorten\?url\=www.baidu.com
```

