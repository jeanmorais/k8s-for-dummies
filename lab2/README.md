# Lab 2

O foco deste lab é explorar um pouco mais dos recursos do [kubectl](https://kubernetes.io/pt/docs/reference/kubectl/cheatsheet/), executando comandos de forma imperativa.

Existem duas maneiras de criar recursos em um cluster Kubernetes: o imperativo e o declarativo.

* A abordagem declarativa é usada para criar recursos a partir de arquivos de manifest (geralmente YAML) usando o comando `kubectl apply`. Essa é a abordagem usada em um ambiente de produção.

* A forma imperativa é usada para gerenciar recursos usando diferentes comandos e não requer nenhum arquivo de manifest. Não é a maneira indicada para se utilizar em produção, mas pode ser útil para testes.

Também iremos explorar alguns conceitos dos Services e o serviço de DNS do K8s.

## Conceitos abordados
* kubectl
* Deployments / Pods
* Logs
* Services
* DNS

## Mãos à obra

1. Crie um Deployment usando o seguinte comando:

```terminal
kubectl create deployment web --image=gcr.io/google-samples/hello-app:1.0
```

2. Agora vamos expor este Deployment como um NodePort:
```terminal
kubectl expose deployment web --type=NodePort --port=8080
```

3. Agora vamos verificar se o Service está criado e disponível como um NodePort:
```terminal
kubectl get service web
```
* A saída deve ser similar a esta:
```terminal
NAME   TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
web    NodePort   10.96.194.208   <none>        8080:32755/TCP   10m
```

* Isso significa que o Deployment está acessível através do IP do nosso cluster, e pode ser acessível externamente através da porta `32755`.

4. Acesse o Service via o NodePort:
```terminal
minikube service web --url
```
* A saída deve ser similar a esta:
```terminal
http://192.168.99.100:32755
```
```terminal
curl http://192.168.99.100:32755

Hello, world!
Version: 1.0.0
Hostname: web-557f59c6cf-7nsjj
```

:heavy_check_mark: Conseguimos acessar via o IP do nosso cluster (Minikube) o Service através de um NodePort.

<br />


5. Continuando o roteiro, vamos criar um novo Deployment:
```terminal
kubectl create deployment web2 --image=gcr.io/google-samples/hello-app:2.0
```

6. Agora vamos expor este Deployment como um ClusterIP:
```
kubectl expose deployment web2 --type=ClusterIP --port=80 --target-port=8080
```

* A saída deve ser similar a esta:
```terminal 
NAME   TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)   AGE
web2   ClusterIP   10.96.95.55   <none>        80/TCP    37s
```

## Faça você mesmo 

Dado um cenário bem comum, de multiplos serviços precisando realizar requisições entre si dentro do cluster. Como poderíamos realizar uma chamada de um Pod escalonado pelo Deployment "web" para o "web2"?

> Os Pods escalonados pelo Deployment "web2" estão visíveis apenas dentro do cluster, pois foram expostos através de um Service do tipo ClusterIP.

O primeiro passo seria buscar algum Pod com que esteja com o seletor "web".
<details>
<summary>Spoiler</summary><br />

```terminal
kubectl get pods --selector=app=web
NAME                   READY   STATUS    RESTARTS   AGE
web-557f59c6cf-7nsjj   1/1     Running   0          42m
```
</details><br /><br />

Agora, tente acessar este Pod, entrando no seu console.

<details>
<summary>Spoiler</summary><br />

```terminal
kubectl exec -it web-557f59c6cf-7nsjj /bin/sh
```
</details><br /><br />

Tente fazer uma requisição para o service "web2" utilizando o `curl`.

<details>
<summary>Spoiler</summary><br />

Para qual endereço devo requisitar? 

Por padrão, o DNS de um Service segue o seguinte formato:
`web2.default.svc.cluster.local`, sendo `web2` o nome do Service e `default` o namespace.

Dentro do Pod, vamos instalar o `curl` e testar a request.

```terminal

/ # apk update && apk add curl

/ # curl web2.default.svc.cluster.local
Hello, world!
Version: 2.0.0
Hostname: web2-644ffbbb4-9rb7b

```

Feito! :tada:

</details>

