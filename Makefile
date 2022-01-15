gen:
	# Auto-generate code
	protoc --twirp_out=. --go_out=. rpc/category-api/service.proto