Разработайте сетевую службу по аналогии с сервером времени (не RPC),
которая бы каждому подключившемуся клиенту показывала раз в 3 секунды
случайную Go-поговорку. Поговорки возьмите с сайта.

Служба должна поддерживать множественные одновременные подключения. (сделаю 3 посредством
реализации простейшего семафора)
Служба НЕ ДОЛЖНА!!! завершать соединение с клиентом.
Вы должны проверить работу приложения с помощью telnet.