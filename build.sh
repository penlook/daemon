CURRENT=`pwd`
USER=`whoami`
ROOT="/home/$USER/src/github.com"
cp -rf ./* $ROOT/penlook/service/component/daemon
cd $ROOT/penlook/service/component/daemon
go build
cd $CURRENT
