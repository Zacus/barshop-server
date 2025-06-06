#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import os
import sys
import socket
import time
import logging
from typing import List, Dict
import paramiko
from dotenv import load_dotenv

# 配置日志
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(levelname)s - %(message)s'
)
logger = logging.getLogger(__name__)

class EnvironmentChecker:
    def __init__(self, host: str, port: int, user: str):
        self.host = host
        self.port = port
        self.user = user
        self.ssh_client = None
        
    def __enter__(self):
        return self
        
    def __exit__(self, exc_type, exc_val, exc_tb):
        if self.ssh_client:
            self.ssh_client.close()
            
    def check_remote_connection(self, retries: int = 3) -> bool:
        """检查远程连接"""
        logger.info(f"正在检查与 {self.host} 的连接...")
        
        for attempt in range(retries):
            try:
                if not self.ssh_client:
                    self.ssh_client = paramiko.SSHClient()
                    self.ssh_client.set_missing_host_key_policy(paramiko.AutoAddPolicy())
                
                self.ssh_client.connect(
                    self.host,
                    username=self.user,
                    timeout=10
                )
                logger.info("SSH 连接成功")
                return True
                
            except Exception as e:
                logger.warning(f"连接尝试 {attempt + 1}/{retries} 失败: {e}")
                if attempt < retries - 1:
                    time.sleep(5)
                    
        logger.error(f"无法连接到远程主机 {self.host}")
        return False
            
    def check_service_health(self, retries: int = 3) -> bool:
        """检查服务健康状态"""
        logger.info(f"正在检查服务端口 {self.port}...")
        
        for attempt in range(retries):
            try:
                with socket.create_connection((self.host, self.port), timeout=5):
                    logger.info(f"服务端口 {self.port} 正常")
                    return True
            except Exception as e:
                logger.warning(f"端口检查尝试 {attempt + 1}/{retries} 失败: {e}")
                if attempt < retries - 1:
                    time.sleep(5)
                    
        logger.error(f"服务端口 {self.port} 不可访问")
        return False
            
    def check_disk_space(self, min_space_gb: float = 5.0) -> bool:
        """检查磁盘空间"""
        if not self.ssh_client:
            logger.error("SSH 客户端未连接")
            return False
            
        try:
            stdin, stdout, stderr = self.ssh_client.exec_command("df -h /")
            df_output = stdout.read().decode()
            
            # 解析 df 输出获取可用空间
            lines = df_output.strip().split('\n')
            if len(lines) >= 2:
                fields = lines[1].split()
                available = fields[3]
                if 'G' in available:
                    available_gb = float(available.replace('G', ''))
                    if available_gb < min_space_gb:
                        logger.error(f"可用空间不足: {available_gb}GB < {min_space_gb}GB")
                        return False
                    logger.info(f"磁盘空间充足: {available_gb}GB")
                    return True
                    
            logger.error("无法解析磁盘空间信息")
            return False
            
        except Exception as e:
            logger.error(f"检查磁盘空间失败: {e}")
            return False
            
    def verify_deployment(self) -> Dict[str, bool]:
        """验证部署环境"""
        results = {
            "ssh_connection": False,
            "service_health": False,
            "disk_space": False
        }
        
        # 检查 SSH 连接
        results["ssh_connection"] = self.check_remote_connection()
        if not results["ssh_connection"]:
            return results
            
        # 检查服务健康状态
        results["service_health"] = self.check_service_health()
        
        # 检查磁盘空间
        results["disk_space"] = self.check_disk_space()
        
        return results

def main():
    """主函数"""
    # 加载环境变量
    load_dotenv()
    
    # 获取必要的环境变量
    host = os.getenv("DEPLOY_HOST")
    port = int(os.getenv("APP_PORT", "8080"))
    user = os.getenv("DEPLOY_USER", "deploy")
    
    if not host:
        logger.error("未设置 DEPLOY_HOST 环境变量")
        sys.exit(1)
        
    # 执行环境检查
    with EnvironmentChecker(host, port, user) as checker:
        results = checker.verify_deployment()
        
        # 输出检查结果
        print("\n环境检查结果:")
        print("-" * 20)
        for check, status in results.items():
            print(f"{check}: {'✅ 通过' if status else '❌ 失败'}")
            
        # 如果有任何检查失败，返回非零状态码
        if not all(results.values()):
            sys.exit(1)

if __name__ == "__main__":
    main()