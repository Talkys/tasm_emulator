# Todo project description

## Little rant about some golang aspects

I don't know why Google decided to make golang such a git repo dependent language. You're expected to host every package as a different repo and install them to build your programs.

At the same time they're all in for static build, static source is not a first class feature of the language.

My biggest problem here is the strange workaround I had to do to divide my program into deveral submodules on go.work.

In python I would import folder.file.function or something like that. In Go... Well... It really does not want you to make projects relative to the project folder instead of the root golang path. This is not ideal as I don't want all my projects sharing a namespace.

A feature to import local packages as import "folder/package" would be great to work otb.