grammar GraphQLPM;

document
   : definition
   ;

definition
   : selectionSet | operationDefinition
   ;

operationDefinition
   : operationType NAME selectionSet
   ;

filters
   : '@filter' '(' pairsNested ')'
   ;

pairsNested
   : '(' pairsNested ')' (filterOperation pairsNested)* |  pairs
   ;
   
pairs
   :  pair (filterOperation pair)*
   ;
   
filterOperation
   : '||' | '&&'
   ;
   
funcName
   : 'anyof' | 'allof'
   ;

pair
   : funcName '(' fieldName ',' value ')'
   ;

fieldName
   : STRING # fieldNameValue
   ;
   
operationType
   : 'query'
   ;

selectionSet
   : '{' field ( ','? field )* '}'
   ;

field
   : NAME arguments? filters? selectionSet?
   ;

arguments
   : '(' argument ( ',' argument )* ')'
   ;

argument
   : NAME ':' value
   ;

value
   : STRING # stringValue
   ;

NAME
   : [_A-Za-z] [._0-9A-Za-z]*
   ;

STRING
   : '"' [ .A-Za-z0-9]* '"'
   ;

WS
   : [ \t\n\r]+ -> skip
   ;
