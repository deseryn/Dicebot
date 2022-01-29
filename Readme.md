# Dicebot

## TODO


### API Routes
* GET /dice?dice=3&eyes=20 
* GET /dice/:talent/:hero
* GET /dice/:attribute/:hero
* GET /:hero/hp
* GET /:hero/money
* GET /:hero/:attribute
* GET /:hero/:talent
* GET /:hero/info <- basic info like race, class, gender etc.
* GET /info/npc/:npc
* GET /info/wiki/:something <- grab from ulisses-regelwiki?
* PUT /info/:npc/:something
* PUT /:hero/hp
* PUT /:hero/money/{amount} <- Params for DK, ST, KR?
* PUT /:hero/:talent/:value
* PUT /:hero/:attribute/:value

### Ideas
* maybe: roll dice talents - channel -> compare with hero's talents -> channel response -> return dice values and success with quality level or failure
* go script to initially convert hero JSON and store in db


### Database 
* Heroes: name string, race string, class string, gender string, weight int, height int, hp int, money float, talents

* Attributes: name string, description string

* Talents: name string, attr1 string, attr2 string, attr3 string, description string

* HeroAttributes: hero string, attribute string, value int

* HeroTalents: hero string, talent string, value int



