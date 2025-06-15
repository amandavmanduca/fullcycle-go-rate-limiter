
### Desafio técnico Rate Limiter

#### 1. Preparando ambiente

Criar arquivo .env baseado no .env.example
```bash
make env
```

Explicação das Variáveis de ambiente
```
API_PORT -- porta do servidor
API_KEY -- chave válida da api

REDIS_HOST -- host do redis
REDIS_PORT -- porta do redis
REDIS_PASSWORD -- senha do redis

IP_RATE_LIMIT -- limite de requests disponível por ip
API_KEY_RATE_LIMIT -- limite de requests disponível pela chave da api
RATE_LIMIT_IP_INTERVAL_IN_SECONDS -- intervalo para o limite de requests por ip (se 0 = desabilitado)
RATE_LIMIT_KEY_INTERVAL_IN_SECONDS -- intervalo para o limite de requests por chave de api (se 0 = desabilitado)
```

### 2. Executando o projeto com Docker

```bash
make run
```

### 3. Exemplos de Requests

#### 3.1 Rota aberta

```bash
curl http://localhost:8080

```


#### 3.2 Rota com api key

```bash
curl -H "API_KEY: secretkey" http://localhost:8080

```

### 4. Testes

Para a execução dos testes (docker dependency) utilize os comandos:
```bash
$ setup-local-tests

$ test

$ teardown-local-tests
```

Ou para executar a combinação dos comandos
```bash
make test-local
```