# ddiff

Identify quickly dependencies difference between two versions of a Gentoo package based on the reverse dependencies QA [report](https://qa-reports.gentoo.org/output/genrdeps/rdeps.tar.xz). It can give an indicator about the package upgrade complexity.

## Usage

```shell
$ ddiff -current net-fs/samba-4.11.13-r1 -target net-fs/samba-4.12.9-r1
--- Current
+++ Target
@@ -13,0 +14 @@
+dev-libs/icu@@ -18,0 +20 @@
+dev-perl/Parse-Yapp@@ -47,0 +50 @@
+sys-libs/liburing@@ -72,0 +76 @@
+dev-libs/icu@@ -102,0 +107 @@
+sys-libs/liburing
```

## Build and install

```shell
make
```

and move `./ddiff` in your `$PATH`.

## Requirements

You'll need to download the QA report. It can be done using:

```shell
make download
```

It could be nice to handle the download directly in the binary but go standard lib does not support yet `xz` decompression.
