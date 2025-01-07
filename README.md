# Labs Go Expert - Observabilidade & Open Telemetry

## Objetivo
Desenvolver um sistema em Go que receba um CEP, identifique a cidade e retorne o clima atual (temperatura em graus Celsius, Fahrenheit e Kelvin), juntamente com o nome da cidade. O sistema deve implementar OpenTelemetry (OTEL) e Zipkin para rastreamento distribuído.

## Contexto
O projeto está dividido em dois serviços principais:

- **Serviço A (zipcode-api)**: Responsável por receber o input do usuário.
- **Serviço B (weather-api)**: Responsável pela orquestração da busca de informações do CEP e das temperaturas correspondentes.

---

## Requisitos do Sistema

### Serviço A (Input)
1. O sistema deve receber um input de 8 dígitos via POST, utilizando o seguinte schema:
    ```json
    {
      "cep": "29902555"
    }
    ```
2. Validação do CEP:
    - Deve conter exatamente 8 dígitos.
    - Deve ser uma string.
3. Comportamento em casos de validação:
    - CEP válido: Encaminhar para o Serviço B via HTTP.
    - CEP inválido: Retornar a seguinte resposta:
      - Código HTTP: **422**
      - Mensagem: `invalid zipcode`

### Serviço B (Orquestração)
1. Deve receber um CEP válido de 8 dígitos.
2. Realizar uma pesquisa do CEP e identificar o nome da localização correspondente.
3. Consultar as temperaturas e formatá-las nas seguintes escalas:
    - **Celsius**
    - **Fahrenheit**
    - **Kelvin**
4. Respostas esperadas:
    - **Sucesso**:
      - Código HTTP: **200**
      - Response Body:
        ```json
        {
          "city": "São Paulo",
          "temp_C": 28.5,
          "temp_F": 83.3,
          "temp_K": 301.65
        }
        ```
    - **Falha - CEP inválido (formato incorreto)**:
      - Código HTTP: **422**
      - Mensagem: `invalid zipcode`
    - **Falha - CEP não encontrado**:
      - Código HTTP: **404**
      - Mensagem: `can not find zipcode`

---

## Integração OTEL e Zipkin

### Objetivos da Implementação
- Implementar rastreamento distribuído entre o **Serviço A** e o **Serviço B**.
- Utilizar spans para medir o tempo de resposta das seguintes operações:
  - Busca de informações do CEP.
  - Consulta de dados climáticos.

### Benefícios Esperados
- Observabilidade completa das requisições entre os serviços.
- Identificação de gargalos nas operações.
- Medidas claras de desempenho e latência.

---

## Executando o Sistema

1. Certifique-se de ter o Docker e o Docker Compose instalados.
2. Clone o repositório e acesse o diretório raiz.
3. Execute o comando:
    ```bash
    docker compose up --build
    ```
4. Os serviços estarão disponíveis nas seguintes portas:
    - **Serviço A**: `http://localhost:8081`
    - **Serviço B**: `http://localhost:8080`
    - **Zipkin UI**: `http://localhost:9411`

---

## Testando o Sistema

### Serviço A
- Endpoint: `POST /zipcode`
- Exemplo de requisição:
    ```bash
    curl -X POST http://localhost:8081/zipcode \
         -H 'Content-Type: application/json' \
         -d '{"cep": "29902555"}'
    ```

### Serviço B
- Endpoint: Interno, acessado pelo Serviço A.

### Observabilidade
- Acesse o **Zipkin UI** em `http://localhost:9411` para visualizar o rastreamento distribuído.
