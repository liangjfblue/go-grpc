# 证书生成
## 私钥
`openssl ecparam -genkey -name secp384r1 -out server.key`

## 自签公钥
`openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650`

## 填写信息
```
Country Name (2 letter code) [AU]:11
State or Province Name (full name) [Some-State]:11
Locality Name (eg, city) []:11
Organization Name (eg, company) [Internet Widgits Pty Ltd]:11
Organizational Unit Name (eg, section) []:11
Common Name (e.g. server FQDN or YOUR name) []:123.56.157.144
Email Address []:123

```

#注意
Common Name (e.g. server FQDN or YOUR name) []:

这里必须填写：`123.56.157.144`。不然会出现一下报错：

`panic: rpc error: code = Unavailable desc = all SubConns are in TransientFailure, latest connection error: connection error: desc = "transport: authentication handshake failed: x509: certificate is valid for 123.56.157.144, not 127.0.0.1"`

这里是证书验证服务器IP，查了下时阿里云的：

IP地址: 123.56.157.144北京市大兴区 阿里云

