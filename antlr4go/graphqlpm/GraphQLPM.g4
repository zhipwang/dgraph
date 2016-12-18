grammar GraphQLPM;

document
   : definition+
   ;

definition
   : selectionSet | operationDefinition
   ;

operationDefinition
   : operationType NAME selectionSet
   ;

operationType
   : 'query'
   ;

selectionSet
   : '{' selection ( ','? selection )* '}'
   ;

selection
   : field
   ;

field
   : fieldName arguments? selectionSet?
   ;

fieldName
   : alias | NAME
   ;

alias
   : NAME ':' NAME
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
   : [_A-Za-z] [_0-9A-Za-z]*
   ;


STRING
   : '"' [A-Za-z0-9]* '"'
   ;

WS
   : [ \t\n\r]+ -> skip
   ;
