#!/bin/sh

pushd `dirname $0` > /dev/null

BASEDIR=`pwd`

nohup $BASEDIR/ims -log_dir=/data/logs/ims ims.cfg >/data/logs/ims/ims.log 2>&1 &

nohup $BASEDIR/imr -log_dir=/data/logs/imr imr.cfg >/data/logs/imr/imr.log 2>&1 &

nohup $BASEDIR/im -log_dir=/data/logs/im im.cfg >/data/logs/im/im.log 2>&1 &
