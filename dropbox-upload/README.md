# Dropbox Upload - [Download](https://github.com/nikitavoloboev/small-workflows/blob/master/birthday/Dropbox%20upload.alfredworkflow?raw=true)

> Upload and share files with Dropbox

To use this workflow, you need to first install [Dropbox Uploader script](https://github.com/andreafabrizi/Dropbox-Uploader/blob/master/dropbox_uploader.sh).

I put this script inside `~/Dropbox` named `upload.sh`.

I then designated another folder `~/Dropbox/Shares` as a place where I put all the files I want to share with Dropbox. This workflow **assumes this structure** and **won't work** if you don't have the [above script](https://github.com/andreafabrizi/Dropbox-Uploader/blob/master/dropbox_uploader.sh) inside `Scripts` dir with name `upload.sh` and if you don't have the `Shares` folder too. Of course you can change all of this and find a structure that works for you if you wish. I just share what I am using myself.

Inside the workflow there are 2 things. First, there is a file action that will take the given file, then put it inside `~/Dropbox/Shares`, then share the file and give you a downloadable link of the file in clipboard.

The second is a script filter that will filter through the contents of `~/Dropbox/Shares` and on enter will give you a downloadable link to the file. You can also delete the file from the directory and Dropbox with a modifier key press.

## Contributing

See [contribution guidelines](../contributing.md).
