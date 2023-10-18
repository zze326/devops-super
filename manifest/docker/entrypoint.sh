#!/bin/bash

# 检查并设置默认值
if [ -z "${SERVER_PORT}" ]; then
  echo "环境变量 SERVER_PORT 未设置，将使用默认值 8001"
  SERVER_PORT=8001
fi


if [ -z "${DB_HOST}" ]; then
  echo "环境变量 DB_HOST 未设置" && exit 1
fi

if [ -z "${DB_PORT}" ]; then
  DB_PORT="3306"
fi

if [ -z "${DB_NAME}" ]; then
  DB_NAME="devops_super"
fi

if [ -z "${DB_USER}" ]; then
  echo "环境变量 DB_USER 未设置" && exit 1
fi

if [ -z "${DB_PWD}" ]; then
  echo "环境变量 DB_PWD 未设置" && exit 1
fi

if [ -z "${JWT_SECRET}" ]; then
  echo "环境变量 JWT_SECRET 未设置，将使用默认值"
  JWT_SECRET="vIIEngfamdsaGZasdsasdasadkseadgF9fe"
fi

# 渲染配置文件模板
sed "s/{{SERVER_PORT}}/$SERVER_PORT/g; s/{{DB_HOST}}/$DB_HOST/g; s/{{DB_PORT}}/$DB_PORT/g; s/{{DB_NAME}}/$DB_NAME/g; s/{{DB_USER}}/$DB_USER/g; s/{{DB_PWD}}/$DB_PWD/g; s/{{JWT_SECRET}}/$JWT_SECRET/g;" config/config-tpl.yaml > config/config.yaml
exec ./main --gf.gcfg.file=config.yaml