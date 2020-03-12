docker-build:
	docker build -t eu.gcr.io/home-kubernetes-261703/fizzbuzz .

docker-push:
	docker push eu.gcr.io/home-kubernetes-261703/fizzbuzz:latest

helm-install:
	helm install fizzbuzz ./kubernetes/fizzbuzz

helm-upgrade:
	helm upgrade fizzbuzz ./kubernetes/fizzbuzz

install: docker-build docker-push helm-install
upgrade: docker-build docker-push helm-upgrade
