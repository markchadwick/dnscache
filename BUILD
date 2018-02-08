load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test", "go_prefix")

go_prefix("github.com/markchadwick")

go_library(
    name = "go_default_library",
    srcs = [
        "cache.go",
        "roundtripper.go",
    ],
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
    library = ":go_default_library",
    deps = [
        "@com_github_markchadwick_spec//:go_default_library",
    ],
)
