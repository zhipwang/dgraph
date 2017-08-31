## Performance Analysis

Running on local machine on the 1M dataset:

0c1b5a101a5749ff76dbd055dbc8a97f4500291b 18s -- uses batch writes to badger
16cf55a478e749bfac8b3201ad8e17e29dc4b5d8 18s -- split out line parsing and rdf parsing
5dab12eb788396f5856953747245c56a8bdde204 18s -- split rdf parsing and rdf processing
9e11cefcd72922d2903f6d3ac12da34d743d4180 17s -- parallelise rdf parsing

Running on i3 AWS machine with 21M dataset:

Hash seems to be AA7920211A1FA9C3

                                         P1 7m22s p2 3m5s  Total 10m27s
729a508a9902bc3eb0942ed34fbb94ce764421b1 P1 7m21s P2 3m16s Total 10m37s -- uses batch writes to badger
9e11cefcd72922d2903f6d3ac12da34d743d4180 P1 7m8s  P2 3m21s Total 10m13s -- parallelise rdf parsing

## Batch size experiment (using 90bee146ce68d6536e1071a453096321ab895c83)

100   - Total: 9m56s Phase1: 7m9s Phase2: 2m46s
1000  - Total: 9m48s Phase1: 7m1s Phase2: 2m47s
1000  - Total: 9m48s Phase1: 6m59s Phase2: 2m49s
10000 - Total: 9m52s Phase1: 6m21s Phase2: 3m30s

Conclusion -- not much difference. As long as there is *some* batching it's ok.

## TODO

- Test cases where the loader should fail. E.g. rdfs that don't match schema,
  or indexes against incorrect schema type.

- Facets (don't know how these work)
- XID edges (don't know how these work)
