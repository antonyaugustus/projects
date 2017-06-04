#!/usr/local/bin/env python

import yum
import os
import subprocess

def influx_repo():
    with open('/etc/yum.repos.d/influxdb.repo', 'w') as repo:
        repo.write('[influxdb]')
        repo.write('name = InfluxDB Repository - RHEL \$releasever')
        repo.write('baseurl = https://repos.influxdata.com/rhel/\$releasever/\$basearch/stable')
        repo.write('enabled = 1')
        repo.write('gpgcheck = 1')
        repo.write('gpgkey = https://repos.influxdata.com/influxdb.key')

def influx_install():
    yb = yum.YumBase()
    yb.conf.assumeyes = True

    #if yb.rpmdb.searchNevra(name=line.strip()):
    #    continue
    #else:
    #    print('this package is not installed {0}'.format(line))

    yb.install(name='influxdb')

def influx_start():
    cmd = 'systemctl start influxdb'
    p = subprocess.Popen(cmd, stderr=subprocess.PIPE)

    while True:
        out = p.stderr.read(1)
    if out == '' and p.poll() != None:
        break
    if out != '':
        sys.stdout.write(out)
        sys.stdout.flush()

def influx_variable():
    INFLUXDB_CONFIG_PATH="/etc/influxdb/influxdb.conf"
