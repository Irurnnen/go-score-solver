# go-score-solver

### Переход между репозиториями проетка:

- [go-score-gui](https://github.com/Irurnnen/go-score-gui)
- [go-score-orchestra](https://github.com/Irurnnen/go-score-orchestra)
- [go-score-db](https://github.com/Irurnnen/go-score-db)
- Вы здесь: [go-score-solver](https://github.com/Irurnnen/go-score-solver)

### Запуск:

1. Создать файл `.env` в корне репозитория по примеру:

    ```env
    PORT=6874 # Порт на котором будет работать программа
    ```

2. Создать файл `score-solver_config.yaml` по директории `../.secrets/` по примеру:

    ```yaml
    port: 6874
    host: 0.0.0.0
    release: release
    ```
3. Запустить docker командой `docker compose up --detach`

### Остановка:

1. Остановить docker контейнер с помощью compose: команда `docker compose down`

2. Остановить docker контейнер: команда `docker stop lpz-solver` 