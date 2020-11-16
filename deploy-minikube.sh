#!/bin/sh
set -o errexit -o nounset
cd `dirname $0`

kubectl config use-context minikube

eval $(minikube docker-env --shell=bash)

docker-compose build go-react-app
kubectl delete deployment go-react-app || true
kubectl apply -f .minikube
kubectl rollout status deployment/go-react-app
kubectl get deployment go-react-app
echo "Deployed on $(minikube service go-react-app --url)"