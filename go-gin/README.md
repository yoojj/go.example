# gin

https://gin-gonic.com/



```bash 
# 프로젝트(모듈) 생성 
go mod init example-gin

# 모듈 설치 
go get -u github.com/gin-gonic/gin
go get -u github.com/go-ini/ini

# 실행 
go run main.go

# 빌드
go build 

# vscode task  
# ctrl + shift + b 


# + auto reloading
npm i -g nodemon
nodemon -v 

## nodemon.json 설정 후 실행 
nodemon
```



**기본 구조**
```
example-gin
├─ config 
├─ middleware 
├─ mvc 
├─ router 
└─ main.go 
```