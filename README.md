# Gontabilizador
Trata-se de um contabilizador de presenças em ensaios para ritmistas. Tanto a comunicação com o banco de dados como o servimento e carregamento de arquivos estáticos (html, css) são realizados por **Go**. O banco de dados foi construído com **MySQL**.

Além disso, também há um script em **Python** que escuta por atualizações no banco de dados, escrevendo no terminal atualizações de inserção na tabela de presença.

## Uso
Primeiramente, faça o clone do repositório
```bash
git clone https://github.com/compermane/gontabilizador
```

Feito o clone, inicie o minikube ou execute o arquivo build.sh (talvez seja necessário executar [chmod +x build.sh])
```bash
minikube start
```

Então, entre na raíz do projeto e instale o projeto em Helm. Isso fará com que os templates em ./gontabilizador/templates sejam aplicados automaticamente pelo minikube.
```bash
helm install gontabilizador ./gontabilizador
```
Então, espere algum tempo para todos os pods estiver em "Running" e acesse o endereço k8s.local

Alternativamente, para obter o endereço da aplicação, execute os seguintes comandos
```bash
export NODE_PORT=$(kubectl get --namespace default -o jsonpath="{.spec.ports[0].nodePort}" services gontabilizador)
export NODE_IP=$(kubectl get nodes --namespace default -o jsonpath="{.items[0].status.addresses[0].address}")
echo http://$NODE_IP:$NODE_PORT
```

Ou, caso esteja usando o kubectl do minikube, execute
```bash
export NODE_PORT=$(minikube kubectl -- get --namespace default -o jsonpath="{.spec.ports[0].nodePort}" services gontabilizador)
export NODE_IP=$(minikube kubectl -- get nodes --namespace default -o jsonpath="{.items[0].status.addresses[0].address}")
echo http://$NODE_IP:$NODE_PORT
```

Para atualizar a aplicação com atualizações feitas em algum template, execute
```bash
helm upgrade gontabilizador ./gontabilizador
```

## Desinstalando
Para desinstalar a aplicação e, como consequência, deletar todos os templates carregados, execute

```bash
helm delete gontabilizador
```

### Autor
Eugênio Akinori Kisi Nishimiya RA: 811598
