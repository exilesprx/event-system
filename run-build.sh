#!/bin/bash

BRANCH=$(git rev-parse --abbrev-ref HEAD)

echo -e "\033[35m Updating app image  exilesprx/golang:latest \033[0m"
DOCKER_BUILDKIT=1 docker build -f .docker/Dockerfile -t exilesprx/golang:latest .

echo -e "\033[32m Done! \033[0m"