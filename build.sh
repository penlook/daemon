CURRENT=`pwd`
USER=`whoami`
ROOT="/home/$USER/src/github.com"
cp -rf ./* $ROOT/penlook/daemon
cd $ROOT/penlook/daemon
go build
cd $CURRENT
