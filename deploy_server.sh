sudo kill -9 $(ps -ef | grep "main.go" | awk '{print $2}')
sudo fuser -k 8080/tcp
cd /root/go/src/anaco-go-backend
git checkout master
git pull origin master
nohup go run main.go &