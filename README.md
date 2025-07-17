# Gontabilizador
Trata-se de um contabilizador de presenças em ensaios para ritmistas. Tanto a comunicação com o banco de dados como o servimento e carregamento de arquivos estáticos (html, css) são realizados por **Go**. O banco de dados foi construído com **MySQL**.

Além disso, também há um script em **Python** que escuta por atualizações no banco de dados, escrevendo no terminal atualizações de inserção na tabela de presença.

## Uso
Primeiramente, faça o clone do repositório
```bash
git clone https://github.com/compermane/gontabilizador
```

Feito o clone, inicie o minikube
```bash
minikube start
```

Então, entre na raíz do projeto e instale o projeto em Helm. Isso fará com que os templates em ./gontabilizador/templates sejam aplicados automaticamente pelo minikube.
```bash
helm install gontabilizador ./gontabilizador
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
