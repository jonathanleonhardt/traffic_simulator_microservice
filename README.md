
microserviço de simulação de tráfego para imersao full cycle

---

Inicializar o container com go rodando:     
```
docker-compose up -d
```

ver se o container de nome ```traffic_simulator``` esta no ar:
```
docker ps
```

entrar no container em modo interativo
```
docker exec -it traffic_simulator bash
```

perceba que os arquivos desta pasta são compartilhados com o container, ou seja, adicionando ou modificando aqui ou la da na mesma.    
Uma vantagem é não precisar nem do golang instalado na sua maquina.

---

Como dar start no customer e enviar a mensagem para ele de dentro do container go

inicie o container kafka
```
docker exec -it kafka-kafka-1 bash
```

execute o comando para subir o consumer
```
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=readteste
```

de dentro do container go
```
docker exec -it traffic-simulator bash
# execute
go run main.go
```

to create a producer 
```
kafka-console-producer --bootstrap-server=localhost:9092 --topic=readteste
```

