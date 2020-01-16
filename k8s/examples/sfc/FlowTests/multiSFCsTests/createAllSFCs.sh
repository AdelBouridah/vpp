#!/bin/bash
set -eu

kubectl apply -f sfc1.yaml

kubectl apply -f sfc2.yaml
kubectl apply -f sfc3.yaml
kubectl apply -f sfc4.yaml
kubectl apply -f sfc5.yaml
