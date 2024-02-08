# Шаблон пустого проекта на Golang

## Инициализация проекта

- Создайте каталог для нового проекта
- Скопируйте туда файлы из шаблона
- Запустить создание проекта:

```cmd
go mod init <name project>
go mod tidy
```

Далее все команды по компиляции, тестированию, выполнению и т.д. производяться четез make-файл

## Используемые библиотеки

- github.com/ilyakaznacheev/cleanenv  (конфигурирование)
- slog (логирование)

## Запуск с разными конфигами

./project -c app.env
./project -c config.yml