workspace(name = "com_cisco_golang_golinters")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

go_rules_version = "0.19.1"

git_repository(
    name = "io_bazel_rules_go",
    remote =  "https://aci-github.cisco.com/shimish2/rules_go.git",
    commit = "ac5f421a63039c8cf4274099b42f5755057699ab",
)

gazelle_version = "0.18.1"

http_archive(
    name = "bazel_gazelle",
    sha256 = "be9296bfd64882e3c08e3283c58fcb461fa6dd3c171764fcc4cf322f60615a9b",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/{}/bazel-gazelle-{}.tar.gz".format(gazelle_version, gazelle_version)],
)


load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(
    go_version = "1.12.7",
    nogo = "@//:nogo"
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "cc_mvdan_interfacer",
    importpath = "mvdan.cc/interfacer",
    sum = "h1:WX1yoOaKQfddO/mLzdV4wptyWgoH/6hwLs7QHTixo0I=",
    build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v0.0.0-20180901003855-c20040233aed",
)

go_repository(
    name = "cc_mvdan_lint",
    importpath = "mvdan.cc/lint",
    sum = "h1:DxJ5nJdkhDlLok9K6qO+5290kphDJbHOQO1DFFFTeBo=",
    build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v0.0.0-20170908181259-adc824a0674b",
)

go_repository(
    name = "cc_mvdan_unparam",
    importpath = "mvdan.cc/unparam",
    sum = "h1:B1LAOfRqg2QUyCdzfjf46quTSYUTAK5OCwbh6pljHbM=",
    build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v0.0.0-20190124213536-fbb59629db34",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
    build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v1.1.1",
)

go_repository(
    name = "com_github_fsnotify_fsnotify",
    importpath = "github.com/fsnotify/fsnotify",
    sum = "h1:IXs+QLmnXW2CcXuY+8Mzv/fWEsPGWxqefPtCP5CnV9I=",
    build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v1.4.7",
)

go_repository(
    name = "com_github_gobwas_glob",
    importpath = "github.com/gobwas/glob",
    sum = "h1:A4xDbljILXROh+kObIiy5kIaPYD8e96x1tgBhUI5J+Y=",
    build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v0.2.3",
)


go_repository(
    name = "com_github_golangci_check",
    importpath = "github.com/golangci/check",
    sum = "h1:23T5iq8rbUYlhpt5DB4XJkc6BU31uODLD1o1gKvZmD0=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v0.0.0-20180506172741-cfe4005ccda2",
)

go_repository(
    name = "com_github_golangci_dupl",
    importpath = "github.com/golangci/dupl",
    sum = "h1:w8hkcTqaFpzKqonE9uMCefW1WDie15eSP/4MssdenaM=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v0.0.0-20180902072040-3e9179ac440a",
)

go_repository(
    name = "com_github_golangci_errcheck",
    importpath = "github.com/golangci/errcheck",
    sum = "h1:i2jIkQFb8RG45DuQs+ElyROY848cSJIoIkBM+7XXypA=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v0.0.0-20181003203344-ef45e06d44b6",
)

go_repository(
    name = "com_github_golangci_go_misc",
    importpath = "github.com/golangci/go-misc",
    sum = "h1:9kfjN3AdxcbsZBf8NjltjWihK2QfBBBZuv91cMFfDHw=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v0.0.0-20180628070357-927a3d87b613",
    patches=["//:third_party/com_github_golangci_go_misc.patch"],
    patch_args = ["-p1"],
)

go_repository(
    name = "com_github_golangci_goconst",
    importpath = "github.com/golangci/goconst",
    sum = "h1:pe9JHs3cHHDQgOFXJJdYkK6fLz2PWyYtP4hthoCMvs8=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v0.0.0-20180610141641-041c5f2b40f3",
)

go_repository(
    name = "com_github_golangci_gocyclo",
    importpath = "github.com/golangci/gocyclo",
    sum = "h1:J2XAy40+7yz70uaOiMbNnluTg7gyQhtGqLQncQh+4J8=",
    patches=["//:third_party/com_github_golangci_gocylo.patch"],
    patch_args = ["-p1"],
    version = "v0.0.0-20180528134321-2becd97e67ee",
)

go_repository(
    name = "com_github_golangci_gosec",
    importpath = "github.com/golangci/gosec",
    sum = "h1:qMomh8bv+kDazm1dSLZ9S3zZ2PJZMHL4ilfBjxFOlmI=",
    patches=["//:third_party/com_github_golangci_gosec.patch"],
    patch_args = ["-p1"],
    version = "v0.0.0-20180901114220-66fb7fc33547",
)

go_repository(
    name = "com_github_golangci_ineffassign",
    importpath = "github.com/golangci/ineffassign",
    sum = "h1:XRFao922N8F3EcIXBSNX8Iywk+GI0dxD/8FicMX2D/c=",
    version = "v0.0.0-20180808204949-42439a7714cc",
)

go_repository(
    name = "com_github_golangci_lint_1",
    importpath = "github.com/golangci/lint-1",
    sum = "h1:r7vyX+SN24x6+5AnpnrRn/bdwBb7U+McZqCHOVtXDuk=",
    version = "v0.0.0-20180610141402-ee948d087217",
)

go_repository(
    name = "com_github_golangci_maligned",
    importpath = "github.com/golangci/maligned",
    sum = "h1:kNY3/svz5T29MYHubXix4aDDuE3RWHkPvopM/EDv/MA=",
    version = "v0.0.0-20180506175553-b1d89398deca",
)

go_repository(
    name = "com_github_golangci_misspell",
    importpath = "github.com/golangci/misspell",
    sum = "h1:EL/O5HGrF7Jaq0yNhBLucz9hTuRzj2LdwGBOaENgxIk=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v0.0.0-20180809174111-950f5d19e770",
)

go_repository(
    name = "com_github_golangci_prealloc",
    importpath = "github.com/golangci/prealloc",
    sum = "h1:leSNB7iYzLYSSx3J/s5sVf4Drkc68W2wm4Ixh/mr0us=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v0.0.0-20180630174525-215b22d4de21",
)

go_repository(
    name = "com_github_golangci_unconvert",
    importpath = "github.com/golangci/unconvert",
    sum = "h1:zwtduBRr5SSWhqsYNgcuWO2kFlpdOZbP0+yRjmvPGys=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v0.0.0-20180507085042-28b1c447d1f4",
)

go_repository(
    name = "com_github_gostaticanalysis_analysisutil",
    importpath = "github.com/gostaticanalysis/analysisutil",
    patches=["//:third_party/com_github_gostaticanalysis_analysisutil.patch"],
    patch_args = ["-p1"],
    tag = "v0.0.2",
)

go_repository(
    name = "com_github_hpcloud_tail",
    importpath = "github.com/hpcloud/tail",
    sum = "h1:nfCOvKYfkgYP8hkirhJocXT2+zOD8yUNjXaWfTlyFKI=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v1.0.0",
)

go_repository(
    name = "com_github_kisielk_gotool",
    importpath = "github.com/kisielk/gotool",
    sum = "h1:AV2c/EiW3KqPNT9ZKl07ehoAGi4C5/01Cfbblndcapg=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v1.0.0",
)

go_repository(
    name = "com_github_kr_pretty",
    importpath = "github.com/kr/pretty",
    sum = "h1:L/CwN0zerZDmRFUapSPitk6f+Q3+0za1rQkzVuMiMFI=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v0.1.0",
)

go_repository(
    name = "com_github_kr_pty",
    importpath = "github.com/kr/pty",
    sum = "h1:VkoXIwSboBpnk99O/KFauAEILuNHv5DVFKZMBN/gUgw=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v1.1.1",
)

go_repository(
    name = "com_github_kr_text",
    importpath = "github.com/kr/text",
    sum = "h1:45sCR5RtlFHMR4UwH9sdQ5TC8v0qDQCHnXt+kaKSTVE=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v0.1.0",
)

go_repository(
    name = "com_github_mozilla_tls_observatory",
    importpath = "github.com/mozilla/tls-observatory",
    sum = "h1:Q0XH6Ql1+Z6YbUKyWyI0sD8/9yH0U8x86yA8LuWMJwY=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v0.0.0-20180409132520-8791a200eb40",
)

go_repository(
    name = "com_github_nbutton23_zxcvbn_go",
    importpath = "github.com/nbutton23/zxcvbn-go",
    sum = "h1:Ri1EhipkbhWsffPJ3IPlrb4SkTOPa2PfRXp3jchBczw=",
    patches=["//:third_party/com_github_nbutton23_zxcvbn_go.patch"],
    patch_args = ["-p1"],
    version = "v0.0.0-20171102151520-eafdab6b0663",
)

go_repository(
    name = "com_github_onsi_ginkgo",
    importpath = "github.com/onsi/ginkgo",
    sum = "h1:Ix8l273rp3QzYgXSR+c8d1fTG7UPgYkOSELPhiY/YGw=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v1.6.0",
)

go_repository(
    name = "com_github_onsi_gomega",
    importpath = "github.com/onsi/gomega",
    sum = "h1:3mYCb7aPxS/RU7TI1y4rkEn1oKmPRjNJLNEXgw7MH2I=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v1.4.2",
)

go_repository(
    name = "com_github_openpeedeep_depguard",
    importpath = "github.com/OpenPeeDeeP/depguard",
    sum = "h1:k9QF73nrHT3nPLz3lu6G5s+3Hi8Je36ODr1F5gjAXXM=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v1.0.0",
)

go_repository(
    name = "com_github_pkg_errors",
    importpath = "github.com/pkg/errors",
    sum = "h1:WdK/asTD0HN+q6hsWO3/vpuAkAr+tw6aNJNDFFf0+qw=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v0.8.0",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v1.0.0",
)

go_repository(
    name = "com_github_rogpeppe_go_internal",
    importpath = "github.com/rogpeppe/go-internal",
    sum = "h1:g0fH8RicVgNl+zVZDCDfbdWxAWoAEJyI7I3TZYXFiig=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v1.1.0",
)

go_repository(
    name = "com_github_ryanuber_go_glob",
    importpath = "github.com/ryanuber/go-glob",
    sum = "h1:7YvPJVmEeFHR1Tj9sZEYsmarJEQfMVYpd/Vyy/A8dqE=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v0.0.0-20170128012129-256dc444b735",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    sum = "h1:bSDNvY7ZPG5RlJ8otE/7V6gMiyenm9RtJ7IUVIAoJ1w=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v1.2.2",
)

go_repository(
    name = "com_github_timakin_bodyclose",
    importpath = "github.com/timakin/bodyclose",
    sum = "h1:lI9ufgFfvuqRctP9Ny8lDDLbSWCMxBPletcSqrnyFYM=",
    patches=["//:third_party/com_github_timakin_bodyclose.patch"],
    patch_args = ["-p1"],
    version = "v0.0.0-20190407043127-4a873e97b2bb",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    tag = "v0.3.0",
)

go_repository(
    name = "in_gopkg_check_v1",
    importpath = "gopkg.in/check.v1",
    sum = "h1:qIbj1fsPNlZgppZ+VLlY7N33q108Sa+fhmuc+sWQYwY=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v1.0.0-20180628173108-788fd7840127",
)

go_repository(
    name = "in_gopkg_errgo_v2",
    importpath = "gopkg.in/errgo.v2",
    sum = "h1:0vLT13EuvQ0hNvakwLuFZ/jYrLp5F3kcWHXdRggjCE8=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v2.1.0",
)

go_repository(
    name = "in_gopkg_fsnotify_v1",
    importpath = "gopkg.in/fsnotify.v1",
    sum = "h1:xOHLXZwVvI9hhs+cLKq5+I5onOuwQLhQwiu63xxlHs4=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v1.4.7",
)

go_repository(
    name = "in_gopkg_tomb_v1",
    importpath = "gopkg.in/tomb.v1",
    sum = "h1:uRGJdciOHaEIrze2W8Q3AKkepLTh2hOroT7a+7czfdQ=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v1.0.0-20141024135613-dd632973f1e7",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    sum = "h1:mUhvW9EsL+naU5Q3cakzfE91YhliOondGd6ZrsDBHQE=",
     build_file_proto_mode = "disable",
    build_extra_args = ["-exclude=vendor"],
    version = "v2.2.1",
)



