#Het Thalia Gastenboek voor Constitutieborrels

_Door Jip J. Dekker_

Ter gevolge van traditie is het geboden dat bij iedere constitutieborrel een gastenboek aanwezig is. Dit object kan door de aanwezigen veilig worden gesteld als de mogelijkheid bestaat om het gastenboek in een geheel te verplaatsen naar een locatie buiten de zaal waar de constitutieborrel plaatsvindt, wat een tegenprestatie afdwingt. Ter voorkoming van een dergelijke situatie heeft Thalia een methode ontwikkeld die tot nog toe nog niet heeft gefaald. In plaats van een papieren boek, gebruikt Thalia een USB-stick waarop alle berichten worden bewaard. Deze USB-stick wordt aan het plafond vastgemaakt en met een lange USB-kabel verbonden aan een oude laptop. De aanwezigen kunnen dan via de laptop hun bericht achterlaten via de software in deze repository, maar kunnen het daadwerkelijke gastenboek niet wegnemen. Belangrijk is het dus dat het gastenboek stevig vast zit aan het plafond. Nu volgt er een beschrijving van de aanwijzingen die moeten worden gevolgd om de laptop gebruiksklaar te maken.

### Deployment
1. We starten met een (oude) laptop waar Linux op is geïnstalleerd. (Of een ander OS wat werkt met Google Go en een _stty_ commando heeft)
2. Zorg ervoor dat de laptop automatisch USB-sticks "mount". Anders gaan berichten verloren als de kabel eenmaal wordt verwijderd.
3. Installeer Google Go(1.0+). [Link](http://golang.org). Voor de meeste distributies is Google Go beschikbaar in de standaard packages.
4. Google Go heeft een GOPATH nodig. Als je bash gebruikt kunt je het volgende doen:
```bash
cd
mkdir go
export GOPATH=$HOME/go
```
5. Download de source code: `go get github.com/jjdekker/Gastenboek`
6. Verander de "Super Secret Code" (om het programma te beëindigen). Deze is te vinden op regel 9 in `gastenboek.go`.
7. Compileer de code opnieuw:
```bash
cd $GOPATH/src/github.com/jjdekker/Gastenboek
go install
```
8. Zorg dat je gastenboek overal kan starten door de executables van Go toe te voegen aan PATH: `export PATH=$PATH:$GOPATH/bin`
9. Sluit de USB-stick aan en ga naar de juiste map met `cd`
10. Start het gastenboek programma met: `exec Gastenboek` (Mocht iemand toch een manier vinden om de applicatie te verlaten zorgt `exec` ervoor dat ze niet terugvallen naar de shell)
