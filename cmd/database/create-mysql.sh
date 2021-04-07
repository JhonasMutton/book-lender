#!/bin/bash
echo "Creating MySql container!"
docker create -it --name=mysql \
  -p 3306:3306 \
  -e "MYSQL_ROOT_PASSWORD=admin" \
  mysql:8.0.23