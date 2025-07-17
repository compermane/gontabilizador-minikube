# Pegar as imagens
eval $(minikube docker-env)

docker build -t gontabilizador-app:latest .
docker build -t mysql:latest ./db
docker build -t custom-nginx:lastest ./nginx
docker build -t python-listener:lasters ./python

# Iniciar o minikube
minikube start

helm install gontabilizador ./gontabilizador