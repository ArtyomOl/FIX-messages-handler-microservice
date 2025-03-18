# FIX-messages-handler-microservice

Микросервис, который обрабатывает сообщения о биржевых транзакциях в формате протокола FIX и, исходя из получаемых данных, строит биржевой стакан для ценных бумаг.  

## Технологии
`language` ~ `Go`  
`database` ~ `Redis`  
`message broker` ~ `Apache Kafka`

## Принцип работы
Микросервис получает сообщения по протоколу FIX через Kafka и сохраняет расшифрованные данные в Redis.  
Через API можно получать информацию о биржевом стакане для конкретной акции.  
  
**Пример FIX-сообщения:**  
`8=FIX.4.2|55=AAPL|44=213|38=15|54=1`  
что означает следующее: совершена заявка на покупку (54=1) 15 (38=15) лотов акций компании Apple (55=AAPL) по цене 213$ (44=213)  
  
**Пример сформированного биржевого стакана (в json-формате):**
```json
{
    "symbol": "AAPL"  
    "bidsprices": [213.45,213.85,214.34,214.80]  
    "asksprices": [212.67,213.05,213.40,214.80]  
    "bidsquantity": [100,200,300,150]  
    "asksquantity": [130,250,220,80] 
} 
```
где `bidsprices` и `askprices` - цены предложения и спроса, `bidsquantity` и `askquantity` - объёмы предложения и спроса соответственно  


## Взаимодействие с микросервисом:

**API**  
POST:  
`curl localhost:8080/api/getorderbook --data "{"symbol": <string>,"depth": <int>}"`  
(symbol - тикер акции, depth - глубина стакана)
  
**Отправка сообщений через Kafka:**  
Сообщение должно быть в формате протокола FIX.  
Пример кода на Go:  
```go
func SendToKafka(message string) error {
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		return err
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: "fix-messages",
		Value: sarama.StringEncoder(message),
	}

	_, _, err = producer.SendMessage(msg)
	return err
}
```
Функция `main`:
```go
func main() {
	err := SendToKafka("8=FIX.4.2|55=AAPL|44=213|38=15|54=1")
	if err != nil {
		panic(err)
	}
}
```

### Планы на будущее:
- Возможность отправки сообщения на закрытие сделки
- Поддержка протоколов FIX разных версий