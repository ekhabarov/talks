load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")
load("@io_bazel_rules_docker//container:container.bzl", "container_image")
load("@io_bazel_rules_k8s//k8s:object.bzl", "k8s_object")
load("//:helpers.bzl", "namespace")

go_library(
    name = "service-one_lib", // HL
    srcs = ["main.go"],
    importpath = "github.com/ekhabarov/blog-code-snippets/bazel-k8s-envoy/service-one",
    visibility = ["//visibility:private"],
    deps = [
        "//service-one/pb", # internal deps // HL
        "//service-one/server",
        "@org_golang_google_grpc//:go_default_library", # external deps // HL
        "@org_golang_google_grpc//reflection",
    ],
)

go_binary(
    name = "service-one", // HL
    embed = [":service-one_lib"],
    visibility = ["//visibility:public"],
)

go_image(
    # uses distroless images
    name = "image", // HL
    embed = ["service-one_lib"],
    importpath = "github.com/ekhabarov/blog-code-snippets/bazel-k8s-envoy/service-one",
)

container_image(
    name = "custom_image", // HL
    base = "@go_base//image",
    entrypoint = ["/service-one"],
    env = {
        "ENV_VAR1": "value1",
        "ENV_VAR2": "value2",
    },
    files = [
        ":service-one",
    ],
    labels = {
        "version": "1.0.0",
    },
    ports = [
        "5001/tcp",
    ],
)

k8s_object(
    name = "yaml", // HL
    kind = "deployment",
    substitutions = {
        "%{apiname}": "service-one",
        "%{namespace}": namespace(),
    },
    template = "//k8s:api-deployment-yaml",
)
