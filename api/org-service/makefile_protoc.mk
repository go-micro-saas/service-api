override ABSOLUTE_MAKEFILE := $(abspath $(lastword $(MAKEFILE_LIST)))
override ABSOLUTE_PATH := $(patsubst %/,%,$(dir $(ABSOLUTE_MAKEFILE)))
override REL_PROJECT_PATH := $(subst $(PROJECT_ABS_PATH)/,,$(ABSOLUTE_PATH))

SAAS_ORG_API_PROTO := $(shell find ./$(REL_PROJECT_PATH) -name "*.proto")
#SAAS_ORG_INTERNAL_PROTO := "app/org-service/internal/conf/config.conf.proto"
SAAS_ORG_INTERNAL_PROTO := ""
SAAS_ORG_PROTO_FILES := ""
ifneq ($(SAAS_ORG_INTERNAL_PROTO), "")
	SAAS_ORG_PROTO_FILES=$(SAAS_ORG_API_PROTO) $(SAAS_ORG_INTERNAL_PROTO)
else
	SAAS_ORG_PROTO_FILES=$(SAAS_ORG_API_PROTO)
endif
.PHONY: protoc-org-protobuf
# protoc :-->: generate org service protobuf
protoc-org-protobuf:
	@echo "# generate testdata service protobuf"
	$(call protoc_protobuf,$(SAAS_ORG_PROTO_FILES))
