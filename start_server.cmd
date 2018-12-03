echo "------Starting Authentication Server------"
start go run cmd/main/auth.go
echo "------Starting Storage Server------"
start go run cmd/main/storage.go
echo "------Starting Web Server------"
start go run cmd/main/web.go