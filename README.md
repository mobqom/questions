# questions
## CI/CD Процесс с Tekton в k3s

Для реализации полноценного цикла CI/CD с разделением на среды `dev` и `prod`, в проекте настроены манифесты Tekton.

### Структура манифестов
Все манифесты находятся в директории `k8s/`:
- `k8s/namespaces.yaml`: Пространства имен `dev`, `prod` и `tekton-pipelines`.
- `k8s/app/`: Деплоймент и Сервис приложения.
- `k8s/tekton/`:
    - `kaniko-task.yaml`: Сборка Docker-образа (без Docker-демона).
    - `deploy-task.yaml`: Применение манифестов (`kubectl apply`).
    - `pipeline.yaml`: Конвейер (Clone -> Build -> Deploy).
    - `pipelinerun-dev.yaml`: Пример запуска для среды `dev`.

### Установка Tekton и Dashboard

1. **Установка Tekton Pipelines:**
   ```bash
   kubectl apply -f https://storage.googleapis.com/tekton-releases/pipeline/previous/v0.44.0/release.yaml
   ```

2. **Установка Tekton Dashboard:**
   ```bash
   kubectl apply -f https://storage.googleapis.com/tekton-releases/dashboard/latest/release.yaml
   ```

3. **Доступ к Dashboard:**
   Дашборд настроен как `NodePort` сервис. Вы можете найти порт командой:
   ```bash
   kubectl get svc tekton-dashboard -n tekton-pipelines
   ```
   Доступ будет по адресу `http://192.168.64.2:<NODE_PORT>`.

### Запуск CI/CD пайплайна

Для запуска деплоя в `dev` среду:
```bash
kubectl create -f k8s/tekton/pipelinerun-dev.yaml
```

### Автоматизация деплоя по пушу (Tekton Triggers)

Для того чтобы деплой в `dev` происходил при пуше в ветку `dev`, а в `prod` — при пуше в `main`, используются **Tekton Triggers**.

#### 1. Установка Triggers
```bash
kubectl apply -f https://storage.googleapis.com/tekton-releases/triggers/previous/v0.24.0/release.yaml
kubectl apply -f https://storage.googleapis.com/tekton-releases/triggers/previous/v0.24.0/interceptors.yaml
```

#### 2. Применение конфигурации триггеров
```bash
kubectl apply -f k8s/tekton/triggers/template.yaml
kubectl apply -f k8s/tekton/triggers/triggers.yaml
kubectl apply -f k8s/tekton/triggers/event-listener.yaml
```

#### 3. Настройка Webhook
EventListener создает сервис. Для доступа извне переведите его в NodePort:
```bash
kubectl patch svc el-app-event-listener -n tekton-pipelines -p '{"spec": {"type": "NodePort"}}'
```
Узнайте порт:
```bash
kubectl get svc el-app-event-listener -n tekton-pipelines
```
Настройте Payload URL в GitHub репозитории: `http://192.168.64.2:<NODE_PORT>`.

### Изоляция сред
- **Namespace-based isolation**: Разные пространства имен (`dev` / `prod`) обеспечивают логическую изоляцию ресурсов.
- **Pipeline Parameterization**: Один и тот же `Pipeline` используется для обеих сред через параметр `TARGET_NAMESPACE`.
