#!/bin/bash

docker build -t go-containerized:latest . && docker run  -p 9000:9000 go-containerized:latest

