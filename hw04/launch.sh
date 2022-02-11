#!/bin/zsh

# Create namespace
kubectl apply -f manifest/00-namespace.yaml

# Deploy application
helm install otus-hw04 ./simple-app/.helm -n otus-hw04-ns

# Initialize istio
kubectl apply -f istio/ -n otus-hw04-ns