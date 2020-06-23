# Contributing

Thank you for taking the time to contribute! ♥️ You can:

- Submit [bug reports or feature requests](../../issues/new/choose). Contribute to discussions. Fix [open issues](../../issues).
- Improve docs, the code and more! Any idea is welcome.

## Run project

Most workflows here are small enough that everything is done with Alfred blocks with minimal code. You can also open the contents of workflow. Like this:

![](https://i.imgur.com/1z1sHEq.png)

Sometimes the workflows are written in [Go](https://golang.org/) and use [AwGo](https://github.com/deanishe/awgo) library for all Alfred related things.

You can run those workflows like so. Assumes you have [Alfred command](https://godoc.org/github.com/jason0x43/go-alfred/alfred) installed.

1. Clone repo
2. cd to folder with workflow you want to edit
3. Run `alfred link` (makes symbolic link of [`workflow`](workflow) directory)
4. Running `alfred build` will build workflow and place binary inside [`workflow`](workflow) directory.
5. Make changes to code or modify Alfred objects to do what you want! Open debugger in Alfred or run the workflow with `workflow:log` passed in as argument to see the logs Alfred produces.

![](https://i.imgur.com/FFYOecx.png)
