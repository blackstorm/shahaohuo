#!/bin/bash
replace() {
  # image version
  sed -i "s/:latest/:$2/g" $1
}

echo "template = $1 Docker image version = $2"
replace $1 $2;