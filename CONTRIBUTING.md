# Contributing to Small Workflows
There are many ways you can contribute. You can:
- Make suggestions and file bugs in [Issues](../../issues/).
- Fix issues and add features to any of the workflows with [Pull Requests](../../pulls/).

## Improving the workflows
If you wish that a workflow did something that it currently doesn't do or perhaps it does something wrong and you wish it was fixed, you can open an [issue](../../issues/new) saying what it is you want fixed or added. Alternatively, you can make the change yourself in the workflow, then export the workflow with something like [transfer.sh](https://transfer.sh) and I can take a look at it and add the changes.

If you are interested in making your own workflows, you can read [this article](https://medium.com/@nikitavoloboev/writing-alfred-workflows-in-go-2a44f62dc432) on how to write Alfred workflows in Go. And I wrote about [how I make and update my workflows](https://wiki.nikitavoloboev.xyz/macOS/apps/alfred/making-workflows.html).

## Developing Go workflows
There are some workflows in here written in [Go](https://golang.org/) which use [AwGo](https://github.com/deanishe/awgo#readme) library for all Alfred related things.

They also use [modd](https://github.com/cortesi/modd#readme) and [Alfred command](https://godoc.org/github.com/jason0x43/go-alfred/alfred) to help with development of them.

To hack on these workflows you can:
1. `cd` to the directory of the Go workflow and run: `alfred link` inside it.
2. Running `modd` will start a process that will automatically build the workflow with `alfred build` on any changes you make to `.go` files.
3. Make changes to the code or modify Alfred objects to do what you want! Open debugger in Alfred or run the workflow with `workflow:log` passed in as argument to see the logs Alfred produces.

![](https://i.imgur.com/FZ91Qkc.png)

## Sending PRs
1. Fork repo
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Create new Pull Request

## Submitting a Pull Request
Please go through [existing issues](../../issues/) and [pull requests](../../pulls/) to check if somebody else is already working on the issue.

#### Thank you for taking the time to contribute! 💜