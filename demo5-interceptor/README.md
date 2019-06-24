## 拦截器
在rpc调用前后做一些特别的处理。

## 两种拦截器：
- 普通方法：一元拦截器（grpc.UnaryInterceptor）
- 流方法：流拦截器（grpc.StreamInterceptor）

一元拦截器的调用方法：

`grpc.UnaryInterceptor()`

从上面的方法可以看到，参数是 UnaryServerInterceptor：

`type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)`

所以，要调用一元拦截器，只要实现以下的方法，并在 grpc.NewServer()传入就ok了。

详细的实现看源码。

> 由于grpc只支持设置一个拦截器，可以采用开源项目 go-grpc-middleware实现设置多个拦截器










