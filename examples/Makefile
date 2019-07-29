all:
	cd ./httpCallBackServer && go build
	cd ./register_device && go build
	cd ./subscriptions && go build -o subscriptions subscriptions.go
	cd ./subscriptions && go build -o subscriptions_batch_query subscriptions_batch_query.go
	cd ./subscriptions && go build -o subscriptions_batch_delete subscriptions_batch_delete.go

clean:
	cd ./httpCallBackServer && rm -rf ./httpCallBackServer
	cd ./register_device &&  rm -rf ./register_device
	cd ./subscriptions && rm -rf ./subscriptions && rm -rf ./subscriptions_batch_query && rm -rf ./subscriptions_batch_delete

