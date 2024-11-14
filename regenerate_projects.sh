#!/bin/bash

cd ~/gbdubs/grady.dev/projectdata

hugo

rm -R ../project/
mkdir ../project
mkdir ../project/img
cp -R public/project/* ../project
cp -R public/img/* ../project/img

rm -R public