### English Version

# Project: Example of Using Events in Ethereum  

This project demonstrates how to work with events in Ethereum smart contracts and process them in a Go application.  

## Description  

The `CustomERC721` smart contract implements the ERC721 standard and includes functionality for:  
- Minting tokens with unique URIs.  
- Listing tokens for sale.  
- Purchasing tokens.  

The Go application connects to an Ethereum node, subscribes to smart contract events, and processes them:  
- `TokenMinted`: Event triggered when a new token is created.  
- `TokenListedForSale`: Event triggered when a token is listed for sale.  
- `TokenSold`: Event triggered when a token is purchased.  

## Installation and Launch  

1. Ensure that Go is installed and an Ethereum node (e.g., Ganache) is set up.  
2. Clone the repository:  
    ```bash  
    git clone <repository URL>  
    cd go-test  
    ```  
3. Install dependencies:  
    ```bash  
    go mod tidy  
    ```  
4. Deploy the `CustomERC721` smart contract to your Ethereum node.  
5. Update the contract address in the `main.go` file.  
6. Run the application:  
    ```bash  
    go run main.go  
    ```  

## Usage Example  

1. Deploy the contract and mint a token:  
    ```solidity  
    contract.mintToken("https://example.com/token/1");  
    ```  
2. Subscribe to events in the Go application.  
3. Observe the event output in the console.  

## Dependencies  

- Go 1.23+  
- Libraries:  
  - `github.com/ethereum/go-ethereum`  
  - `github.com/ethereum/go-verkle`  
  - Others (see `go.mod`).  

## License  

The project is distributed under the MIT license.  


---


### Russian Version

# Проект: Пример использования событий в Ethereum  

Этот проект демонстрирует, как работать с событиями в смарт-контрактах Ethereum и обрабатывать их в приложении на Go.  

## Описание  

Смарт-контракт `CustomERC721` реализует стандарт ERC721 и включает функционал для:  
- Создания токенов с уникальными URI.  
- Выставления токенов на продажу.  
- Покупки токенов.  

Приложение на Go подключается к узлу Ethereum, подписывается на события смарт-контракта и обрабатывает их:  
- `TokenMinted`: Событие, вызываемое при создании нового токена.  
- `TokenListedForSale`: Событие, вызываемое при выставлении токена на продажу.  
- `TokenSold`: Событие, вызываемое при покупке токена.  

## Установка и запуск  

1. Убедитесь, что Go установлен, а узел Ethereum (например, Ganache) настроен.  
2. Клонируйте репозиторий:  
    ```bash  
    git clone <repository URL>  
    cd go-test  
    ```  
3. Установите зависимости:  
    ```bash  
    go mod tidy  
    ```  
4. Разверните смарт-контракт `CustomERC721` на вашем узле Ethereum.  
5. Обновите адрес контракта в файле `main.go`.  
6. Запустите приложение:  
    ```bash  
    go run main.go  
    ```  

## Пример использования  

1. Разверните контракт и создайте токен:  
    ```solidity  
    contract.mintToken("https://example.com/token/1");  
    ```  
2. Подпишитесь на события в приложении на Go.  
3. Наблюдайте за выводом событий в консоли.  

## Зависимости  

- Go 1.23+  
- Библиотеки:  
  - `github.com/ethereum/go-ethereum`  
  - `github.com/ethereum/go-verkle`  
  - Другие (см. `go.mod`).  

## Лицензия  

Проект распространяется под лицензией MIT.  