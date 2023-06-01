load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//container:container.bzl", "container_image", "container_push")

go_library(
    name = "service-one_lib", // HL
    srcs = ["main.go"],
    importpath = "github.com/ekhabarov/bazel-k8s-envoy/service-one",
    deps = [
        "//service-one/pb", # internal dependency // HL
        "//service-one/server",
        "@org_golang_google_grpc//:go_default_library", # external dependency // HL
    ],
)

go_binary(
    name = "service-one", // HL
    embed = [":service-one_lib"],
    visibility = ["//visibility:public"],
)

container_image(
    name = "image", // HL
    base = "@go_base//image",
    entrypoint = ["/service-one"],
    files = [
        ":service-one",
    ],
    ports = [
        "5000/tcp",
    ],
)

genrule(
    name = "yaml", // HL
    srcs = [
        "//k8s:defaults.yaml",
        "//k8s/base:service.yaml",
        ":values.yaml",
    ],
    outs = ["manifest.yaml"],
    cmd = """
      echo $(location @com_github_vmware_tanzu_carvel_ytt//cmd/ytt) \
         `ls $(SRCS) | sed 's/^/-f /g'` \
          > $@
    """,
    executable = True,
    tools = [
        "@com_github_vmware_tanzu_carvel_ytt//cmd/ytt",
    ],
)

container_push(
    name = "push_image",
    format = "Docker",
    image = ":image",
    registry = "ghcr.io",
    repository = "ekhabarov/service-one",
    tag = "$(IMAGE_TAG)",
)
