# find top-k in time sliding window

# 查找滑动时间窗口频繁项

### 算法：

1. sketchWindow 记录n秒内的top k个频繁线
2. sketchWindow满时生成摘要添加到下一个window窗口
3. window窗口如果不是最后一个窗口，判断当前窗口是否已满，如果满生成摘要添加到下一个窗口并清空当前窗口，如果不满进行滑动计算
4. 滑动计算，将滑出的摘要做减法，滑入的做加法
5. window窗口是最后一个窗口时进行滑动计算不生成摘要

具体实现需要考虑时间跨越检测周期的情况

### 数据结构

* minheap 最小堆
* circular_queue 循环队列
* countMinSketch 


### 测出用例设计
 
