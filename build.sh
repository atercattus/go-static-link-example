#!/bin/bash

rm -f libsmth.{o,a} test
g++ -Wall -c -o libsmth.o smth.cpp && ar rcs libsmth.a libsmth.o
go build --ldflags '-extldflags "-static"' -o test main.go && rm -f libsmth.{o,a}
