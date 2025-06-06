#!/bin/bash

set -e

# 定义变量
APP_NAME="barshop-server"
APP_PORT=8080
DEPLOY_USER="deploy"
DEPLOY_PATH="/opt/barshop"
BACKUP_PATH="/opt/barshop/backups"
VERSION=$(git describe --tags --always --dirty)

# 检查 Python 环境
check_python() {
    if ! command -v python3 &> /dev/null; then
        echo "❌ Python3 未安装"
        exit 1
    fi
    
    # 安装依赖
    if [ ! -f "scripts/requirements.txt" ]; then
        echo "❌ scripts/requirements.txt 不存在"
        exit 1
    fi
    
    python3 -m pip install -r scripts/requirements.txt
}

# 检查环境变量
check_env() {
    if [ -z "$DEPLOY_HOST" ]; then
        echo "❌ Error: DEPLOY_HOST environment variable is not set"
        exit 1
    fi
    
    # 运行 Python 环境检查脚本
    if ! python3 scripts/env_checker.py; then
        echo "❌ 环境检查失败"
        exit 1
    fi
}

# 创建远程目录结构
setup_remote_dirs() {
    echo "📁 Setting up remote directories..."
    ssh $DEPLOY_USER@$DEPLOY_HOST "mkdir -p $DEPLOY_PATH/{bin,config,logs,backups}"
}

# 备份当前版本
backup_current() {
    echo "💾 Backing up current version..."
    ssh $DEPLOY_USER@$DEPLOY_HOST "if [ -f $DEPLOY_PATH/bin/$APP_NAME ]; then \
        cp $DEPLOY_PATH/bin/$APP_NAME $BACKUP_PATH/${APP_NAME}_$(date +%Y%m%d_%H%M%S); \
    fi"
}

# 部署新版本
deploy_new() {
    echo "🚀 Deploying new version..."
    
    # 复制文件
    scp dist/$APP_NAME $DEPLOY_USER@$DEPLOY_HOST:$DEPLOY_PATH/bin/
    scp config.yaml $DEPLOY_USER@$DEPLOY_HOST:$DEPLOY_PATH/config/
    
    # 设置权限
    ssh $DEPLOY_USER@$DEPLOY_HOST "chmod +x $DEPLOY_PATH/bin/$APP_NAME"
}

# 重启服务
restart_service() {
    echo "🔄 Restarting service..."
    ssh $DEPLOY_USER@$DEPLOY_HOST "sudo systemctl restart $APP_NAME"
}

# 检查服务状态
check_service() {
    echo "🔍 Checking service status..."
    ssh $DEPLOY_USER@$DEPLOY_HOST "sudo systemctl status $APP_NAME"
    
    # 等待服务启动
    sleep 5
    
    # 检查端口是否正常监听
    if ssh $DEPLOY_USER@$DEPLOY_HOST "nc -z localhost $APP_PORT"; then
        echo "✅ Service is running and listening on port $APP_PORT"
    else
        echo "❌ Service is not responding on port $APP_PORT"
        exit 1
    fi
}

# Docker 部署
deploy_docker() {
    echo "🐳 Deploying Docker container..."
    
    # 停止并删除旧容器
    ssh $DEPLOY_USER@$DEPLOY_HOST "docker stop $APP_NAME || true"
    ssh $DEPLOY_USER@$DEPLOY_HOST "docker rm $APP_NAME || true"
    
    # 运行新容器
    ssh $DEPLOY_USER@$DEPLOY_HOST "docker run -d \
        --name $APP_NAME \
        --restart unless-stopped \
        -p $APP_PORT:$APP_PORT \
        -v $DEPLOY_PATH/config:/app/config \
        -v $DEPLOY_PATH/logs:/app/logs \
        $APP_NAME:$VERSION"
}

# 主函数
main() {
    check_python
    
    case "$1" in
        "setup")
            check_env
            setup_remote_dirs
            ;;
        "deploy")
            check_env
            backup_current
            deploy_new
            restart_service
            check_service
            ;;
        "docker")
            check_env
            deploy_docker
            check_service
            ;;
        *)
            echo "Usage: $0 {setup|deploy|docker}"
            exit 1
            ;;
    esac
}

# 执行主函数
main "$@"