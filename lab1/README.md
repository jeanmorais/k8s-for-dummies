# Lab 1

O foco deste lab é colocar em prática os conceitos principais do Kubernetes, realizando o deploy de uma API e expondo-a para os usuários.

Também iremos explorar o uso da command line responsável pela comunicação com o cluster, o [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/).

## Conceitos abordados
* kubectl
* Deployment
* Pod
* ConfigMap
* Service
* Ingress

## API

A API "greeting-app" é escrita em Go, e possui apenas um endpoint que retorna na index uma saudação em diversos idiomas:

* GET `http://<host>/`

## Mãos à obra

1. Realize o deploy da API *greeting-app*
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

4. Crie o ConfigMap com as variáveis de ambiente necessárias para a API executar corretamente
```terminal
kubectl apply -f configmap.yml
```

5. Para forçar a criação de um novo Pod, com as novas variáveis de ambiente definidas, é possível deletar o  Pod defeituoso

```terminal
kubectl delete pod <pod-name>
```

Todavia, o K8s tentará o retry do Pod durante intervalos de tempo, tentando fazer com que o estado atual do Deployment seja igual ao desejado. Mais detalhes em [Restart Policy](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy).


6. Realize a criação do Service, para que o Pod possa ser visível dentro da rede do cluster
```terminal
kubectl apply -f service.yml
```
7. Para que possamos expor nossa API para o público, vamos precisar de um [Ingress Controller](https://kubernetes.io/docs/concepts/services-networking/ingress-controllers/) instalado no cluster.
O minikube, como é uma ferramenta voltada para desenvolvimento, já possui um [NGINX Ingress controller](https://kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/) embutido. Para instalar basta executar o seguinte comando:

```terminal
minikube addons enable ingress
```

8. Verifique se o NGINX Ingress está em execução
```terminal
kubectl get pods -n kube-system
```

9. Realize a criação do [Ingress Resource](https://kubernetes.io/docs/concepts/services-networking/ingress/#the-ingress-resource), para que possamos expor o Service para fora do cluster
```terminal
kubectl apply -f ingress.yml
```

10. Requisitando a API através do ingress
```terminal
curl <minikube-ip>/api/greetings
```