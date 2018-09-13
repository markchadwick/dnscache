http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.14.0/rules_go-0.14.0.tar.gz"],
    sha256 = "5756a4ad75b3703eb68249d50e23f5d64eaf1593e886b9aa931aa6e938c4e301",
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "c0a5739d12c6d05b6c1ad56f2200cb0b57c5a70e03ebd2f7b87ce88cabf09c7b",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.14.0/bazel-gazelle-0.14.0.tar.gz"],
)

load("@io_bazel_rules_go//go:def.bzl",
     "go_rules_dependencies", "go_register_toolchains")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
go_rules_dependencies()
go_register_toolchains()
gazelle_dependencies()

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
