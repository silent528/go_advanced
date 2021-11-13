# KIT 基础库
protoc --proto_path=. \
--proto_path=../third_party \
--go_out=paths=source_relative:. \
--go-http_out=paths=source_relative:. \
$(API_PROTO_FILES)