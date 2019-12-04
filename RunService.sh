#!/bin/bash
# 撰寫人員: Neil_Hsieh
# 撰寫日期：2019/01/14
# 說明： 啟動Golang的服務
#
# 備註：
#

# 執行GoFormat的目錄,
WORK_PATH=$(dirname $(readlink -f $0))
# 執行各容器，須掛載的資料夾位置
VOLUME_PATH=$(dirname $(readlink -f $0))/../
# 專案名稱(取當前資料夾路徑最後一個資料夾名稱)
PROJECT_NAME=${WORK_PATH##*/}
# Log存放的目錄(預設local路徑)
LOG="/var/log/app/$PROJECT_NAME"
# 讀取圖片路徑(預設dev路徑)
IMG="$VOLUME_PATH/images"

# 顯示環境
printf "\033[1;31m%s  %+10s\n" "(1).Dev" "(2).QATest"
printf "\033[1;31m%s  %s\n" "(3).Sit" "(4).Local"
printf "\033[1;31m%s  %+7s\n" "(5).T2" "(6).T4"

# 選擇環境
printf "\033[0;36m"
read -p "請選擇服務環境：" ENV_ID

# 執行環境
# 執行選項

printf "\033[37m"

case $ENV_ID in
    1)
        ENV="develop"
        LOG="/home/log/$PROJECT_NAME"
        ;;
    2)
        ENV="qatest"
        LOG="/home/log/$PROJECT_NAME"
        ;;
    3)
        ENV="sit"
        LOG="/home/log/$PROJECT_NAME"

        ;;
    4)
        ENV="local"
        IMG="./upload/images"

        go get github.com/pilu/fresh
        go mod download

        # 本機開發須安裝swagger + 初始化文件
        go get -u github.com/swaggo/swag/cmd/swag
        cd $WORK_PATH
        swag init

        ;;
    5)
        ENV="prodT2"
        ;;
    6)
        ENV="prodT4"
        ;;
    *)
        echo "格式不符合(環境變數)"
        exit
        ;;
esac

#############################
#############################
docker network ls | grep "web_service" >/dev/null 2>&1
    if  [ $? -ne 0 ]; then
        docker network create web_service
    fi

echo "ENV=$ENV">.env
echo "LOG=$LOG">>.env
echo "IMG=$IMG">>.env
echo "PROJECT_NAME=$PROJECT_NAME">>.env

# 啟動容器服務
if [ "$ENV" = "local" ]
then
    docker-compose up -d
else
    docker-compose up -d golang
fi