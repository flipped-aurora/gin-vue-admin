
<div align=center>
<img src="http://qmplusimg.henrongyi.top/gvalogo.jpg" width="300" height="300" />
</div>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.18-blue"/>
<img src="https://img.shields.io/badge/gin-1.9.1-lightBlue"/>
<img src="https://img.shields.io/badge/vue-3.3.4-brightgreen"/>
<img src="https://img.shields.io/badge/element--plus-2.3.8-green"/>
<img src="https://img.shields.io/badge/gorm-1.25.2-red"/>
</div>

[English ](./README-en.md) | [简体中文](./README.md) | Русский

[gitee](https://gitee.com/pixelmax/gin-vue-admin): https://gitee.com/pixelmax/gin-vue-admin

[github](https://github.com/flipped-aurora/gin-vue-admin): https://github.com/flipped-aurora/gin-vue-admin

# Руководство по проекту

[Онлайн-документация](https://www.gin-vue-admin.com/) : https://www.gin-vue-admin.com/

[Видеоуроки от настройки окружения до развертывания](https://www.bilibili.com/video/BV1fV411y7dT)

[Этапы разработки](https://www.gin-vue-admin.com/guide/start-quickly/env.html) (Авторы: <a href="https://github.com/LLemonGreen">LLemonGreen</a> и <a href="https://github.com/fkk0509">Fann</a>)

## 1. Основная информация

### 1.1 О проекте

> Gin-vue-admin — это система управления бэкендом, основанная на [Vue](https://vuejs.org) и [Gin](https://gin-gonic.com), с разделением фронтенда и бэкенда. Она интегрирует JWT-аутентификацию, динамическую маршрутизацию, динамическое меню, аутентификацию Casbin, генератор форм, генератор кода и другие функции. Проект предоставляет множество примеров файлов, позволяя вам сосредоточиться на разработке бизнес-логики.

[Онлайн-демо](http://demo.gin-vue-admin.com): http://demo.gin-vue-admin.com

Имя пользователя: admin

Пароль: 123456

### 1.2 Руководство по участию

Привет! Спасибо, что выбрали gin-vue-admin.

Gin-vue-admin — это фреймворк с разделением фронтенда и бэкенда для разработчиков, дизайнеров и менеджеров продуктов.

Мы рады, что вы заинтересованы в участии в проекте gin-vue-admin. Прежде чем отправить свой вклад, пожалуйста, ознакомьтесь со следующими рекомендациями.

#### 1.2.1 Рекомендации по созданию Issues

- Issues предназначены исключительно для сообщений об ошибках, запросов новых функций и обсуждения дизайна. Другие вопросы могут быть закрыты напрямую. Если у вас возникли вопросы по использованию Element, обратитесь за помощью в [Gitter](https://gitter.im/element-en/Lobby).

- Перед созданием issue проверьте, не были ли уже зарегистрированы похожие проблемы.

#### 1.2.2 Рекомендации по Pull Request

- Сделайте форк этого репозитория на свой аккаунт. Не создавайте ветки в этом репозитории.

- Формат сообщения коммита должен быть: `[Имя файла]: Описание коммита.` (например, `README.md: Исправлена ошибка xxx`)

- <font color=red>Убедитесь, что Pull Request создается в ветку `develop`, а не в `master`.</font>

- Если ваш PR исправляет ошибку, предоставьте описание связанной ошибки.

- Для принятия PR требуется участие двух мейнтейнеров: один проверяет и одобряет изменения, второй проверяет и объединяет их.

### 1.3 Список версий

- master: Код версии 2.0 для продакшена
- develop: Код версии 2.0 для тестирования
- [gin-vue-admin_v2_dev](https://github.com/flipped-aurora/gin-vue-admin/tree/gin-vue-admin_v2_dev) (версия 2.0 с [GormV1](https://v1.gorm.io), стабильная ветка)
- [gva_gormv2_dev](https://github.com/flipped-aurora/gin-vue-admin/tree/gva_gormv2_dev) (версия 2.0 с [GormV2](https://v2.gorm.io), ветка разработки)

## 2. Начало работы

```
- Версия Node > v8.6.0
- Версия Golang >= v1.14
- Рекомендуемая IDE: Goland
- Инициализация проекта: разные версии базы данных не инициализированы. См. подробности в документации по инициализации https://www.gin-vue-admin.com/docs/first
- Замените публичный ключ, приватный ключ, имя хранилища и URL по умолчанию для Qiniuyun в проекте, чтобы избежать путаницы с тестовыми файлами.
```

### 2.1 Серверная часть проекта

Используйте `Goland` или другие редакторы кода, откройте каталог `server`. Не открывайте корневой каталог `gin-vue-admin`.

```bash
# Клонировать проект
git clone https://github.com/flipped-aurora/gin-vue-admin.git

# Перейти в каталог server
cd server

# Использовать go mod и установить зависимости Go
go generate

# Скомпилировать
go build -o server main.go (для Windows команда компиляции: go build -o server.exe main.go)

# Запустить бинарный файл
./server (для Windows команда запуска: server.exe)
```

### 2.2 Фронтенд

```bash
# Перейти в каталог проекта
cd web

# Установить зависимости
npm install

# Запустить в режиме разработки
npm run serve
```

### 2.3 Сервер

```bash
# Использование go.mod

# Установить модули Go
go list (go mod tidy)

# Скомпилировать сервер
go build
```

### 2.4 Автогенерация документации API с помощью Swagger

#### 2.4.1 Установка Swagger

##### (1) Использование VPN или за пределами материкового Китая
```bash
go get -u github.com/swaggo/swag/cmd/swag
```

##### (2) В материковом Китае

В материковом Китае доступ к go.org/x ограничен, рекомендуем использовать [goproxy.io](https://goproxy.io/zh/) или [goproxy.cn](https://goproxy.cn).

```bash
# Если вы используете Go версии 1.13–1.15, необходимо вручную включить GO111MODULE=on. Для версий Go 1.16 и новее этот шаг можно пропустить
# Шаг 1: Включить Go Modules
go env -w GO111MODULE=on 
# Шаг 2: Настроить переменную окружения GOPROXY
go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct

# Если вы хотите избежать лишних действий, используйте go generate для автоматического выполнения кода перед компиляцией. Однако это не работает в терминале командной строки `Goland` или `VSCode`
cd server
go generate -run "go env -w .*?"

# Установить swag с помощью следующей команды
go get -u github.com/swaggo/swag/cmd/swag
```

#### 2.4.2 Генерация документации API

```bash
cd server
swag init
```

> После выполнения приведенной выше команды в каталоге `server` появится папка `docs` с файлами `docs.go`, `swagger.json`, `swagger.yaml`. После запуска сервиса Go введите в браузере [http://localhost:8888/swagger/index.html](http://localhost:8888/swagger/index.html), чтобы просмотреть документацию Swagger.

## 3. Технологический стек

- Фронтенд: Используется [Element](https://github.com/ElemeFE/element) на основе [Vue](https://vuejs.org) для создания страниц.
- Бэкенд: Используется [Gin](https://gin-gonic.com/) для быстрого создания RESTful API. [Gin](https://gin-gonic.com/) — это веб-фреймворк, написанный на Go (Golang).
- База данных: `MySQL` (5.6.44), используется [GORM](http://gorm.io) для работы с данными, добавлена поддержка баз данных SQLite.
- Кэш: Используется `Redis` для хранения JWT-токенов активных пользователей и реализации ограничения множественного входа.
- API: Используется Swagger для автоматической генерации документации API.
- Конфигурация: Используются [fsnotify](https://github.com/fsnotify/fsnotify) и [viper](https://github.com/spf13/viper) для работы с файлами конфигурации в формате `yaml`.
- Логи: Используется [zap](https://github.com/uber-go/zap) для записи логов.

## 4. Архитектура проекта

### 4.1 Схема архитектуры

![Схема архитектуры](http://qmplusimg.henrongyi.top/gva/gin-vue-admin.png)

### 4.2 Подробная схема дизайна фронтенда (Автор: <a href="https://github.com/baobeisuper">baobeisuper</a>)

![Подробная схема дизайна фронтенда](http://qmplusimg.henrongyi.top/naotu.png)

### 4.3 Структура проекта

```
    ├── server
        ├── api             (вход для API)
        │   └── v1          (интерфейсы версии v1)
        ├── config          (пакет конфигурации)
        ├── core            (основные файлы)
        ├── docs            (каталог документации Swagger)
        ├── global          (глобальные объекты)                    
        ├── initialize      (инициализация)                        
        │   └── internal    (внутренние функции инициализации)                            
        ├── middleware      (слой промежуточного ПО)                        
        ├── model           (слой моделей)                    
        │   ├── request     (структуры входных параметров)                        
        │   └── response    (структуры выходных параметров)                            
        ├── packfile        (упаковка статических файлов)                        
        ├── resource        (папка статических ресурсов)                        
        │   ├── excel       (путь по умолчанию для импорта/экспорта Excel)                        
        │   ├── page        (генератор форм)                        
        │   └── template    (шаблоны)                            
        ├── router          (слой маршрутизации)                    
        ├── service         (слой сервисов)                    
        ├── source          (слой источников)                    
        └── utils           (набор утилит)                    
            ├── timer       (инкапсуляция интерфейса таймера)                        
            └── upload      (инкапсуляция интерфейса OSS)  
            
    └─web            (фронтенд)
        ├─public        (шаблоны развертывания)
        └─src           (исходный код)
            ├─api       (API фронтенда)
            ├─assets    (статические файлы)
            ├─components (компоненты)
            ├─router    (маршруты фронтенда)
            ├─store     (управление состоянием Vuex)
            ├─style     (общие стили)
            ├─utils     (общие утилиты фронтенда)
            └─view      (страницы)
```

## 5. Возможности

- Управление правами: Управление правами на основе `JWT` и `Casbin`.
- Загрузка и скачивание файлов: Реализация операций загрузки файлов на основе `Qiniuyun`, `Aliyun` и `Tencent Cloud` (для каждого сервиса необходимо самостоятельно получить `token` или `key`).
- Инкапсуляция пагинации: Фронтенд использует `mixins` для инкапсуляции пагинации, методы пагинации вызываются через `mixins`.
- Управление пользователями: Системный администратор назначает роли пользователей и их разрешения.
- Управление ролями: Создание основного объекта управления правами, назначение различных разрешений API и меню для роли.
- Управление меню: Реализация динамической конфигурации меню для пользователей, назначение различных меню разным ролям.
- Управление API: Разные пользователи могут вызывать разные разрешения API.
- Управление конфигурацией: Конфигурационный файл можно изменять через интерфейс (функция недоступна на сайте онлайн-демо).
- Условный поиск: Добавлен пример условного поиска.
- Пример RESTful: Примеры API можно увидеть в модуле управления пользователями.
  - Ссылка на файл фронтенда: [web/src/view/superAdmin/api/api.vue](https://github.com/flipped-aurora/gin-vue-admin/blob/master/web/src/view/superAdmin/api/api.vue).
  - Ссылка на бэкенд: [server/router/sys_api.go](https://github.com/flipped-aurora/gin-vue-admin/blob/master/server/router/sys_api.go).
- Ограничение множественного входа: Измените `user-multipoint` на `true` в разделе `system` файла `config.yaml` (необходимо самостоятельно настроить Redis и его параметры. Во время тестирования сообщайте о багах).
- Загрузка файлов по частям: Предоставлены примеры загрузки файлов и загрузки больших файлов по частям.
- Конструктор форм: Используется [@form-generator](https://github.com/JakHuang/form-generator).
- Генератор кода: Предоставляет бэкенду базовую логику и простой генератор кода для операций CRUD.

## 6. База знаний

### 6.1 Блог команды

> https://www.yuque.com/flipped-aurora
>
> В нашем блоге есть видеоуроки о фреймворке фронтенда. Если проект оказался полезным для вас, вы можете добавить меня в WeChat: shouzi_1994, ваши комментарии приветствуются.

### 6.2 Видеоуроки

(1) Курс по настройке среды разработки

> Bilibili: https://www.bilibili.com/video/BV1Fg4y187Bw/

(2) Курс по шаблонам

> Bilibili: https://www.bilibili.com/video/BV16K4y1r7BD/

(3) Введение в версию 2.0 и опыт разработки

> Bilibili: https://www.bilibili.com/video/BV1aV411d7Gm#reply2831798461

(4) Базовый курс по Golang

> https://space.bilibili.com/322210472/channel/detail?cid=108884

(5) Базовый курс по фреймворку Gin

> Bilibili: https://space.bilibili.com/322210472/channel/detail?cid=126418&ctype=0

(6) Видео с введением в обновления версии gin-vue-admin

> Bilibili: https://space.bilibili.com/322210472/channel/detail?cid=126418&ctype=0

## 7. Контакты

### 7.1 Группы

#### Группа QQ: 622360840

| Группа QQ |
|  :---:  |
| <img src="http://qmplusimg.henrongyi.top/qq.jpg" width="180"/> |

#### Группа WeChat: напишите в комментариях "加入gin-vue-admin交流群"

| WeChat |
|  :---:  |
| <img width="150" src="http://qmplusimg.henrongyi.top/qrjjz.png"> 

#### [О нас](https://www.gin-vue-admin.com/about/join.html)

## 8. Участники

Спасибо за ваш интерес к участию в gin-vue-admin!

<a href="https://openomy.app/github/flipped-aurora/gin-vue-admin" target="_blank" style="display: block; width: 100%;" align="center">
  <img src="https://openomy.app/svg?repo=flipped-aurora/gin-vue-admin&chart=bubble&latestMonth=3" target="_blank" alt="Contribution Leaderboard" style="display: block; width: 100%;" />
</a>

<a href="https://github.com/flipped-aurora/gin-vue-admin/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=flipped-aurora/gin-vue-admin" />
</a>

## 9. Поддержка

Если вы находите этот проект полезным, вы можете угостить автора стаканом сока :tropical_drink: [здесь](https://www.gin-vue-admin.com/coffee/index.html)

## 10. Коммерческое использование

Если вы используете этот проект в коммерческих целях, пожалуйста, соблюдайте лицензию Apache 2.0 и сохраняйте заявление о технической поддержке автора.

## 11. Перевод

Поддержка и консультация на рускком языке не производится. 

Переведено: [Vlaidslav Iaskin](https://t.me/xenon007)