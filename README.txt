# merge method - not yet.
jenkins + sonarsource
gitlab + sonarsource

# sonarsource module install - installed.
Traditional Chinese Language Pack
YAML Analyzer
SonarPythonLanguages
PMD
Java I18n Rules
GitLab Auth
Checkstyle
Ansible Lint

# gitlab sonarqube's Impersonation Token - api only, verified okay.
5LH3unBdMu5pytrvimbv

# gitlab -> 管理區塊 -> Applications -> create sonarqube
In your GitLab profile, you need to create a Developer Application for which the 'Authorization callback URL' must be set to '<value_of_sonar.core.serverBaseURL_property>/oauth2/callback/gitlab'.
--> http://192.168.100.190/oauth2/callback/gitlab
enable -> Trusted , api

# local test
application id:ffb8c08a70c5dce979801944ebc75ad13b8dd6c20645c3f75065e309e7dbd695
application secret:ff4a449012e1ad491aa943194be0d88ba169bcd282e055db83cd2aad5213900x

# production setting
application id:508ba07fdd8c40abccc5160c96c4171f2522072289c4dfe4dfe8054f33b8976c
application secret:cdfc9d43e49ac33bc3c68b9ca2c42bddca6311bee2299afb3ac70d631eb54a2x

# 配置sonar gitlab-plugin
admin 登录 SonarQube，点击 配置 —> 通用配置 —> GitLab
设置 -> GitLab url , Application ID , Secret
Gitlab access scope -> api
Default groups -> devops
Synchronize user groups -> enable

# local test
GitLab url: http://192.168.100.190:8080
Application ID: 

# production setting
GitLab url: http://gitlab.funpodium.net
Application ID: 508ba07fdd8c40abccc5160c96c4171f2522072289c4dfe4dfe8054f33b8976x

# scanner verified okay. (without /opt/sonar-scanner/conf/sonar-scanner.properties, get SONAR_TOKEN from sonarqube web gui.)
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


# https://docs.gitlab.com/ee/ci/variables/
variables:
  TEST: "HELLO WORLD"
