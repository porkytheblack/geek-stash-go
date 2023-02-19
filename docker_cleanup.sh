#!/bin/bash

if [ $(docker ps -aq) ]; then
    echo "Stopping and removing Docker containers"
    docker stop $(docker ps -aq) && docker rm $(docker ps -aq)
else
    echo "No Docker containers to stop and remove"
fi

if [ $(docker volume ls -q) ]; then
    echo "Removing Docker volumes"
    docker volume rm $(docker volume ls -q)
else
    echo "No Docker volumes to remove"
fi
