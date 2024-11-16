
# Bank Transfer System

Um sistema de transferência bancária simples, desenvolvido em **GoLang**, para praticar conceitos como arquitetura hexagonal, manipulação de APIs REST, integração com MongoDB e implementação de funcionalidades avançadas como conversão de moedas e histórico de transferências.

## 🛠️ Funcionalidades

- ✅ Registro de usuários com saldo inicial.
- ✅ Transferência de saldo entre contas, com:
  - Taxa fixa ou percentual.
  - Limite máximo por transferência.
  - Conversão automática entre moedas (USD, BRL, EUR).
- ✅ Histórico de transferências para auditoria.
- ✅ Autenticação JWT para proteger os endpoints.
- ✅ Integração com MongoDB para persistência.

---

## 🏗️ Arquitetura

O projeto segue o padrão de **Arquitetura Hexagonal (Ports and Adapters)**, separando a lógica de negócios, interfaces e infraestrutura:

- **Core (Núcleo):** Regras de negócios e lógica central (`internal/core`).
- **Adapters:** Camadas de infraestrutura para acesso a banco de dados e serviços externos (`internal/infra`).
- **Application:** Configuração de rotas e endpoints HTTP (`internal/app`).
- **Entrypoint:** Ponto de entrada para iniciar o servidor (`cmd/main.go`).

---

## 🚀 Tecnologias Utilizadas

- **Linguagem:** GoLang
- **Banco de Dados:** MongoDB
- **Autenticação:** JWT
- **Frameworks/Libs:**
  - `github.com/gorilla/mux` para roteamento.
  - `github.com/golang-jwt/jwt/v4` para autenticação.
  - `go.mongodb.org/mongo-driver` para integração com MongoDB.

---

## 📦 Estrutura do Projeto

```
bank-transfer-system/
├── cmd/
│   └── main.go                # Ponto de entrada da aplicação
├── internal/
│   ├── app/
│   │   ├── config/            # Configuração do sistema (ex.: MongoDB, JWT)
│   │   ├── handlers/          # Handlers HTTP
│   │   └── routes/            # Configuração de rotas
│   ├── core/
│   │   ├── user/              # Regras de negócio de usuários
│   │   ├── transfer/          # Lógica de histórico de transferências
│   │   └── currency/          # Conversão de moedas
│   └── infra/
│       ├── db/                # Repositórios e acesso ao MongoDB
│       └── http/              # Middleware e inicialização do servidor
└── README.md                  # Documentação do projeto
```

---

## ⚙️ Configuração e Execução

### 1. Clone o Repositório
```bash
git clone https://github.com/seu-usuario/bank-transfer-system.git
cd bank-transfer-system
```

### 2. Configure o Banco de Dados
Certifique-se de que o MongoDB está rodando e crie as coleções necessárias:
- **users**
- **CurrencyRates**
- **TransferHistory**

### 3. Configure o Arquivo `.env`
Crie um arquivo `.env` com as seguintes variáveis:
```env
MONGO_URI=mongodb://localhost:27017
DATABASE_NAME=bank_transfer
JWT_SECRET=supersecretkey
SERVER_PORT=8080
```

### 4. Execute o Projeto
Instale as dependências e inicie o servidor:
```bash
go mod tidy
go run cmd/main.go
```

---

## 🛠️ Endpoints Disponíveis

### **1. Registrar Usuário**
- **POST** `/users`
- **Body (JSON):**
  ```json
  {
      "name": "Otávio Mendes",
      "email": "otavio.mendes@example.com",
      "balance": 1000.00
  }
  ```
- **Resposta:**
  ```json
  {
      "message": "Usuário criado com sucesso"
  }
  ```

### **2. Realizar Transferência**
- **POST** `/transfer`
- **Body (JSON):**
  ```json
  {
      "from_id": "user1_id",
      "to_id": "user2_id",
      "amount": 100.00,
      "from_currency": "USD",
      "to_currency": "BRL"
  }
  ```
- **Resposta:**
  ```json
  {
      "message": "Transferência realizada com sucesso"
  }
  ```

### **3. Listar Usuários**
- **GET** `/users`
- **Resposta:**
  ```json
  [
      {
          "id": "user1_id",
          "name": "Otávio Mendes",
          "email": "otavio.mendes@example.com",
          "balance": 1000.00
      }
  ]
  ```

### **4. Consultar Histórico de Transferências**
- **GET** `/users/{id}/transfers`
- **Resposta:**
  ```json
  [
      {
          "from_id": "user1_id",
          "to_id": "user2_id",
          "amount": 100.00,
          "currency": "BRL",
          "created_at": "2024-11-14T15:00:00Z"
      }
  ]
  ```

---

## 🤔 Melhorias Futuras

- 📊 Relatórios detalhados de transferências.
- 🌍 Integração com APIs externas para taxas de câmbio em tempo real.
- 🔐 Melhorar autenticação com suporte a permissões (Admin, Usuário).
- 💬 Suporte a notificações (e.g., via e-mail ou WebSocket).

---

## 🧑‍💻 Contribuições

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar PRs. 😉

---

## 📄 Licença

Este projeto está licenciado sob a [MIT License](LICENSE).

---

**Desenvolvido com 💻 e ☕ por Otávio Mendes.**
