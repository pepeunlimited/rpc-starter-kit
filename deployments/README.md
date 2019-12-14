# `/deployments`

IaaS, PaaS, system and container orchestration deployment configurations and templates (docker-compose, kubernetes/helm, mesos, terraform, bosh).

# NATS (MessageBus)

kubectl apply -f https://bitbucket.org/siimooo/k8-defaults/raw/62b28f29a824804846d0a36ad0ba297fd299ecc1/nats-deployment.yaml,https://bitbucket.org/siimooo/k8-defaults/raw/62b28f29a824804846d0a36ad0ba297fd299ecc1/nats-service.yaml

https://stackoverflow.com/questions/51874503/kubernetes-ingress-network-deny-some-paths

# Enabled JWT Auth

kubernetes.io/ingress.class: nginx
nginx.ingress.kubernetes.io/auth-url: http://nginx-jwt.ingress-nginx.svc.cluster.local:8000/verify
nginx.ingress.kubernetes.io/auth-response-headers: X-JWT-Username