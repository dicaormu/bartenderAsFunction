#!/bin/bash

P_FOLDER=bin

P_FILE_TO_BUILD=$1

if [[ -z "${P_FILE_TO_BUILD}" ]]; then
  echo "You must give a lambda name to build"
  exit
fi

export GOOS=linux

echo "building $P_FILE_TO_BUILD...."
target=${P_FOLDER}
go build -o  ${P_FOLDER}/$P_FILE_TO_BUILD bartenderAsFunction/functions/$P_FILE_TO_BUILD;
echo ".....built "


