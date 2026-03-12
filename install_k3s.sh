#!/bin/bash
set -e

# Установка k3s без traefik
curl -sfL https://get.k3s.io | INSTALL_K3S_EXEC="--disable traefik" sh -

# Подождать запуск
sleep 10

# Настройка прав для kubeconfig, чтобы его можно было прочитать по ssh без sudo
sudo chmod 644 /etc/rancher/k3s/k3s.yaml
