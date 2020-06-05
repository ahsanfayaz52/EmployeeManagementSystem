#!/bin/bash

if [ "$(docker ps -q -f name=employee-system-mysql-db)" ]; then
  docker rm -f employee-system-mysql-db
fi

if [ "$(docker ps -q -f name=employee-employee-mongo-db)" ]; then
  docker rm -f employee-system-mongo-db
fi
