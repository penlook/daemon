CURRENT=`pwd`
USER=`whoami`
ROOT="/home/$USER/src/github.com"
sudo mkdir -p $ROOT/penlook/service/component/daemon
sudo cp -rf ./* $ROOT/penlook/service/component/daemon
cd $ROOT/penlook/service/component/daemon
go build
cd $CURRENT
