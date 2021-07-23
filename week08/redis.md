
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
#### 
1. 生成value是10字节大小， 数量2w
test_10.py
```

#!/usr/bin/python
for i in range(20000):
    print 'set name'+str(i),'0123456789'
 ```
 ```
python test_10.py > test_10.cmd
cat test_10.cmd|redis-cli --pipe
```
result:
```
 ./redis-cli
127.0.0.1:6379> info memory
# Memory
used_memory:2349152
used_memory_human:2.24M
used_memory_rss:6963200
used_memory_rss_human:6.64M
used_memory_peak:2408992
used_memory_peak_human:2.30M
used_memory_peak_perc:97.52%
used_memory_overhead:1663288
used_memory_startup:809832
used_memory_dataset:685864
used_memory_dataset_perc:44.56%
allocator_allocated:2399848
allocator_active:2666496
allocator_resident:5259264
total_system_memory:2096828416
total_system_memory_human:1.95G
```
used_memory_dataset:685864, 685864/20000= 34.2932

2. 生成value是20字节大小， 数量2w
  test_20.py
 ```
#!/usr/bin/python
for i in range(20000):
    print 'set name'+str(i),'01234567890123456789'
 ```
```
python test_20.py > test_20.cmd
cat test_20.cmd|redis-cli --pipe
```
result:
```
127.0.0.1:6379> info memory
# Memory
used_memory:2733160
used_memory_human:2.61M
used_memory_rss:7356416
used_memory_rss_human:7.02M
used_memory_peak:2851856
used_memory_peak_human:2.72M
used_memory_peak_perc:95.84%
used_memory_overhead:1892472
used_memory_startup:809832
used_memory_dataset:840688
used_memory_dataset_perc:43.71%
allocator_allocated:2781768
allocator_active:3031040
allocator_resident:5607424
total_system_memory:2096828416
total_system_memory_human:1.95G
```
used_memory_dataset:840688, 840688/20000 = 42.0344

