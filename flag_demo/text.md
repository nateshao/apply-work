flag有两种方式：

1、flag.Type，其中Type可以是：int、string、bool，float等类型，返回指针类型。

var port = flag.Int("port", 0, "相关描述")
参数1：flag的名称

参数2：flag的值，上例中默认值是0

参数3：flag的描述

2、flag.TypeVar，将类型绑定到一个变量上。

var port int
flag.IntVar(&port, "port", 0, "相关描述")
参数1：flag的值

参数2：flag的名称

参数3：flag的值，上例中默认值是0

参数4：flag的描述
