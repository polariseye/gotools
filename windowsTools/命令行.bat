echo off
title golang命令行工具
echo 当前盘符：%~d0
echo 当前盘符和路径：%~dp0
echo 当前盘符和路径的短文件名格式：%~sdp0
echo 当前批处理全路径：%~f0
echo 当前CMD默认目录：%cd%

echo 开始设置gopath环境变量
set gopath=%~sdp0
echo on
cmd.exe