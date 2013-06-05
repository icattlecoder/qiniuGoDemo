#!/bin/bash
if [ -e ./qrsb.keys ]
then
	make clean
fi
while true 
do
	go run check.go 
	if [ $? -eq 0 ]
	then
		make clean
		./qrsb_nodownload ~/fixbuges/qrsb
	else
		echo "find the bug\n"
		break
	fi
done
