load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.16.4/rules_go-0.16.4.tar.gz",
    sha256 = "62ec3496a00445889a843062de9930c228b770218c735eca89c67949cd967c3f",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

git_repository(
    name = "herumi_mcl",
    remote = "https://github.com/prysmaticlabs/mcl",
    commit = "01f8ce2606cb93b2ec611207f2d4a9237e7758c9",
)
