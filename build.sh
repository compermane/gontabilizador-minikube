#!/bin/bash

# Iniciar o minikube
minikube start

# Iniciar o ingress
minikube addons enable ingress

# Pegar as imagens
eval $(minikube docker-env)

docker build -t gontabilizador-app:latest .
docker build -t custom-mysql:latest ./db
docker build -t custom-nginx:latest ./nginx
docker build -t python-listener:latest ./python

helm install gontabilizador ./gontabilizador