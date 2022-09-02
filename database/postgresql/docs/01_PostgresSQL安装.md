## Windows
```bash
docker run --name postgres \
    -e POSTGRES_PASSWORD=123456 \
    -p 5432:5432 \
    -v D:\\DevData\\postgres\\data:/var/lib/postgresql/data \
    -d postgres
```

## Linux

```bash
docker run --name postgres \
    -e POSTGRES_PASSWORD=123456 \
    -p 5432:5432 \
    -v /root/docker/postgresql/data:/var/lib/postgresql/data \
    -d postgres
```

参考:
- https://blog.csdn.net/qq_44732146/article/details/124795972