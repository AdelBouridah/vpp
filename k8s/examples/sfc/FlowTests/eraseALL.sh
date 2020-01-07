#!/bin/bash
set -eu

kubectl delete --all servicefunctionchain

kubectl delete --all replicaset
