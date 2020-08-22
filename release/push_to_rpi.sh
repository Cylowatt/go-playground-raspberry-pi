#!/usr/bin/env bash

cd ..

env GOOS=linux GOARCH=arm go build -o server src/main.go

zip -r server.zip ./* -x src/ src/* release/ release/*

scp ./server.zip pi@raspberrypi:/home/pi/server.zip
