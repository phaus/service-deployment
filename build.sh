#!/bin/bash

docker build --pull -t phaus/service-deployment .

docker images | grep "phaus/service-deployment"