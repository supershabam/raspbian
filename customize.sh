#!/bin/bash

set -Eeuo pipefail
# echo commands as they are being run
set -x

site_name=$1
site_port=$2

cp /app/base.img /image.img
mount -o loop,offset=50331648 /image.img /mnt

cp /app/files/authorized_keys /mnt/home/pi/.ssh/authorized_keys

sed -i "s/\$site-name/$site_name/" /mnt/etc/frpc-brooklyn.ini
sed -i "s/\$site-name/$site_name/" /mnt/etc/frpc-london.ini
sed -i "s/\$site-port/$site_port/" /mnt/etc/frpc-brooklyn.ini
sed -i "s/\$site-port/$site_port/" /mnt/etc/frpc-london.ini

mv /image.img /app/image.img
