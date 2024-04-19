pb-gen: pb-gen-coremenu pb-gen-menu
pb-gen-menu:
	protoc --go_out=. ./pkg/shared/protobuff/menu.proto
pb-gen-coremenu:
	protoc --go_out=. ./pkg/shared/protobuff/core-menu.proto