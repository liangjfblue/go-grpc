# CA
> ca的作用是保证证书的有效性和可靠性，与ca颁发有关的是根证书

## 根证书
- 公钥
- 密钥

根证书是属于根证书颁发机构（CA）的公钥证书。通过验证 CA 的签名从而信任 CA ，任何人都可以得到 CA 的证书（含公钥），用来验证它所签发的证书（客户端、服务端）

## 生成 Key
`openssl genrsa -out ca.key 2048`

## 生成密钥
`openssl req -new -x509 -days 7200 -key ca.key -out ca.pem`

## 填写信息
```
Country Name (2 letter code) []:
State or Province Name (full name) []:
Locality Name (eg, city) []:
Organization Name (eg, company) []:
Organizational Unit Name (eg, section) []:
Common Name (eg, fully qualified host name) []:123.56.157.144
Email Address []:
```

# Server
`openssl ecparam -genkey -name secp384r1 -out server.key`

## 生成 CSR
CSR 是 证书请求文件。作用是 CA 利用 CSR 文件进行签名使得攻击者无法伪装或篡改原有证书

`openssl req -new -key server.key -out server.csr`

## 基于 CA 签发
`openssl x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in server.csr -out server.pem`

# Client
##生成 Key
`openssl ecparam -genkey -name secp384r1 -out client.key`

## 生成 CSR
`openssl req -new -key client.key -out client.csr`

## 基于 CA 签发
`openssl x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in client.csr -out client.pem`



error:
> panic: rpc error: code = Unavailable desc = all SubConns are in TransientFailure, latest connection error: connection closed

> panic: rpc error: code = Unavailable desc = all SubConns are in TransientFailure, latest connection error: connection error: desc = "transport: authentication handshake failed: x509: certificate is valid for 123.56.157.144, not go_protoc.HelloServer

