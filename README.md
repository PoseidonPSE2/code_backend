# Go Server

Dieser Server ist in Go geschrieben und ermöglicht es, Daten zu einer Datenbank hinzuzufügen und abzurufen.

## Server starten

Um den Server zu starten, navigieren Sie zu dem Verzeichnis, in dem sich die `server.go` Datei befindet, und führen Sie den folgenden Befehl aus:

go run server.go

Der Server läuft nun und hört auf Port 8080.

Funktionen aufrufen
Es gibt drei Hauptfunktionen, die Sie aufrufen können:

Daten hinzufügen
Um Daten zur Datenbank hinzuzufügen, senden Sie eine POST-Anfrage an http://localhost:8080/add mit einem JSON-Body, der die id, ml und waterType enthält. Zum Beispiel:

{
    "id": "123",
    "ml": "500",
    "waterType": "still"
}


Daten manuell hinzufügen
Um Daten manuell zur Datenbank hinzuzufügen, senden Sie eine GET-Anfrage an http://localhost:8080/addManually mit den id, ml und waterType als URL-Parameter. Zum Beispiel:

http://localhost:8080/addManually?id=123&ml=500&waterType=still

Daten abrufen
Um Daten aus der Datenbank abzurufen, senden Sie eine POST-Anfrage an http://localhost:8080/ mit einem JSON-Body, der die id enthält. Zum Beispiel:

{
    "id": "123"
}

Der Server wird die zugehörigen ml und waterType für die angegebene id zurückgeben.

## Server testen

TestAddData: Dieser Test überprüft die addData Funktion. Er sendet eine POST-Anfrage an den Testserver mit einigen JSON-Daten, die hinzugefügt werden sollen. Der Test überprüft dann, ob der Statuscode der Antwort 200 OK ist, was bedeutet, dass die Daten erfolgreich hinzugefügt wurden.

TestHandleRequest: Dieser Test überprüft die handleRequest Funktion. Zuerst sendet er eine POST-Anfrage an die addData Funktion, um einige Daten hinzuzufügen. Dann sendet er eine POST-Anfrage an die handleRequest Funktion, um die zuvor hinzugefügten Daten abzurufen. Der Test überprüft dann, ob der Statuscode der Antwort 200 OK ist, was bedeutet, dass die Daten erfolgreich abgerufen wurden.

TestAddDataManually: Dieser Test überprüft die addDataManually Funktion. Er sendet eine GET-Anfrage an den Testserver mit einigen Parametern in der URL, die hinzugefügt werden sollen. Der Test überprüft dann, ob der Statuscode der Antwort 200 OK ist, was bedeutet, dass die Daten erfolgreich hinzugefügt wurden.