Команда для сборки docker образа: docker build . -t=server \
Команда для запуска приложения: docker-compose up \
\
Api:\
  http://localhost:8080/ping               --- проверка работоспособности сервера и БД\
  http://localhost:8080/set/{key}/{value}  --- добавить в базу данных ячейку с ключом {key} и значением {value}\
  http://localhost:8080/get/{key}          --- получить значение ячейки с ключом {key}, если такого ключа нет, то вернётся сообщение "key does not exist" \
