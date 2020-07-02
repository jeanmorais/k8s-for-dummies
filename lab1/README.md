# Lab 1

O foco deste lab é colocar em prática os conceitos principais do Kubernetes, realizando o deploy de uma API e expondo a mesma para os usuários.

Também iremos explorar o uso da command line responsável pela comunicação com o cluster, o [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/).

## Conceitos abordados
* kubectl
* Deployment
* Pod
* ConfigMap
* Service
* Ingress

## Mãos à obra


1. Realize o deploy da API "greeting-app"
```terminal
kubectl apply -f deploy.yml
```

2. Verifique se o deploy foi realizado com sucesso
```terminal
kubectl get pods
```

3. Caso o deploy não tenha sido realizado com sucesso, utilize o comando abaixo para obter mais detalhes sobre o status do Pod
```terminal
kubectl describe pod <pod-name>
```

4. Aplique o ConfigMap com as variáveis necessárias para a API executar
```terminal
kubectl apply -f configmap.yml
```

5. Para forçar a criação de um novo Pod, com as novas variáveis de ambiente definidas, é possível deletar antigo Pod

```terminal
kubectl delete pod <pod-name>
```

Todavia, o K8s tentará o retry do Pod durante intervalos de tempo, tentando fazer com que o estado atual do Deployment seja igual ao desejado.


6. Realize a criação do Service, para que o Pod possa ser visível dentro da rede do cluster
```terminal
kubectl apply -f service.yml
```

6. Realize a criação do Ingress, para que possamos expor o Service para fora do cluster
```terminal
kubectl apply -f ingress.yml
```

7. Requisitando a API através do ingress
```terminal
curl <minikube-ip>/api/greetings
```