# HoldMyLinkBot 🔗

**HoldMyLinkBot** — это минималистичный бот, который сохраняет ссылки, отправленные пользователем, и по запросу возвращает случайную из сохранённых.

Полезен для тех, кто любит сохранять статьи «на потом», но забывает их читать. Просто скинь ссылку — бот запомнит. Попроси — он напомнит.

---

## ⚙️ Возможности

- ✅ Сохраняет ссылки, отправленные пользователями
- 🎲 Отдаёт случайную сохранённую ссылку по запросу
- 🧱 Архитектура, удобная для расширения
- 🧼 Написан на Go без сторонних библиотек

---

## 🚀 Быстрый старт

### 1. Клонируй репозиторий
```bash
git clone https://github.com/yourusername/holdmylinkbot.git
cd holdmylinkbot
````

### 2. Скомпилируй и запусти

```bash
go build -o holdmylinkbot
./holdmylinkbot
```

> Убедись, что у тебя установлен Go (1.18+)

---

## 🧩 Архитектура

Проект построен по принципу разделения ответственности:

* **Хранилище ссылок** — обособлено, легко заменяется (в будущем можно подключить БД)
* **Обработка команд** — отделена от платформы
* **Интерфейс клиента (messaging client)** — позволяет подключать любые мессенджеры (Telegram, Discord, CLI и др.)

Добавление нового клиента требует только реализации интерфейса клиента. Вся остальная логика остаётся неизменной.

---

## 📦 Пример использования (CLI или чат)

```text
Пользователь: https://example.com/article

Бот: Ссылка сохранена!

Пользователь: /random

Бот: Попробуй почитать вот это: https://example.com/article
``
