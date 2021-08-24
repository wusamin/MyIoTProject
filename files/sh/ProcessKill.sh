#!/bin/bash

sleep $2s

process=`ps aux | grep $1 | grep -v "grep" | grep -v "ProcessKill" | head -1`

if [ -z  "${process}" ]; then
  echo "$1 processes were not available."
  exit 0
fi

arrays=(${process})

kill ${arrays[1]}
echo "'killed ${process}'"
