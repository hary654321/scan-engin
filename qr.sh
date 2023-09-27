#!/bin/bash
start_date="20230601";
end_date="20230701";

 ##通过循环，返回日期值（包含开始和结束日期，闭区间）
 for i in `seq 0 100000`
 do
     t_date=`date -d "${start_date} +$(($i+1)) day" "+%Y_%m_%d"`
     echo $t_date
     cnt_days=$i
     

    ps -ef | grep qr | grep -v grep | awk '{print $2}' | xargs kill -9
    cd /u4/logstashqr

    yes|cp -f config/qr.base config/qr.conf
    sed -i "s/day/$t_date/g" config/qr.conf
    rm -rf data
    ./bin/logstash -f config/qr.conf


    ##如果循环到当天，就退出
     if [ $t_date == $end_date ]
     then
         break
     fi
 done

