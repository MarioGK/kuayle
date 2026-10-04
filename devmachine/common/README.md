# npm security backports in development images

`patch-npm-security.cjs` replaces pinned dependencies bundled by npm. Its versions
stay within npm's existing dependency major versions. Run it after installing npm
and the agent providers. Optional directory arguments locate and patch additional
`http-cache-semantics` copies, for example the code-server installation tree.

Upstream `http-cache-semantics@4.3.0` still contains
[CVE-2026-93748](https://github.com/advisories/GHSA-ch52-4w7c-c8xp).
The local backport prevents `max-stale` from reusing responses prohibited by
`no-cache`, `no-store`, shared `private`, shared `proxy-revalidate`, or shared
`Set-Cookie` without an explicit cache opt-in. It returns the normal revalidation
result before either stale-serving path. Ordinary expiration and explicit
`public`/`immutable` opt-ins retain their existing behavior. The guard also
applies to policies restored from serialized cache entries.

The upstream [security report](https://github.com/kornelski/http-cache-semantics/issues/56)
provides the rationale. The package manifest retains the actual upstream version;
the source patch is explicit, and no scanner suppressions are added.

Every patched copy runs `test-http-cache-semantics.cjs` during the image build.
The patch fails closed if its expected source anchor changes. When changing the
pins, review the source and dependency compatibility, build the IDE and all three
agent images, run their version checks, and scan at HIGH/CRITICAL severity. Remove
the local guard only after an upstream fix passes the negative and positive cache
regressions; updating the package version alone is insufficient.
