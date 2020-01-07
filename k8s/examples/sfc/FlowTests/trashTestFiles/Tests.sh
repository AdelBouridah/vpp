#!/bin/bash
set -eu

while : 
do
	who | cut -d' ' -f1 >fic 
	sleep 300 
done &


echo "Linux SFC tests done."
