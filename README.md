# 用自己的域名作为 go module path

你是否好奇 `go.uber.org/zap`, `gopkg.in/yaml.v3` 这样的包是怎么来的？为什么不是 `github.com/uber/zap` 或者 `github.com/go-yaml/yaml`？

这些包都是通过自定义的 go module path 来实现的，在 [Finding a repository for a module path](https://go.dev/ref/mod#vcs-find) 中说明了工作原理。
简单来说，在我们 `go get go.uber.org/zap` 时，go 会先去 `https://go.uber.org/zap?go-get=1` 获取 `meta` 标签，然后根据 `meta` 标签中的 `go-import` 信息，去对应的仓库中查找 `go.uber.org/zap` 这个包。

```shell
$ curl -s https://go.uber.org/zap?go-get=1 | grep go-import

<meta name="go-import" content="go.uber.org/fx git https://github.com/uber-go/fx">
```

上面的 `meta` 标签告诉 go，通过 git 从 https://github.com/uber-go/fx 获取 `go.uber.org/fx` 这个包的源码。

## 为什么要用自定义的 go module path

ChatGPT 的回答是：

1. Semantic Import Versioning: Custom module paths allow organizations to use semantic import versioning, which means they can release breaking changes without changing the import path. This enables smooth upgrades for consumers of the module.
2. Improved Readability: With custom module paths, organizations can create a more descriptive and intuitive import path that aligns with their own namespaces and project structures. It can make the codebase more readable and understandable to developers.
3. Consistent Import Paths: Custom module paths allow organizations to maintain consistent import paths across different repositories and projects. This is particularly useful when projects are split across multiple repositories or when an organization has multiple repositories with related modules.
4. Vendor-agnostic: Using a custom module path makes the codebase vendor-agnostic, meaning the code can be easily integrated with other version control systems or package managers. This can be beneficial if an organization decides to move away from GitHub or migrate to alternative package managers.
5. Branding and Identity: Custom module paths provide organizations with the opportunity to promote their brand and establish a unique identity. By using a custom module path, organizations can reinforce their branding and differentiate themselves from other projects or libraries.

## 如何实现自定义的 go module path

1. 拥有一个域名，比如 j178.dev
2. 在 Vercel 上部署该项目
3. 将 go.j178.dev 设置为该项目的 custom domain

## Try it out

```shell
$ go run go.j178.dev/go-custom-module-path@latest

Hello from go.j178.dev/go-custom-module-path
```
