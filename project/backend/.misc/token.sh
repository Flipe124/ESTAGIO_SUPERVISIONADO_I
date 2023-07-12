#!/bin/bash
USERNAME=${1:-'rhuanpk'}
PASSWORD=${2:-'password'}
curl \
	-fsSL \
	localhost:9999/api/v0/auth \
	--request POST \
	--data "{\"username\": \"$USERNAME\", \"password\": \"$PASSWORD\"}" \
	| sed -nE 's/^.*":"(.*)"}$/\1/p'
