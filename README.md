I made this to demonstrate a bug in golint

Run golint on this repo and you'll see the following:

```
doublecontext.go:6:1: context.Context should be the first parameter of a function
doublecontext.go:11:1: context.Context should be the first parameter of a function
```

This is likely happening because it takes two contexts in one signature

But it shouldn't matter because not only is a context in the first param, all contexts occur ahead of other types.