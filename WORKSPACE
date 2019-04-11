load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.17.0/rules_go-0.17.0.tar.gz"],
    sha256 = "492c3ac68ed9dcf527a07e6a1b2dcbf199c6bf8b35517951467ac32e421c06c1",
)

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.16.0/bazel-gazelle-0.16.0.tar.gz"],
    sha256 = "7949fc6cc17b5b191103e97481cf8889217263acf52e00b560683413af204fcb",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

git_repository(
    name = "herumi_mcl",
    remote = "https://github.com/prysmaticlabs/mcl",
    commit = "79b3a33e21072712f00985ed2adf34b3bcf0d74e",
)

git_repository(
    name = "bazelify_gmp",
    remote = "https://github.com/robin-thomas/bazelify-gmp",
    commit = "bb4881b35e6864c90493980d035e1d984cafd093",
)

git_repository(
    name = "boringssl",
    commit = "4a8c05ffe826c61d50fdf13483b35097168faa5c",
    remote = "https://github.com/google/boringssl",
)

git_repository(
    name = "io_bazel_rules_m4",
    remote = "https://github.com/jmillikin/rules_m4",
    commit = "2bf69df77dfb6b3ba6b7fc95c304b0dc279375bc",
)

load("@io_bazel_rules_m4//:m4.bzl", "m4_register_toolchains")

m4_register_toolchains()
