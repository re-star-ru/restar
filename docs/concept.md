## Обсуждение концепции 

1 - предлагаю использовать vanilla golang, не использовать ORM - только нативный SQL и только нативные веши из GOLANG чтобы потом не иметь проблем с совместимостью.

2 - предлагаю для начала использовать встроенную кроссплатформенную базу например genji.dev если ее совсем не будет хватать возьмем чтоб нибудь помощнее типа postgres но думаю если правильно организовать то ее должно хватить

3 - первый запущенный бинарь автоматически организовывает кластер и генерит QR код или какую то ссылку котору потом скармливаем всем остальным участникам

4 - допустим все бинари сообщают в общий чат сколько у них есть ресурсов чтобы им могли закидывать данные которые не помесятся другим участникам

5 - допустим какому то бинарю можно выдать привелегию "надежный" это значит что данные на нем резервируются и его ресурсы можно будет использовать для бекапа данных остальных участников


вплане идея встроенной чтобы где бы не запустили бинарь - там бы база образовалась 
далее по сетевым каналам передать что в данной базе имеется (какой то легковесный индекс) "чтобы по разным местам одновременно можно было поискать" вообще супер будет.

типа делаем сетевые каналы
запустили пару бинарей на разных машинах
они бегут в какой то общий чат знакомятся
находят друг друга
и далее когда спрашиваешь где данные они у себя смотрят и выдают если такие данные у них есть в общий канал
чтобы сколько бы не добавили бинарников они только наращивали мощность системы

и может как то резервировать эти данные на сервере где больше размер диска - допустим отдельным гуфером

в базе ставить отметку время жизни данных и отталкиваясь от этого выбирать время бекапа данных (допустим в три раза дольше) 


короче чтобы все что нужно было сделать для установки программы это просто скачать бинарник и запустить
и он предложит как то авторизоваться (допустим скормить ему QR код своего кластера) и оттуда узнает в каком кластере и в каких чатах он сидит