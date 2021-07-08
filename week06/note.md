#### 第六周， 评论系统学习记录

一. 对于读，流程：
   1.  commom_service: 先去redis cache 读
   2.  cache miss 后， singalflight 去读db，同时只有一个线程去读db.
   3.  读完db,   如果直接往redis 构建缓存，高并发的情况下，可能都卡在redis, 程序可能oom.
        + 所以 进程内存里 生成一个lru临时小缓存， 过期时间短，让kafka 完成异步构建就行， 小缓存，避免过多占用进程内存
        + 如果某个topic 热点，很大qps 也可能把redis 打死，这个时候，也可以动态滑动统计最近top N  的热点，在进程内存构建缓存，牺牲短时一致性
   4.  以评论topic 为key , singalfight 向kafka 投递指令；同时实现一个lru 指令缓存，避免3-5秒短时间内向kafka 投递相同的构建缓存指令。
   5.  + common_job 根据指令，先读redis，（因为判断redis key 是否存在很快）, miss 就读 db, 这也减少读db次数, 读完db 再构建redis 缓存。因为构建缓存有时需要一两秒，这时很可能同时收到多个实例相同的指令，那么common_job也可以singalflight 来构建缓存。
        + 为什么用kafka 和 common_job来异步构建缓存，避免redis 或者db 某些情况下性能抖动，导致进程挂掉或者db 被打死，加kafka, 就把压力放在kafka, 
         同时kafka 还可以削峰的作用，万一db 性能下降，也不会打死db, 因为 common_job 根据db 的性能线性消费kafka 的消息。

二、 写：写是每个用户发表评论，没办法用signalfight 方案。
   1. commom_service 直接 返回成功，并且 往kafka 发送消息。
   2. common_job 根据db 的性能不断消费消息，写到db,  然后马上构建redis缓存？如果写的并发大，会不会打死redis? 
           不会，common_job 是根据db和redis 的处理能力来消费kafka消息的，如果db 和redis 性能不行，只会在kafka积累消息，用户感觉评论延迟显示出来。
   3. 问题:  
     +  每次写db,都更新这topic 相关的所有缓存吗？ 更新缓存 content_index ，sortedset 排序。     
     +  mysql db 是读写分离的吗， 那么写完后，common_job 还要等slave 更新后，再读，还是直接在master db 读就行了。

三、  es :  运营多维度检索、 OLAP 
  1.  canal---》kafka ---> join 大表、转换昵称等--》es 多维度检索 <-------> common_admin -----> CRUD 操作 db、redis 。

四；总结
 + commom_service 归并回源读db, 归并回源 向kafka 投递相同指令，甚至实现lru 缓存缓存指令， 和lru 缓存redis数据，避免读db，  
+ common_job 预判的缓存是否存在来放弃构建缓存， 也能减少读db.
