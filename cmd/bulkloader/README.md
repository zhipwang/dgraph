## Performance Analysis

Running on local machine on the 21M dataset.

0c1b5a101a5749ff76dbd055dbc8a97f4500291b 18s -- uses batch writes to badger
16cf55a478e749bfac8b3201ad8e17e29dc4b5d8 18s -- split out line parsing and rdf parsing
5dab12eb788396f5856953747245c56a8bdde204 18s -- split rdf parsing and rdf processing

## TODO

- Test cases where the loader should fail. E.g. rdfs that don't match schema,
  or indexes against incorrect schema type.

- Facets (don't know how these work)
- XID edges (don't know how these work)
