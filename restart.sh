#! /bin/bash
git pull
go build -o scanengin worker.go
ps -ef | grep ./scanengin | grep -v grep | awk '{print $2}' | xargs kill -9
cd /zrtx/cy/scan-engin && ./scanengin >> scan.log &
