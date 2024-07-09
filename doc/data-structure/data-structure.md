# Model

In time-trace data model we have 3 main concepts, Sets, Sub-sets ans Elements. The first important thing you should consider when read this document is to don't confuse this with sets in math, they are similar and have same name, but not completely same.

## Element

Elements in time-trace is name of each separated data we are save and stream. each element has 2 field, value and time.

Something like this:
```
[value, time]
```

And this is how they look likes in code:
```go
type Element struct {
	value []byte
	time  time.Time
}
```

## Sub-Set

Sub-sets is a set of Elements in a array. we can have similar elements (similar value) in same sub-set.

Something like this:
```
[someDate, time-1]
[anotherData, time-2]
[someData, time-3]
[firstValue, time-4]
[secondValue, time-5]
[aValue, time-6]
[aValue, time-7]
```

And this is how it look likes in code:
```go
type SubSet []Element
```

## Set

Sets is a map of string to sub-sets, the string is name of each sub-set (each set itself has a name).

Something like this:
```
Set[goldPrice]:
sub-set[AUD] # GOLD/AUD price
sub-set[USD] # GOLD/USD price
```

And this is how it look likes in code:
```go
type Sets   map[string]Set
type Set    map[string]SubSet
```

## Examples

Before starting you can check [types](../../core/database/types.go) to understand the main 3 concepts better if you wish.

Now we are going to show an example of how it works. our example is for a fintech service for store and stream prices.
(other use cases of time-trace can be IoT, multiplayer games and saving services logs, but we are going to explain it with an fintech example).

In this example our service going to show and store prices of gold (and other stuff). so, we can start by making our database and name it `price-service` in config.
Then we make a Set and assign `GOLD` name to it. each sub set of this set is name of one currency, like AUD, EUR, USD and ...
Each element of each sub-set is the live price of GOLD/SUB-SET-NAME-CURRENCY in value, and the default time.

Like:
```
Set: GOLD
Sub-set: USD
Element example: ["1,970.30", 1698228966]
```
