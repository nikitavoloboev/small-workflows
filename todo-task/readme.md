# TODO Task - [Download](https://github.com/nikitavoloboev/small-workflows/blob/master/todo-task/TODO%20Task.alfredworkflow?raw=true)
> Write a task and display it in middle of screen.

## Why use this?
I use this workflow to create (or update) my current **active TODO**.

The way this workflow works is that it creates a small `todo` file the Alfred data directory of this workflow.

You can then by calling the external trigger or with `todo` trigger write the task you want to work on and it will be added as active TODO.

You can then run the external trigger `view`, I call mine from [Karabiner](https://wiki.nikitavoloboev.xyz/macOS/apps/karabiner/karabiner.html) and see the task you should be working on in the middle of the screen.

This lets me work on tasks one by one in GTD style. No multi tasking.

You can also point [BitBar](https://github.com/matryer/bitbar) or [TextBar](http://richsomerfield.com/apps/textbar/) to read the content of the file and that will show the `todo` task in your menu bar.

There is also an action that will take selected text and put it as your active TODO. Like for example copying some `TODO:` in code or taking tasks from some app like MindNode or 2Do.

## Contributing
See [contribution guidelines](../CONTRIBUTING.md#readme).
