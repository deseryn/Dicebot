# Dicebot

## How to use
* run _go run main.go_, GET localhost:8080/heroes

## TODO

* Julian:
** Dice 

* Maximilian:
** Wiki

* Chriss:
** DB stuff (SQLite, Reader/Writer)


### API Routes
* GET /dice?dice=3&eyes=20 
* GET /dice/talents/{talent}/{hero}
* GET /dice/attributes/{attribute}/{hero}
* GET /:hero/{field}
* GET /:hero/attributes/{attribute}
* GET /:hero/talents/{talent}
* GET /:hero/info <- basic info like race, class, gender etc.
* GET /info/npc/{npc}
* GET /info/wiki/{something} <- grab from ulisses-regelwiki?
* PUT /info/:npc?something={something}
* PUT /:hero/hp?difference={difference}
* PUT /:hero/money?amount={amount} <- Params for DK, ST, KR?
* PUT /:hero/talents?talent={talent}&value={value}
* PUT /:hero/attributes?attribute={attribute}&value={value}

### Ideas
* maybe: roll dice talents - channel -> compare with hero's talents -> channel response -> return dice values and success with quality level or failure
* go script to initially convert hero JSON and store in db


### Database 
* Heroes: name string, race string, class string, gender string, weight int, height int, hp int, money float, talents

* Attributes: name string, description string

* Talents: name string, attr1 string, attr2 string, attr3 string, description string

* HeroAttributes: hero string, attribute string, value int

* HeroTalents: hero string, talent string, value int



