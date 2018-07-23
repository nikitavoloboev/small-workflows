JsOsaDAS1.001.00bplist00�Vscript_�var system = Application('System Events');
var proc = system.processes.whose({ frontmost: {'=': true } }).name()[0];

var fileName = "testing";

if (proc == "Finder") {
	fileName = finderFile();
} else if (proc == "Path Finder") {
	fileName = pathFinderFile();
} else if (proc == "Commander One") {
	fileName = commanderOneFile();
} else if (proc == "fman") {
	fileName = fmanFile();
}

function finderFile() {
	var app = Application("Finder");
	app.includeStandardAdditions = true;
	app.activate();
	var filePath = app.selection();
	return(filePath);
}

function pathFinderFile() {
	var app = Application("Path Finder");
	app.includeStandardAdditions = true;
	app.activate();
	var filePath = app.selection()[0].posixPath;
	return(filePath);
}

function commanderOneFile() {
	var app = Application("Commander One");
	app.includeStandardAdditions = true;
	app.activate();
	system.keystroke('c', { using: ['command down','control down'] })
	delay(0.5);
	return(app.theClipboard());
}

function fmanFile() {
	var app = Application.currentApplication();
	app.includeStandardAdditions = true;
	system.keyCode(103);
	delay(0.5);
	return(app.theClipboard());
}

pathFinderFile()                              � jscr  ��ޭ