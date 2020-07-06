# TODO Task - [Download](https://github.com/nikitavoloboev/small-workflows/blob/master/todo-task/TODO%20Task.alfredworkflow?raw=true)

> Write a task and display it on TouchBar, in middle of screen or menu bar.

## Why use this?

I use this workflow to create (or update) my current **active TODO**.

The way this workflow works is that it creates a small `todo` file the Alfred data directory of this workflow.

You can then by calling the external trigger or with `todo` trigger write the task you want to work on and it will be added as active TODO.

![](https://i.imgur.com/wxFuODw.png)

In order to make sure I always keep my active TODO in mind, I display it on my TouchBar with [help of BetterTouchTool](https://github.com/nikitavoloboev/my-mac-os/tree/master/btt).

![](https://i.imgur.com/l9ueqja.png)

There is also an action that will take selected text and put it as your active TODO. Or if you use [Keyboard Maestro](https://www.keyboardmaestro.com/main/), you can use [this macro](https://www.dropbox.com/s/mkn483urqme9hs2/Add%20selected%20text%20as%20todo.kmmacros?dl=1) to smartly grab text from various apps like 2Do tasks, Trello cards, or `TODO:`'s in code as well as any arbitrary text you want to make your active TODO.

This workflow lets me work on tasks one by one in GTD style. No multi tasking. This approach works really well with [breaking down of complex tasks into smaller parts](https://wiki.nikitavoloboev.xyz/research/solving-problems.html).

## Contributing

See [contribution guidelines](../contributing.md).
