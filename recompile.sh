export GOPATH="$HOME/go"
eval $(minikube docker-env)
build/build-linux.sh
build/docker.sh fission latest

helm uninstall --namespace fission workflows
sleep 0.5 
helm install workflows --namespace fission charts/fission-workflows 
sleep 3
kubectl get pods -A | grep workflows