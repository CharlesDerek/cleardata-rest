#!/bin/bash

# Make go package (used dev.mrgi-portal.com as example domain).
go mod init dev.mrgi-portal.com

# Remove unused packages (if any exists).
go mody tidy

# String Conversions:
go get -u strconv