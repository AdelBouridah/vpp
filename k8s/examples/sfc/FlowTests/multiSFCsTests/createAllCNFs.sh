#!/bin/bash
set -eu

kubectl apply -f cnfs-sfc1.yaml

kubectl apply -f cnfs-sfc2.yaml
kubectl apply -f cnfs-sfc3.yaml
kubectl apply -f cnfs-sfc4.yaml
kubectl apply -f cnfs-sfc5.yaml
