#!/bin/bash

set -e

# 定义变量
APP_NAME="barshop-server"
VERSION=$(git describe --tags --always --dirty)
BUILD_TIME=$(date -u '+%Y-%m-%d_%H:%M:%S')
DOCKER_REGISTRY="your-registry.com"  # 替换为实际的 Docker 仓库地址

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

# 检查 main 包位置
check_main_location() {
    local main_locations=(
        "."                # 根目录
        "./cmd/server"     # 标准 Go 项目布局
        "./cmd/api"       # 替代布局
        "./main"           # 简单布局
    )
    
    for location in "${main_locations[@]}"; do
        if [ -f "$location/main.go" ]; then
            echo "$location"
            return 0
        fi
    done
    
    echo "❌ 找不到 main.go 文件"
    exit 1
}

# 本地构建
build_local() {
    echo "🔨 Building locally..."
    
    # 检查并获取 main 包位置
    MAIN_PATH=$(check_main_location)
    echo "📍 找到 main 包位置: $MAIN_PATH"
    
    # 创建构建目录
    mkdir -p bin
    
    # 构建应用
    echo "🔧 开始构建..."
    go build -o ./bin/$APP_NAME \
        -ldflags "-X main.Version=$VERSION -X main.BuildTime=$BUILD_TIME" \
        "$MAIN_PATH"
        
    echo "✅ Local build complete: ./bin/$APP_NAME"
}

# Docker 构建
build_docker() {
    echo "🐳 Building Docker image..."
    docker build -t $APP_NAME:$VERSION \
        --build-arg VERSION=$VERSION \
        --build-arg BUILD_TIME=$BUILD_TIME \
        .
    
    # 标记最新版本
    docker tag $APP_NAME:$VERSION $APP_NAME:latest
    
    echo "✅ Docker build complete"
}

# 主函数
main() {
    check_python
    
    case "$1" in
        "clean")
            python3 scripts/package.py clean
            ;;
        "local")
            python3 scripts/package.py clean
            build_local
            ;;
        "docker")
            python3 scripts/package.py clean
            build_docker
            ;;
        "release")
            # 先清理发布目录
            python3 scripts/package.py clean
            
            # 如果没有构建文件，先构建
            if [ ! -f "./bin/$APP_NAME" ]; then
                build_local
            fi
            
            # 创建发布包
            python3 scripts/package.py release
            ;;
        *)
            echo "Usage: $0 {clean|local|docker|release}"
            exit 1
            ;;
    esac
}

# 执行主函数
main "$@"