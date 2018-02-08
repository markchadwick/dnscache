git_repository(
    name = "io_bazel_rules_go",
    commit = "561efc61f3daa04ad16ff6f75908a88d48c01bb5",
    remote = "https://github.com/bazelbuild/rules_go.git",
)

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_rules_dependencies",
    "go_register_toolchains",
    "go_repository",
)

go_rules_dependencies()

go_register_toolchains()

go_repository(
    name = "com_github_rcrowley_go_metrics",
    commit = "ab2277b1c5d15c3cba104e9cbddbdfc622df5ad8",
    importpath = "github.com/rcrowley/go-metrics",
)

# TEST Deps

go_repository(
    name = "com_github_markchadwick_spec",
    commit = "743340b6dc03c8362bd98be28b891eb7cf0a0c13",
    importpath = "github.com/markchadwick/spec",
)

go_repository(
    name = "com_github_markchadwick_assert",
    commit = "0c5f925a0c673ccaf96f1b37566fdba2f8e3a992",
    importpath = "github.com/markchadwick/assert",
)

go_repository(
    name = "com_github_mgutz_ansi",
    commit = "9520e82c474b0a04dd04f8a40959027271bab992",
    importpath = "github.com/mgutz/ansi",
)

go_repository(
    name = "com_github_mattn_go_colorable",
    commit = "7dc3415be66d7cc68bf0182f35c8d31f8d2ad8a7",
    importpath = "github.com/mattn/go-colorable",
)

go_repository(
    name = "com_github_mattn_go_isatty",
    commit = "6ca4dbf54d38eea1a992b3c722a76a5d1c4cb25c",
    importpath = "github.com/mattn/go-isatty",
)
