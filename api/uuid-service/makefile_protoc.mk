override ABSOLUTE_MAKEFILE := $(abspath $(lastword $(MAKEFILE_LIST)))
override ABSOLUTE_PATH := $(patsubst %/,%,$(dir $(ABSOLUTE_MAKEFILE)))
override REL_PROJECT_PATH := $(subst $(PROJECT_ABS_PATH)/,,$(ABSOLUTE_PATH))

# config
SAAS_UUID_API_PROTO := $(shell find ./$(REL_PROJECT_PATH) -name "*.proto")
#SAAS_UUID_INTERNAL_PROTO := "app/uuid-service/internal/conf/config.conf.proto"
SAAS_UUID_INTERNAL_PROTO := ""
SAAS_UUID_PROTO_FILES := ""
ifneq ($(SAAS_UUID_INTERNAL_PROTO), "")
	SAAS_UUID_PROTO_FILES=$(SAAS_UUID_API_PROTO) $(SAAS_UUID_INTERNAL_PROTO)
else
	SAAS_UUID_PROTO_FILES=$(SAAS_UUID_API_PROTO)
endif
.PHONY: protoc-uuid-protobuf
# protoc :-->: generate uuid service protobuf
protoc-uuid-protobuf:
	@echo "# generate uuid service protobuf"
	$(call protoc_protobuf,$(SAAS_UUID_PROTO_FILES))
