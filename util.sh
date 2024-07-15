#!/bin/bash

case "$1" in
"build")
	cd server/ && make && cd -
	cd web/ && make
	echo "Successfully built web client and server"
	;;
"exec")
	./server/bin/udp-server &
	cd web && ./bin/udp-client &
	cd -
	;;
"clean")
	cd server && make clean && cd -
	cd web && make clean
	;;
"stop")
	ps au | grep 'udp-server\|udp-client' | grep -v grep | awk '{print $2}' | xargs kill
	;;
*)
	echo "None or incorrect argument provided. Allowed arguments are 'build', 'exec', 'stop' or 'clean'"
	;;
esac
