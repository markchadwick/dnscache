load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "dnscache",
    srcs = [
        "cache.go",
        "roundtripper.go",
    ],
    importpath = "github.com/markchadwick/dnscache",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_rcrowley_go_metrics//:go_default_library",
    ],
)

go_test(
    name = "test",
    size = "small",
    srcs = [
        "cache_test.go",
        "dnscache_test.go",
        "roundtripper_test.go",
    ],
    embed = [":dnscache"],
    deps = [
        "@com_github_markchadwick_spec//:go_default_library",
    ],
)
