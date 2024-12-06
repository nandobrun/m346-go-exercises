# Aufgaben

Die folgende Aufgaben können Sie direkt in die angegebene Datei lösen. Beachten
Sie hierzu die `// TODO: `-Kommentare im Code und die folgenden Instruktionen:

## Steckbrief

`go-1-ex-1/main.go`: Es werden verschiedene Variablen ausgegeben, die jedoch noch
nicht deklariert und initialisiert worden sind. Deklarieren Sie die Variablen
mit einem passenden Datentyp und verwenden Sie Ihre persönlichen Angaben um die
Variablen zu initialisieren.

Tipp: Das Sternzeichen können Sie mit der Unicode-Schreibweise `'\uXXXX'`
eingeben, wobei Sie `XXXX` durch die vier letzten Zeichen des Werts in der
[Spalte
"Unicode"](https://en.wikipedia.org/wiki/Astrological_symbols#Signs_of_the_zodiac)
ersetzen können.

## Umrechnungen

`go-1-ex-2/main.go`: Es werden verschiedene imperiale Einheiten (Meilen, Grad
Fahrenheit) in die entsprechenden metrischen Einheiten umgerechnet. Sie müssen
im unteren Teil des Programms zwei Berechnungen in die andere Richtung
ausführen, d.h. vom metrischen ins imperiale System.

Falls Sie die Aufgabe nicht lösen können, notieren Sie sich die Umrechnungen auf
ein Blatt Papier und stellen Sie die Formeln entsprechend um.

## Standardausgabe, Standardfehler

`go-1-ex-3/main.go`: Das Programm würfelt eine zufällige Zahl (zwischen 1 und 6) und
zeigt diese an. Weiter zeigt das Programm an, wann die Zahl gewürfelt worden
ist. Anstelle von `fmt.Println` verwenden Sie die Funktion `fmt.Fprintln` um die
beiden Informationen auszugeben.

Als erstes Argument müssen Sie hierzu `os.Stdout` oder `os.Stderr` verwenden.
Für welche Art der Information (erwürfelte Augen oder Zeitpunkt des Würfelns)
ist welche Ausgabe (`os.Stdout`/`os.Stderr`) besser geeignet?

Rufen Sie anschliessend das Programm so auf, dass die Augenzahl in die Datei
`eyes.txt` geschrieben wird, und der Zeitpunkt in die Datei `dice.log`. Ergänzen
Sie den Kommentar unten am Programm um diesen Aufruf entsprechend zu
dokumentieren.

## Unterlagen

Durchstöbern Sie bei Gelegenheit die folgenden Webseiten:

- [go.dev](https://go.dev/): Die offizielle Webseite von Go
- [go.dev/doc](https://go.dev/doc/): Die offizielle Dokumentation zu Go
- [pkg.go.dev](https://pkg.go.dev/): Suche nach Packages
- [The Go Playground](https://go.dev/play/): Interaktive
  Online-Entwicklungsumgebung für kleine Beispiele
- [Go by Example](https://gobyexample.com/): Beispiele zu verschiedenen
  Bereichen von Go

Geben Sie in der Kommandozeile die folgenden Befehle ein:

- `go doc`
- `go doc fmt`
- `go doc fmt.Println`

Machen Sie sich mit dieser Form der Dokumentation vertraut und betrachten Sie
ein paar weitere Funktionen aus dem `fmt`-Package.

Sie können mit den Pfeiltasten durch die Ausgabe navigieren. Mit `q` können Sie
die Anzeige verlassen.
