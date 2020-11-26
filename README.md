# BEXS

![Bexs CI](https://github.com/domarcio/bexs/workflows/Go/badge.svg?branch=main)

Encontrar a **melhor rota (conexões)** independentemente do valor ($). Feito em **GO**!

1. **[Sobre o projeto](#sobre-o-projeto)**;
2. **[Estrutura & Design adotados](#estrutura--design-adotados)**;
3. **[Como executar](#como-executar)**;
4. **[API](#api)**.

## Sobre o projeto
O objetivo é encontrar a **melhor alternativa (valor mais baixo)** para um viajante. O viajante **não quer saber o número de conexões que o avião fará**, ele quer pagar mais barato.

Supondo que ele queira viajar de **São Paulo** (_source_) para o **Acre** (_target_), as opções disponíveis no momento são:
1. Saindo do aeroporto **SPA** em São Paulo com vôo direito para o **ACR** no Acre no valor de **$ 100.00**.
2. Saindo do aeroporto **SPA** em São Paulo com uma escala para o **RJA** no Rio de Janeiro por apenas **$ 20.00**, na sequência partirá outro vôo para o **ACR** no Acre no valor de **$ 50.00**.

Nós daremos a opção **2** para o nosso viajante, pois ele cumprirá seu objetivo pagando **$ 70.00**, uma econômia de **$ 30.00** :)

```
SPA - RJA - ACR > $ 70
```

## Estrutura & Design adotados
A estrutura é relativamente simples:

|Diretório|Objetivo|
|---------|--------|
|**bin**|Onde estão todos os binários da app|
|**config**|Configurações usadas pela app separadas por ambientes|
|**data**|Onde guardamos os logs, *storages files*, cache, etc|
|**docker**|Nosso Dockerfile e demais configurações usadas pela imagem|
|**driver**|Aqui é, de fato, as "portas de entrada" para a app, tanto API, *command line*, etc. Cada driver disponível é composto por um *main.go* que inicia a app|
|**scripts**|Scripts auxiliares para uso|
|**src**|Onde contém toda a lógica. É o coração da app|

### Design
A ideia foi separar o máximo possível em camadas e usar `interfaces` (sempre que possível) para elas se comunicarem entre si: Composição (até porque é feito em Go, né :).

Algumas peças chaves que foram pensadas:
* Uma *entity* não conhece a lógica do domínio.
* Um *service* é algo que faz alguma coisa, ele pode saber trabalhar com N *repositories* ou não, com serviços de *log* ou *metrics* (ou não).
* Um *service* **sempre tem** *interfaces* definidas com coisas no qual ele sabe trabalhar (externas ou não). Dessa forma garantimos baixo acoplamento e testes mais tranquilos.
* Todo *service*/*lib* **externo** (**que não faz parte da lógica do domńio**) fica dentro da *infra*, seja um serviço de *log*, *file manager* e/ou *repositories*.

### Script para encontrar a melhor rota com baixo custo
A rotina é relativamente simples: Usar uma função recursiva que sabe quando quando estamos no ínicio (*source*) e no fim (*target*), e dessa forma ela vai adicionando os percursos em um `map` enquanto não chegamos no final da nossa jornada.

> confesso que eu quebrei a cabeça para chegar nesse resultado. Caso tenha uma forma mais simples eu gostaria de saber :)


## Como executar
**ATENÇÂO:** Os comandos `make` usam o [Docker](./docker/Dockerfile). Caso prefira executar com o Golang na tua máquina, veja o arquivo [make.sh](scripts/make.sh)

### Criar a imagem
```bash
$ make image
```

### Gerar os binários para o *cmd* e *api*
```bash
$ make build-cmd
$ make build-api
```

### Executar via linha de comando
```bash
$ make run-cmd FILENAME=./data/storage/routes.csv
```

### Subir a API
```bash
$ make run-api
```

## API

Abaixo uma rápida doc sobre a API e como executá-la.

### Breve documentação

Nossa API consiste em dois endpoins, um para a criação das rotas e outro para consultar a melhor opção de baixo custo. Todos os *responses* dos *endpoints* são no formato **JSON**.

#### Endpoints
|PATH|Request Method|Campo|Obrigatório?|Objetivo|Características|
|----|--------------|-----|------------|--------|---------------|
|**/api/connection**|POST|source|Sim|Informar a rota de origem do vôo|`min=3,max=3,pattern=A-Z`|
|||target|Sim|Informar a rota de destino do vôo|`min=3,max=3,pattern=A-Z`|
|||price|Sim|Valor do vôo|`min=1,pattern=0-9`|
|**/api/cost**|GET|source|Sim|Informar a rota de origem do vôo|`min=3,max=3,pattern=A-Z`|
|||target|Sim|Informar a rota de destino do vôo|`min=3,max=3,pattern=A-Z`|

### Como executar

Tem o Postman? Pode usar:

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/67af6193814145e23ddc)

Se preferir, pode usar o `curl`:

#### Criar uma conexão
```bash
$ curl --location --request POST 'http://localhost:7007/api/connection' \
--header 'Content-Type: application/json' \
--data-raw '{
    "source": "SPA",
    "target": "ACR",
    "price": 100
}'
```

#### Encontrar a rota com menor custo
```bash
$ curl --location --request GET 'http://localhost:7007/api/cost?source=CDG&target=GRU' \
--header 'Content-Type: application/json'
```