rpc_go:
	@echo "输入app名称"
	@read app; \
	zeroctl rpc protoc $$app/rpc/$$app.proto --go_out=$$app/rpc --go-grpc_out=$$app/rpc --zrpc_out=$$app/rpc;

api_go:
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
