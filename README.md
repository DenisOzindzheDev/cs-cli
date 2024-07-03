# VAULT-AUTO EXTRACTOR

## Пререквезиты

:warning: У вас должен быть кубконфиг с доступами на чтение секретов в кластере где расположен канал, кубконфиги читаются из ./.kubeconfig

:warning: Установленный go не ниже 1.22.4

## Установка 

1. `make build` - скомпилирует приложение
2. `mv <приложение(cs-cli)> /Users/user/go/bin` - user заменить на свой и путь заменить на диррективу где установлен go
3. `export PATH=/Users/user/go/bin:$PATH` добавляем переменную окружения чтобы запускать исполняемый файл
4. в `.zshrc` нужно тоже прописать пункт 3 

## Использование

```
Usage:
  cs-cli vault [flags]

Flags:
  -n, --channel-namespace string   channel namespace в кубернетесе
  -p, --data-path string           vault path путь к секретам 
  -h, --help                       help for vault
  -v, --vault-namespace string     vault namespace(optional) если есть неймспейс указать
```

пример 
`cs-cli vault -n apim-channel-prod -p platformeco/data/apim-channel-prod/auth`

## TODO

1. Exec KUBECONFIG from $KUBECONFIG environment variable [ * ]
2. Move vault host to config
3. Move vault secret name to config
4. Handle connection to clusters and don't crash while it not works [ * ]
5. Pretify and refactor code