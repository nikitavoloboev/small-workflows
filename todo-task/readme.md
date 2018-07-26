# TODO Task - [Download](https://github.com/nikitavoloboev/small-workflows/blob/master/todo-task/TODO%20Task.alfredworkflow?raw=true)
> Write a task and display it in middle of screen.

## Why use this?
I use this workflow to create (or update) my current **active TODO**.

The way this workflow works is that it creates a small `todo` file the Alfred data directory of this workflow.

You can then by calling the external trigger or with `todo` trigger write the task you want to work on and it will be added as active TODO.

![](https://i.imgur.com/sgJ1xEN.png)

You can then run the external trigger `view`, I call mine from [Karabiner](https://wiki.nikitavoloboev.xyz/macOS/apps/karabiner/karabiner.html) and see the task you should be working on in the middle of the screen. Activating the external trigger again will hide showing of it from the screen.

![](https://i.imgur.com/L5YY3TK.png)

This lets me work on tasks one by one in GTD style. No multi tasking.

There is also an action that will take selected text and put it as your active TODO. Or if you use [Keyboard Maestro](https://www.keyboardmaestro.com/main/), you can use [this macro](https://www.dropbox.com/s/fk7hwt26to1exwf/Add%20selected%20text%20as%20todo.kmmacros?dl=1) to smartly grab text from various apps like 2Do tasks, Trello cards, or `TODO:`'s in code as well as any arbitrary text you want to make your active TODO.

You can also point [BitBar](https://github.com/matryer/bitbar) or [TextBar](http://richsomerfield.com/apps/textbar/) to read the content of the file and that will show the `todo` task in your menu bar.

## Contributing
See [contribution guidelines](../CONTRIBUTING.md#readme).
