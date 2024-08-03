# criando tópico do kafka

kafka-topics --create --bootstrap-server=localhost:9092 --topic=test --partitions=3

# consumindo tópico pelo terminal

kafka-console-consumer --bootstrap-server=localhost:9092 --topic=test
