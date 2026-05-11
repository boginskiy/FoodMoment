### Особенности

Сообщения в Kafka идут только по конкретно заданному уровню логирования. Например "INFO"




TODO ...
logger/
├── logger.go          // Logger — публичный фасад
├── preparer.go        // LogPreparer (или Serializer)
├── sender.go          // KafkaSender (или MessageSender)
├── config.go          // настройки для всех слоев
└── logger_test.go


Старое название	Новое название	Роль
loggProducer.go	KafkaLogger	Публичный фасад, точка входа
preparProducer.go	LogEnricher или MessageBuilder	Подготовка + сериализация
producer.go	KafkaWriter	Чистая отправка в брокер