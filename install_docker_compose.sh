#!/bin/bash

# Check if Docker Compose is already installed
if ! [ -x "$(command -v docker-compose)" ]; then
  # Install Docker Compose
  echo "Installing Docker Compose..."
  sudo curl -L "https://github.com/docker/compose/releases/download/3.8/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
  sudo chmod +x /usr/local/bin/docker-compose
  echo "Docker Compose installation complete."
else
  echo "Docker Compose is already installed."
fi
