# Lead-SCM `pb`
[![travis][travis-badge]][travis]
[![Go Report Card][goreport-badge]][goreport]
[![GoDoc][godoc-badge]][godoc]

Lead-SCM is a distributed version control system (DVCS); it aims to replace Git.

> Git doesn’t so much have a leaky abstraction as **no** abstraction. There is
> essentially no distinction between implementation detail and user interface.
> – [Steve Bennett][quote]

## Short-term Goals

- Provide a simplified* Git-like experience for version control.
- Furthermore, reach feature-parity with Git with select exceptions.

## Long-term Goals

- Lead will be self-hosting. No third-party (like GitHub) will be necessary to
  open pull-requests, perform code-reviews, etc. Everything can be performed either
  via HTML (`pb serve localhost:8080`) or on the CLI.
- Lead will not need a full-checkout in 'lazy' mode. Instead blobs and trees will
  be lazy-loaded from any peer providing the SHA1 hash of the object. This is so that
  Lead can adequately host extremely large projects that preclude a full-checkout†.

## Inspiration

- [Git](https://git-scm.com)
- [Mercurial](https://mercurial-scm.org)
- [Fossil](http://fossil-scm.org)

*Git is known to have a bloated, unclear user interface. Such problems have attempted
to be solved in the past by providing new abstractions over Git. Lead aims to replace
Git.

†Companies with monolithic repos and GBs of history (e.g. Microsoft Windows) cannot
realistically be downloaded to a single hard-drive. Even if it could, forcing each
developer to download the complete history would take an incredible amount of time.

[travis]: https://travis-ci.org/Lead-SCM/pb
[travis-badge]: https://api.travis-ci.org/Lead-SCM/pb.svg?branch=master
[goreport]: https://goreportcard.com/report/github.com/Lead-SCM/pb
[goreport-badge]: https://goreportcard.com/badge/github.com/Lead-SCM/pb
[godoc]: https://godoc.org/github.com/Lead-SCM/pb
[godoc-badge]: https://godoc.org/github.com/Lead-SCM/pb?status.svg
[quote]: https://stevebennett.me/2012/02/24/10-things-i-hate-about-git/
