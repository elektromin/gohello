#!/bin/bash

gcloud config set project autoscaling-gke-test

gcloud config set compute/zone europe-west4-a

gcloud container clusters resize k0 --size 0

## in cloud shell
kubectl create secret generic environment --from-file=NEARBYSTOPS_KEY

go build -o gohello .