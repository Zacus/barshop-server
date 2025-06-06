#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import os
import sys
import subprocess
import shutil
import tarfile
from datetime import datetime
import logging

# 配置日志
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(levelname)s - %(message)s'
)
logger = logging.getLogger(__name__)

class PackageManager:
    def __init__(self):
        self.app_name = "barshop-server"
        self.version = self._get_version()
        self.build_time = datetime.utcnow().strftime('%Y-%m-%d_%H:%M:%S')
        
    def _get_version(self):
        """获取 git 版本信息"""
        try:
            result = subprocess.run(
                ['git', 'describe', '--tags', '--always', '--dirty'],
                capture_output=True,
                text=True,
                check=True
            )
            return result.stdout.strip()
        except subprocess.CalledProcessError as e:
            logger.error(f"获取版本信息失败: {e}")
            return "unknown"
                
    def clean(self):
        """清理构建文件"""
        logger.info("开始清理旧的构建文件...")
        paths = ['./dist']  # 只清理 dist 目录，不清理 bin 目录
        for path in paths:
            if os.path.exists(path):
                try:
                    shutil.rmtree(path)
                    logger.info(f"已删除: {path}")
                except Exception as e:
                    logger.error(f"删除 {path} 失败: {e}")
                    
    def create_release(self):
        """创建发布包"""
        logger.info("开始创建发布包...")
        
        # 检查必要文件是否存在
        if not os.path.exists('bin/' + self.app_name):
            logger.error(f"构建文件不存在: bin/{self.app_name}")
            return False
            
        # 创建目标目录结构
        dist_dirs = [
            'dist/bin',           # 可执行文件目录
            'dist/config',        # 配置文件目录
            'dist/scripts',       # 脚本目录
            'dist/logs',          # 日志目录
        ]
        
        for dir_path in dist_dirs:
            os.makedirs(dir_path, exist_ok=True)
            logger.info(f"创建目录: {dir_path}")
        
        # 定义需要复制的文件
        files_to_copy = {
            # 二进制文件
            f'bin/{self.app_name}': f'dist/bin/{self.app_name}',
            
            # 配置文件
            'config.yaml': 'dist/config.yaml',
            '.env.example': 'dist/config/.env.example',
            
            # 部署相关
            'Dockerfile': 'dist/Dockerfile',
            'scripts/deploy.sh': 'dist/scripts/deploy.sh',
            'scripts/env_checker.py': 'dist/scripts/env_checker.py',
        }
        
        # 复制文件
        for src, dst in files_to_copy.items():
            try:
                if os.path.exists(src):
                    os.makedirs(os.path.dirname(dst), exist_ok=True)
                    if os.path.isdir(src):
                        shutil.copytree(src, dst, dirs_exist_ok=True)
                    else:
                        shutil.copy2(src, dst)
                    logger.info(f"已复制: {src} -> {dst}")
                else:
                    logger.warning(f"文件不存在，跳过: {src}")
            except Exception as e:
                logger.error(f"复制 {src} 失败: {e}")
                
        # 创建启动脚本
        self._create_startup_script()
        
        # 创建压缩包
        archive_name = f'dist/{self.app_name}-{self.version}.tar.gz'
        try:
            with tarfile.open(archive_name, 'w:gz') as tar:
                # 切换到 dist 目录
                original_dir = os.getcwd()
                os.chdir('dist')
                
                # 添加所有文件到压缩包
                for item in os.listdir('.'):
                    tar.add(item)
                
                # 恢复工作目录
                os.chdir(original_dir)
                
            logger.info(f"已创建压缩包: {archive_name}")
            return True
            
        except Exception as e:
            logger.error(f"创建压缩包失败: {e}")
            return False
            
    def _create_startup_script(self):
        """创建启动脚本"""
        startup_script = """#!/bin/bash

# 获取脚本所在目录的绝对路径
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
APP_HOME="$(dirname "$SCRIPT_DIR")"

# 设置工作目录
cd "$APP_HOME"

# 确保日志目录存在
mkdir -p logs

# 设置环境变量
export CONFIG_PATH="$APP_HOME/config/config.yaml"

# 启动应用
exec ./bin/barshop-server "$@"
"""
        
        startup_path = 'dist/scripts/start.sh'
        try:
            with open(startup_path, 'w') as f:
                f.write(startup_script)
            os.chmod(startup_path, 0o755)  # 设置可执行权限
            logger.info(f"已创建启动脚本: {startup_path}")
        except Exception as e:
            logger.error(f"创建启动脚本失败: {e}")

def main():
    """主函数"""
    if len(sys.argv) < 2:
        print("Usage: python package.py {clean|release}")
        sys.exit(1)
        
    pm = PackageManager()
    command = sys.argv[1]
    
    if command == "clean":
        pm.clean()
    elif command == "release":
        if not pm.create_release():
            sys.exit(1)
    else:
        print(f"未知命令: {command}")
        sys.exit(1)

if __name__ == "__main__":
    main() 