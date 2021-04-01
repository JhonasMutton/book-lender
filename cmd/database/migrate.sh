#!/bin/bash
export $(cat ../../.env)

action=$1
if [ -z "$1" ]
  then
    action='up'
fi

docker run -it -v /$(pwd)/../../migrations:/migrations \
--network host migrate/migrate \
-path=/migrations/ \
-database "mysql://$DB_USER:$DB_PASSWORD@tcp($DB_HOST:$DB_PORT)/$DB_NAME?multiStatements=true" $action