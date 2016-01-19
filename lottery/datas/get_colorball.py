#coding=utf-8

import urllib2, os, re

def get_ssq_data():
    try:
        ssq_url = urllib2.urlopen('http://www.zhcw.com/data-js/nowdata50.js')
        f = open('./ssq_data.txt', 'w')
        data = ssq_url.read()
        print data
        f.write(data)
        f.close()
        ssq_url.close()

    except Exception, e:
        print "get ssq data err ", e
        return


get_ssq_data()
