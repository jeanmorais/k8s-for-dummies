# k8s-for-dummies 
- [Setup](#setup)
  - [Minikube](#minikube)

# Setup

## Minikube

O [Minikube](https://kubernetes.io/docs/setup/learning-environment/minikube/) é uma ferramenta *open source* que permite a execução de um cluster de Kubernetes em uma máquina local. O Minikube inicia um cluster de nó único localmente, com baixa utilização de recursos. É ideal para testes de desenvolvimento e realização de POC's.

### Instalação

O site oficial fornece um breve tutorial para a instalação do Minikube. Basta acessar [Install Minikube
](https://kubernetes.io/docs/tasks/tools/install-minikube/).

Os requisitos para instalação do Minikube estão [mencionados no tutorial](https://kubernetes.io/docs/tasks/tools/install-minikube/#before-you-begin), mas vale frisar:
- Tenha um hypervisor instalado. Sugestão: [VirtualBox](https://www.virtualbox.org/wiki/Downloads).
- Instale o [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/#install-kubectl-on-linux) (ferramenta de linha de comando para interagir com o cluster)

