# ohgit

Git log statistics tool written in Go.

It is also my Go beginner's practice project.

## install

```zsh
go install github.com/QuentinHsu/ohgit
```

## command

- `path`: 指定 Git 仓库的路径，不指定则默认是命令执行路径

  ```zsh
  ohgit --path ./

  ohgit --path ~
  ```

- `stat-day`：Git log 统计天数，从命令执行时间所在日期起倒推 n 天，默认 1 天

  ```zsh
  ohgit --stat-day 7
  ```

- `user`：指定查询的 Git commit 作者，不指定则不按作者过滤查询

  ```zsh
  ohgit --user QuentinHsu
  ```
