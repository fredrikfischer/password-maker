# password-maker
Lösenordsgenerator för extremt starka lösenord, som samtidigt är enkla att minnas 

- Bakgrund och motiv

Detta program generar starka lösenord menat att användas för svensktalande användare. Programmet är inspirerat av NIST lösenordsstandard, vilken uppmuntrar lösenord uppbyggda av flera ihopsatta ord. Detta ger komplexta lösenord som samtidigt är enkla för en människa att minnas. Dessa principer har i detta program anamats och vidareutvecklats, med målet att generera komplexa lösenord, som samtidigtär lätta att minnas och återge.

- Processen för att generera lösenord

Standardversionen av detta program avser att:

    Välja ut tre slumpmässigt valda ord ur en svensk ordlista

    Binda samman dessa m.h.a en "bind-charachter" t.ex punkt eller bindesträck.

    Byter ut en av charachters i lösenordet till en "special charachter". Exempel: Ifall lösenordet innan denna process är "banan-fyra-lindas" och special char är "@", kommer lösenordet efter denna process vara "b@n@n-fyr@-lind@s".

    Slumpvis göra lösenordet antingen lowercase, capitalcase eller uppercase.

    Lägga till ett nummer mellan 0-99, antingen i början eller slutet av lösenordet.

    - Hur man kör detta program

Detta program är skrivet i C# körs från terminalen.

    Ladda ned programmet

    Navigera till mappen i termianlen

    Kör kommandot "dotnet run"

    - Resultat

Exempel av 10 st slumpmässigt genererade lösenord:

54.vargpar.gracer.int0rkad

30+garvar3+ljusramp+byggboom

Av1ghet+Husfru+Lavasten+12

NA7E,RUNDBÅGE,7REAXLAD,24

m@t@ffär-helgelse-ljusrigg-58

gynn@nde_f@luröd_koffert_69

VALBORG,VOLTIGÖR,LIB3RIA,35

8,BROSKBIT,RYSA,WOBBL3R

Förfog@,Scones,M@gs@ft,99

TJOGTAL ROND$KÅL HÅRLUGG 17

- Tillägg och uppmuntran

Användare av programmet uppmuntras att hämta hem och göra egna modifikationer av programmet. Detta för att öka variationen och komplexiteten av de lösenord som tas i bruk.
