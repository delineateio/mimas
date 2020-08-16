#!/usr/bin/env bash
set -e

###################################################################
# Script Name	: setup.sh
# Description	: Sets up the local development environment
# Author       	: Jonathan Fenwick
# Email         : jonathan.fenwick@delineate.io
###################################################################

# Isolates dev environment
rm -rf ./venv
python3 -m venv ./venv
source ./venv/bin/activate
./venv/bin/python3 -m pip install --upgrade pip
pip3 install -r requirements.txt

# Setup pre-commit
pre-commit install
