#!/bin/bash

GOOS=linux go build -ldflags="-s -w" 

upx --brute configtool
