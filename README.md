# Model2MQL

Convert 
```
{
    "id_gte": 4,
    "name_contains: "Bart"
}
```

to
`id >= 4 and name % "Bart"`

Only meant to be used with chained AND statements. No grouping with parentheses or OR statements.


## Example
```
type Search struct {
    IdGte int `mql:"id"`
    NameContains int `mql:"name"`
}

c := NewConstructor(Search{})
response, _ := c.Convert(
    Search{
        IdGte: 4,
        NameContains: Bart
    }
)

w, err := mql.Parse(response,User{}, mql.WithPgPlaceholders())
if err != nil {
  return nil, err
}
q := fmt.Sprintf("select * from users where %s", w.Condition)
rows, err := db.Query(q, w.Args...)

```
