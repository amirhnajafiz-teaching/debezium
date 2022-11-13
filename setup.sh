#!/bin/bash

# make an http call to send debezium configs to debezium server.
curl -i -X POST -H "Accept:application/json" \
 -H "Content-Type:application/json" \
  localhost:8083/connectors/ -d @debezium.json


# check the results
curl -i -X GET -H "Accept:application/json" localhost:8083/connectors/inventory-connector
