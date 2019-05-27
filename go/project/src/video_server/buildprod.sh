#! /bin/bash

cd ~/work/src/github.com/avenssi/video_server/api
env GOOS=linux GOARCH=amd64 go build -o ../bin/api