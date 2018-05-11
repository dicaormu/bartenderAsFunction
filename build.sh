#!/bin/bash

P_FOLDER=bin

export GOOS=linux

for folder in functions/*;
do
package_name=`basename $folder | sed 's/{//' | sed 's/}//'`
echo $package_name
target=${folder}
#function=`basename $target | sed 's/{//' | sed 's/}//'`
go build -o  ${P_FOLDER}/$package_name bartenderAsFunction/$target;
done
