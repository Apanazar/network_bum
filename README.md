## Directory Tree:

```
├── 📂 code
│   ├── go.mod
│   ├── go.sum
│   └── netstalk.go
├── 📂 log
│   ├── database.txt
│   └── ip.txt
├── 📂 screen
├── rand-stalk
|── screen.py
├── out.json
└── result.txt
```


## EN
```This project was written for random travel on the network by generating random IP addresses and scanning them```
1. ```./rand-stalk gen N``` — generation of random IP addresses by the number of N , as well as their recording in ip.txt and database.txt

2. ```macscan --rate=Xp 80,8080  -iL log/ip.txt --output-format json --output-file out.json``` — scanning of IP addresses from a file ip.txt in the log directory, followed by writing the working addresses to the out file.json

3. ```./rand-stalk format``` — formats IP addresses from out.json to a human-readable template format IP:PORT followed by writing to a file result.txt

* ```database.txt``` the file is used to store all generated addresses, to avoid repetition during subsequent generation

* ```screen.py``` script takes screenshots of addresses from a file result.txt and saves them in the screen directory

* ```geckodriver``` is needed for Selenium to work in screen.py

> From the project: [ꜱᴜʀ.ɴᴇᴛꜱᴛᴀʟᴋɪɴɢ](https://t.me/sur_NET)

## RU
```Данный проект писался для рандомного путешествия по сети путем генерации случайных IP-адресов и их сканирования```
1. ```./rand-stalk gen N``` — генерация рандомных IP-адресов количеством N ,так же их запись в ip.txt и database.txt

2. ```masscan --rate=X -p80,8080  -iL log/ip.txt --output-format json --output-file out.json``` — сканирование IP-адресов из файла ip.txt в директории log с последующей записью рабочих адресов в файл out.json

3. ```./rand-stalk format``` — отформатирует IP-адреса из out.json в удобочитаемый формат шаблона IP:PORT с последующей записью в файл result.txt


* ```database.txt``` файл используется для хранения всех сгенерированных адресов, для исключения повторения при последующей генерации

* ```screen.py``` скрипт делает скриншоты адресов из файла result.txt и сохраняет их в директории screen

* ```geckodriver``` нужен для работы Selenium в screen.py

> От проекта: [ꜱᴜʀ.ɴᴇᴛꜱᴛᴀʟᴋɪɴɢ](https://t.me/sur_NET)
