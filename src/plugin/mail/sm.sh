#!/bin/sh
echo $1
echo $2
cat $! | sendmail $2