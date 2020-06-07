#!/bin/bash

if [ ! -d "gen" ]; then
  mkdir gen
fi

swagger generate server -t gen --exclude-main -A EmployeeManagementSystem
swagger generate client -t gen -A EmployeeManagementSystem
