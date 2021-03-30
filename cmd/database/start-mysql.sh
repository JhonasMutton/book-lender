#!/bin/bash
echo "Starting MySql!"
docker run -it -d --name=mysql \
  -p 3306:3306 \
  -e "MYSQL_ROOT_PASSWORD=admin" \
  mysql:8.0.23