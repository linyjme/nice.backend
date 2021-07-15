# Raysync web后端
- RESTful API service for raysync-client, powered by golang + gin, required sqlite

## Operation

### 接口文档

    Operation/static/docs

### postman api

    docs/raysync_backend.postman_collection.json
    
### 启动 调试环境
```shell script
python manage.py runserver -h 0.0.0.0 -p 9090 
```

### 正式环境
```shell script
gunicorn -w 4 -b 0.0.0.0:5000 manage:app
```

### 打包成exe
```shell script
pyinstaller -F raysync-web.py
```

### 生成在线api docs
python manage.py apidoc

### 查看在线api
go mod init project
```
