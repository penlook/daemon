CURRENT=`pwd`
USER=`whoami`
ROOT="/home/$USER/src/github.com"
mkdir -p $ROOT/penlook/daemon
cp -rf ./* $ROOT/penlook/daemon
cd $ROOT/penlook/daemon
go build
cd $CURRENT
