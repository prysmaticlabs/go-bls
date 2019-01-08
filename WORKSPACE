load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.16.5/rules_go-0.16.5.tar.gz",
    sha256 = "7be7dc01f1e0afdba6c8eb2b43d2fa01c743be1b9273ab1eaf6c233df078d705",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

local_repository(
    name = "herumi_mcl",
    path = "../mcl",
#    remote = "https://github.com/prysmaticlabs/mcl",
#    commit ="b90f1fbfea78c47012fcbc9243a616e654794421",
)

local_repository(
    name = "bazelify_gmp",
    path = "../bazelify-gmp",
#    remote = "https://github.com/robin-thomas/bazelify-gmp",
#    commit = "519a5b41c27bc5003a94116ed1b7d05ce732a331",
)

git_repository(
    name = "boringssl",
    commit = "fafc4482e85c09e7af5f71b2eb287b73ccd1020a",
    remote = "https://github.com/google/boringssl",
)

git_repository(
    name = "io_bazel_rules_m4",
    remote = "https://github.com/jmillikin/rules_m4",
    commit = "2bf69df77dfb6b3ba6b7fc95c304b0dc279375bc",
)

load("@io_bazel_rules_m4//:m4.bzl", "m4_register_toolchains")

m4_register_toolchains()
