rpc-go:
	@echo "输入app名称"
	@read app; \
	zeroctl rpc protoc apps/$$app/rpc/$$app.proto --go_out=apps/$$app/rpc --go-grpc_out=apps/$$app/rpc --zrpc_out=apps/$$app/rpc;

api-go:
	@echo "输入app与api名称,以空格隔开"
	@read app api; \
	zeroctl api go -api apps/$$app/$$api/$$api.api -dir apps/$$app/$$api/;

api:
	@echo "输入app与api名称,以空格隔开"
	@read app api; \
	zeroctl api new apps/$$app/$$api;

rpc:
	@echo "输入app与api名称,以空格隔开"
	@read app rpc; \
    zeroctl api new apps/$$app/$$rpc;

model-c:
	@echo  "输入app、db与table名称,以空格隔开"
	@read app db table; \
	zeroctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/$$db" -table=$$table  -dir="apps/$$app/model" -c

model:
	@echo  "输入app、db与table名称,以空格隔开"
	@read app db table; \
	zeroctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/$$db" -table=$$table  -dir="apps/$$app/model"

test:
	@read foo; echo "foo=$$foo"
