
#### 作业
1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
2、写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。


1. redis server 和 client 同在一台机器上测试
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
待续：
