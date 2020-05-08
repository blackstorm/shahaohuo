#!/bin/bash

echo "template = $1 image version = $2"
sed -i.bu "s/latest/$2/g" $1
