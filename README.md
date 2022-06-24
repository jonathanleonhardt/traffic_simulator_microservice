
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

