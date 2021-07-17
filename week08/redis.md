
redis server 和 client 同在一台机器上测试
```root@ubuntu:~/redis/redis-6.2.4/src# ./redis-benchmark -h 127.0.0.1 -p 6379 -t get,set -d 10 -n 10000 -q
SET: 90090.09 requests per second, p50=0.247 msec
GET: 79365.08 requests per second, p50=0.239 msec

root@ubuntu:~/redis/redis-6.2.4/src# ./redis-benchmark -h 127.0.0.1 -p 6379 -t get,set -d 20 -n 10000 -q
SET: 58823.53 requests per second, p50=0.263 msec
GET: 86956.52 requests per second, p50=0.231 msec

root@ubuntu:~/redis/redis-6.2.4/src# ./redis-benchmark -h 127.0.0.1 -p 6379 -t get,set -d 50 -n 10000 -q
SET: 63291.14 requests per second, p50=0.239 msec
GET: 84033.61 requests per second, p50=0.239 msec

root@ubuntu:~/redis/redis-6.2.4/src# ./redis-benchmark -h 127.0.0.1 -p 6379 -t get,set -d 50 -n 10000 -q
SET: 85470.09 requests per second, p50=0.255 msec
GET: 98039.22 requests per second, p50=0.231 msec

root@ubuntu:~/redis/redis-6.2.4/src#
root@ubuntu:~/redis/redis-6.2.4/src# ./redis-benchmark -h 127.0.0.1 -p 6379 -t get,set -d 100 -n 10000 -q
SET: 88495.58 requests per second, p50=0.247 msec
GET: 100000.00 requests per second, p50=0.231 msec

root@ubuntu:~/redis/redis-6.2.4/src# ./redis-benchmark -h 127.0.0.1 -p 6379 -t get,set -d 200 -n 10000 -q
SET: 54945.05 requests per second, p50=0.263 msec
GET: 84033.61 requests per second, p50=0.255 msec

root@ubuntu:~/redis/redis-6.2.4/src# ./redis-benchmark -h 127.0.0.1 -p 6379 -t get,set -d 1000 -n 10000 -q
SET: 74626.87 requests per second, p50=0.271 msec
GET: 90090.09 requests per second, p50=0.255 msec

root@ubuntu:~/redis/redis-6.2.4/src# ./redis-benchmark -h 127.0.0.1 -p 6379 -t get,set -d 5000 -n 10000 -q
SET: 76923.08 requests per second, p50=0.295 msec
GET: 80000.00 requests per second, p50=0.271 msec

```
