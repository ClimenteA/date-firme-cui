# DATE CUI API

Un api simplu in Go pentru a prelua date firme in functie de CUI. 

- `go run main.go` - pentru a rula API-ul;
- Un GET request la adresa url `http://localhost:3240/date-cui/43000099` va da un raspuns similar:
```json
{
    "cui": "43000099",
    "nume": "POPESCU GOE FIZICA AUTORIZATA",
    "nr_reg_com": "F22/1029/2020",
    "adresa": "Municipiul Iasi, Iasi etc etc",
    "value": "RO-IS, jud. Iasi, Iasi",
    "cod": "RO-IS",
    "judet": "IASI",
    "comuna": "",
    "localitate": "IASI"
}
```

**Observatii**: 
- Setul de date oferit de ONRC este corupt si trebuie "curatat". Am renuntat la aceasta idee din lipsa de timp.


Poti vedea date CUI si la pagina web oferita de Ministerul Finantelor:
- https://mfinante.gov.ro/apps/agenticod.html?pagina=domenii


In docker: 
- `docker-compose up -d` url disponibil la:  `https://datecuiapi.localhost/date-cui/{cui}`; 

Fisierul `procesare_csv_onrc_date_firme.ipynb` din aceste repo te vor ajuta sa creezi baza de date cu firme (lipsesc PFI si SA). Poti optimiza scriptul daca timpul iti permite. 