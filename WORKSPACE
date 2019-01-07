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

git_repository(
    name = "herumi_mcl",
    remote = "https://github.com/prysmaticlabs/mcl",
    commit ="9355e1d3cef12321fb1002f5707a3655bed7c361",
)

git_repository(
    name = "bazelify_gmp",
    remote = "https://github.com/robin-thomas/bazelify-gmp",
<<<<<<< HEAD
    commit = "3819642a7f67cb6cd001edb862079238eb6d0a18",
)

http_archive(
    name = "gmp_6_1_2",
    build_file = "gmp.6.1.2.BUILD",
    sha256 = "87b565e89a9a684fe4ebeeddb8399dce2599f9c9049854ca8c0dfbdea0e21912",
    strip_prefix = "gmp-6.1.2",
    url = "https://gmplib.org/download/gmp/gmp-6.1.2.tar.xz",
=======
    commit = "d5c10aea1c593f17fe118dbb25623ec6c372cdd3",
>>>>>>> ca1878c... use the bew bazelify_gmp
)

git_repository(
    name = "boringssl",
    commit = "fafc4482e85c09e7af5f71b2eb287b73ccd1020a",
    remote = "https://github.com/google/boringssl",
)
