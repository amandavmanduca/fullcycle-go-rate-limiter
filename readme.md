
### Desafio técnico Rate Limiter

#### 1. Preparando ambiente

Criar arquivo .env baseado no .env.example
```bash
make env
```

### 2. Executando o projeto

Docker
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