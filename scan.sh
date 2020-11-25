#!/bin/bash

export SONARQUBE_URL="http://192.168.100.190"
export YOUR_REPO="/opt/allure-docker"
export YOUR_CACHE_DIR="sonar_cache"

export SONAR_TOKEN="f015ed7c02e4da7b47503ce96a29cb01ba59fc7d"
export SONAR_PROJECT="test2"
export SONAR_ENCODING="UTF-8"
export SONAR_PROJECT_BASE_DIR="/usr/src"

# run as a non-root container.
docker run \
    --rm \
    --user="$(id -u):$(id -g)" \
    -e SONAR_HOST_URL="${SONARQUBE_URL}"  \
    -e SONAR_SCANNER_OPTS="-Dsonar.projectKey=${SONAR_PROJECT} -Dsonar.sourceEncoding=${SONAR_ENCODING}" \
    -e SONAR_TOKEN="${SONAR_TOKEN}" \
    -e SONAR_PROJECT_BASE_DIR="${SONAR_PROJECT_BASE_DIR}" \
    -v "${YOUR_REPO}:${SONAR_PROJECT_BASE_DIR}" \
    -v "${YOUR_CACHE_DIR}:/opt/sonar-scanner/.sonar/cache" \
    sonarsource/sonar-scanner-cli

# backup for use.
#    -e SONAR_SCANNER_OPTS="-Dsonar.projectKey=test -Dsonar.sources=. -Dsonar.login=f015ed7c02e4da7b47503ce96a29cb01ba59fc7d" \
#    -e SONAR_SCANNER_DEBUG_OPTS=""
# --- END --- #
