#!/bin/bash

user=$1
host=$2

# On Server AllowTcpForwarding needs to be set to yes: https://unix.stackexchange.com/a/58756
ssh -N -L 0.0.0.0:5433:localhost:5432 $user@$host