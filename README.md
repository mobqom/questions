# questions
## CICD Process

Проект использует комбинированный CICD процесс:
1. **GitHub Actions**: Используется для первичной проверки кода (build, lint) и сборки Docker образа.
2. **Tekton**: Используется для непрерывного деплоя (CD) внутри Kubernetes кластера.

### Tekton Setup

Манифесты для Tekton находятся в директории `tekton/`.

Для развертывания Tekton ресурсов в кластере:
```bash
kubectl apply -f tekton/tasks/
kubectl apply -f tekton/pipelines/
```

### GitHub Actions integration

Для того чтобы GitHub Actions мог триггерить Tekton пайплайны, необходимо:
1. Иметь установленный Tekton Triggers в кластере (опционально для webhook).
2. Настроить `KUBECONFIG` или `TOKEN` в GitHub Secrets для возможности запуска `PipelineRun`.

В текущей конфигурации `.github/workflows/ci.yml` содержит пример триггера `PipelineRun` через `kubectl`.
