package main

import (
	"testing"
)

func TestSingleNodeWithName(t *testing.T) {
	rdfs := `<peter> <name> "Peter" .`
	runTestCaseFromString(t, "", rdfs)
}

func TestSingleNodeWithNameAndAge(t *testing.T) {
	rdfs := `
	<peter> <name> "Peter" .

	    <peter> <age> "28"^^<xs:int> .` // Also test blank lines/weird spacing while we're here.
	runTestCaseFromString(t, "", rdfs)
}

func TestUpdatedValue(t *testing.T) {
	rdfs := `
	<peter> <name> "NotPeter" .
	<peter> <name> "Peter" .`
	runTestCaseFromString(t, "", rdfs)
}

func TestAppleIsAFruit(t *testing.T) {
	rdfs := `<apple> <is> <fruit> .`
	runTestCaseFromString(t, "", rdfs)
}

func TestTwoFruits(t *testing.T) {
	rdfs := `
	<apple> <is> <fruit> .
	<banana> <is> <fruit> .`
	runTestCaseFromString(t, "", rdfs)
}

func TestTwoFruitsWithNames(t *testing.T) {
	rdfs := `
	<apple> <is> <fruit> .
	<banana> <is> <fruit> .
	<apple> <name> "MrApple" .
	<banana> <name> "MrBanana" .`
	runTestCaseFromString(t, "", rdfs)
}

func TestBadSelfGeneratedSchema(t *testing.T) {
	rdfs := `
	<abc> <pred> "hello"^^<xs:string> .
	<def> <pred> "42"^^<xs:int> .`
	runTestCaseFromString(t, "", rdfs)
}

func TestBadSelfGeneratedSchemaReverse(t *testing.T) {
	rdfs := `
	<def> <pred> "42"^^<xs:int> .
	<abc> <pred> "hello"^^<xs:string> .`
	runTestCaseFromString(t, "", rdfs)
}

func TestIntConversion(t *testing.T) {
	rdfs := `
	<a> <age> "15"^^<xs:int> .
	<b> <age> "13" .`
	runTestCaseFromString(t, "", rdfs)
}

func TestIntConversionHex(t *testing.T) {
	rdfs := `
	<a> <age> "15"^^<xs:int> .
	<b> <age> "0xff" .`
	runTestCaseFromString(t, "", rdfs)
}

func TestAgeExampleFromDocos(t *testing.T) {
	rdfs := `
	<a> <age> "15"^^<xs:int> .
	<b> <age> "13" .
	<c> <age> "14"^^<xs:string> .
	<d> <age> "14.5"^^<xs:string> .
	<e> <age> "14.5" .`
	runTestCaseFromString(t, "", rdfs)
}

func TestSchemaMismatch(t *testing.T) {
	rdfs := `
	<s_default>  <p_default> "default" .
	<s_string>   <p_default> "str"^^<xs:string> .
	<s_dateTime> <p_default> "2017-08-24T14:31:07.475773659"^^<xs:dateTime> .
	<s_date>     <p_default> "2017-08-24"^^<xs:date> .
	<s_int>      <p_default> "100"^^<xs:int> .
	<s_boolean>  <p_default> "true"^^<xs:boolean> .
	<s_double>   <p_default> "3.14159"^^<xs:double> .

	<s_string>   <p_string> "str"^^<xs:string> .
	<s_default>  <p_string> "default" .
	<s_dateTime> <p_string> "2017-08-24T14:31:07.475773659"^^<xs:dateTime> .
	<s_date>     <p_string> "2017-08-24"^^<xs:date> .
	<s_int>      <p_string> "100"^^<xs:int> .
	<s_boolean>  <p_string> "true"^^<xs:boolean> .
	<s_double>   <p_string> "3.14159"^^<xs:double> .

	<s_dateTime> <p_datetime> "2017-08-24T14:31:07.475773659"^^<xs:dateTime> .
	<s_default>  <p_datetime> "default" .
	<s_string>   <p_datetime> "str"^^<xs:string> .
	<s_date>     <p_datetime> "2017-08-24"^^<xs:date> .
	<s_int>      <p_datetime> "100"^^<xs:int> .
	<s_boolean>  <p_datetime> "true"^^<xs:boolean> .
	<s_double>   <p_datetime> "3.14159"^^<xs:double> .

	<s_date>     <p_date> "2017-08-24"^^<xs:date> .
	<s_dateTime> <p_date> "2017-08-24T14:31:07.475773659"^^<xs:dateTime> .
	<s_default>  <p_date> "default" .
	<s_string>   <p_date> "str"^^<xs:string> .
	<s_int>      <p_date> "100"^^<xs:int> .
	<s_boolean>  <p_date> "true"^^<xs:boolean> .
	<s_double>   <p_date> "3.14159"^^<xs:double> .

	<s_int>      <p_int> "100"^^<xs:int> .
	<s_date>     <p_int> "2017-08-24"^^<xs:date> .
	<s_dateTime> <p_int> "2017-08-24T14:31:07.475773659"^^<xs:dateTime> .
	<s_default>  <p_int> "default" .
	<s_string>   <p_int> "str"^^<xs:string> .
	<s_boolean>  <p_int> "true"^^<xs:boolean> .
	<s_double>   <p_int> "3.14159"^^<xs:double> .

	<s_boolean>  <p_boolean> "true"^^<xs:boolean> .
	<s_int>      <p_boolean> "100"^^<xs:int> .
	<s_date>     <p_boolean> "2017-08-24"^^<xs:date> .
	<s_dateTime> <p_boolean> "2017-08-24T14:31:07.475773659"^^<xs:dateTime> .
	<s_default>  <p_boolean> "default" .
	<s_string>   <p_boolean> "str"^^<xs:string> .
	<s_double>   <p_boolean> "3.14159"^^<xs:double> .

	<s_double>   <p_double> "3.14159"^^<xs:double> .
	<s_boolean>  <p_double> "true"^^<xs:boolean> .
	<s_int>      <p_double> "100"^^<xs:int> .
	<s_date>     <p_double> "2017-08-24"^^<xs:date> .
	<s_dateTime> <p_double> "2017-08-24T14:31:07.475773659"^^<xs:dateTime> .
	<s_default>  <p_double> "default" .
	<s_string>   <p_double> "str"^^<xs:string> .
	`
	runTestCaseFromString(t, "", rdfs)
}

func TestUIDThenDefaultScalar(t *testing.T) {
	rdfs := `
	<subject> <predicate> <object> .
	<subject> <predicate> "object" .
	`
	runTestCaseFromString(t, "", rdfs)
}

func TestDefaultScalarThenUID(t *testing.T) {
	rdfs := `
	<subject> <predicate> "object" .
	<subject> <predicate> <object> .
	`
	runTestCaseFromString(t, "", rdfs)
}

func TestUIDThenString(t *testing.T) {
	rdfs := `
	<subject> <predicate> <object> .
	<subject> <predicate> "object"^^<xs:string> .
	`
	runTestCaseFromString(t, "", rdfs)
}

func TestStringThenUID(t *testing.T) {
	rdfs := `
	<subject> <predicate> "object"^^<xs:string> .
	<subject> <predicate> <object> .
	`
	runTestCaseFromString(t, "", rdfs)
}

func TestSchemaWithPredicateAsString(t *testing.T) {
	rdfs := `<peter> <name> "Peter" .`
	sche := `name: string .`
	runTestCaseFromString(t, sche, rdfs)
}

func TestSchemaUID(t *testing.T) {
	sche := `friend: uid .`
	rdfs := `<alice> <friend> <bob> .`
	runTestCaseFromString(t, sche, rdfs)
}

// TODO: Test cases that cause the loader to fail.

// TODO: Causes the loader to fail.
//func TestSchemaWithPredicateAsInt(t *testing.T) {
//	sche := `age: int .`
//	rdfs := `
//	<pawan> <age> "oaeu"^^<xs:string> .
//	`
//	runTestCaseFromString(t, sche, rdfs)
//}

func TestCountIndexSimple(t *testing.T) {
	sche := `friend: uid @count .`
	rdfs := `
	<alice> <friend> <bob> .
	<alice> <friend> <carol> .
	`
	runTestCaseFromString(t, sche, rdfs)
}

func TestCountIndexComplex(t *testing.T) {
	sche := `friend: uid @count .`
	rdfs := `
	<alice> <friend> <bob>   .
	<alice> <friend> <carol> .
	<alice> <friend> <dave>  .

	<bob>   <friend> <carol> .

	<carol> <friend> <bob>   .
	<carol> <friend> <dave>  .

	<erin>  <friend> <bob>   .
	<erin>  <friend> <carol> .

	<frank> <friend> <carol> .
	<frank> <friend> <dave>  .
	<frank> <friend> <erin>  .

	<grace> <friend> <alice> .
	<grace> <friend> <bob>   .
	<grace> <friend> <carol> .
	<grace> <friend> <dave>  .
	<grace> <friend> <erin>  .
	<grace> <friend> <frank> .
	`
	runTestCaseFromString(t, sche, rdfs)
}

func TestCountIndexMultiplePredicates(t *testing.T) {
	sche := `
	a: uid @count .
	b: uid @count .
	c: uid @count .
	b: uid @count .
	`
	rdfs := `
	<a1> <a> <a3> .
	<a2> <a> <a1> .
	<a2> <a> <a3> .

	<b1> <b> <b2> .

	<c1> <c> <c2> .
	<c2> <c> <c1> .

	<d1> <d> <d2> .
	<d1> <d> <d3> .
	<d3> <d> <d1> .
	<d3> <d> <d2> .
	<d3> <d> <d4> .

	<d4> <d> <d1> .
	<d4> <d> <d2> .
	<d4> <d> <d3> .
	`
	runTestCaseFromString(t, sche, rdfs)
}

func TestIndexIntSimple(t *testing.T) {
	sche := `age: int @index(int) .`
	rdfs := `<peter> <age> "28"^^<xs:int> .`
	runTestCaseFromString(t, sche, rdfs)
}

func TestIndexIntMultiple(t *testing.T) {
	sche := `age: int @index(int) .`
	rdfs := `
	<peter> <age> "28"^^<xs:int> .
	<jim>   <age> "100"^^<xs:int> .
	<adam>  <age> "42"^^<xs:int> .
	<jess>  <age> "42"^^<xs:int> .
	`
	runTestCaseFromString(t, sche, rdfs)
}

func TestIndexFloat(t *testing.T) {
	sche := `shoe_size: float @index(float) .`
	rdfs := `
	<peter> <shoe_size> "9.5"<xs:float> .
	<young_peter> <shoe_size> "2" .
	`
	runTestCaseFromString(t, sche, rdfs)
}

func TestIndexBool(t *testing.T) {
	runTestCaseFromString(t,
		"exists: bool @index(bool) .",
		`
		<santa>           <exists> "false" .
		<climate_change>  <exists> "true"  .
		<xenu>            <exists> "false" .
		<chuck_norris>    <exists> "true"  .
		<rumpelstiltskin> <exists> "false" .
		`,
	)
}

func TestIndexGeo(t *testing.T) {
	runTestCaseFromString(t,
		"location: geo @index(geo) .",
		`
		<presidio_visitor_center> <location> "{'type':'Point','coordinates':[-122.4560447,37.8012321]}"^^<geo:geojson> .
		`,
	)
}

func TestIndexStringExact(t *testing.T) {
	runTestCaseFromString(t,
		"name: string @index(exact) .",
		`
		<peter> <name> "Peter" .
		<doppelganger> <name> "Peter" .
		<john> <name> "John" .
		`,
	)
}

func TestIndexStringHash(t *testing.T) {
	runTestCaseFromString(t,
		"name: string @index(hash) .",
		`
		<peter> <name> "Peter" .
		<doppelganger> <name> "Peter" .
		<john> <name> "John" .
		`,
	)
}

func TestIndexStringTerm(t *testing.T) {
	runTestCaseFromString(t,
		"name: string @index(term) .",
		`
		<ps> <name> "Peter Stace" .
		<pj> <name> "Peter Jackson" .
		`,
	)
}

// TODO: I don't really understand what this index does. Something to with
// languages. Revisit this once language support has been implemented.
func TestIndexStringFullText(t *testing.T) {
	runTestCaseFromString(t,
		"name: string @index(fulltext) .",
		`
		<ps> <name> "Peter Stace" .
		<pj> <name> "Peter Jackson" .
		`,
	)
}

func TestIndexStringTrigram(t *testing.T) {
	runTestCaseFromString(t,
		"text: string @index(trigram) .",
		`
		<lorum> <text> "Sed ut perspiciatis, unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam eaque ipsa, quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt, explicabo. Nemo enim ipsam voluptatem, quia voluptas sit, aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos, qui ratione voluptatem sequi nesciunt, neque porro quisquam est, qui dolorem ipsum, quia dolor sit amet consectetur adipisci[ng] velit, sed quia non numquam [do] eius modi tempora inci[di]dunt, ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur? Quis autem vel eum iure reprehenderit, qui in ea voluptate velit esse, quam nihil molestiae consequatur, vel illum, qui dolorem eum fugiat, quo voluptas nulla pariatur?" .
		<ipsum> <text> "At vero eos et accusamus et iusto odio dignissimos ducimus, qui blanditiis praesentium voluptatum deleniti atque corrupti, quos dolores et quas molestias excepturi sint, obcaecati cupiditate non provident, similique sunt in culpa, qui officia deserunt mollitia animi, id est laborum et dolorum fuga. Et harum quidem rerum facilis est et expedita distinctio. Nam libero tempore, cum soluta nobis est eligendi optio, cumque nihil impedit, quo minus id, quod maxime placeat, facere possimus, omnis voluptas assumenda est, omnis dolor repellendus. Temporibus autem quibusdam et aut officiis debitis aut rerum necessitatibus saepe eveniet, ut et voluptates repudiandae sint et molestiae non recusandae. Itaque earum rerum hic tenetur a sapiente delectus, ut aut reiciendis voluptatibus maiores alias consequatur aut perferendis doloribus asperiores repellat" .
		`,
	)
}

const birthdayRDFs = `
	<a> <birthday> "1997-08-24"^^<xs:dateTime> .
	<b> <birthday> "1997-04-17"^^<xs:dateTime> .
	<c> <birthday> "1998-04-17"^^<xs:dateTime> .
	<d> <birthday> "1998-04-04"^^<xs:dateTime> .
`

func TestIndexDateTimeYear(t *testing.T) {
	runTestCaseFromString(t,
		"birthday: dateTime @index(year) .",
		birthdayRDFs,
	)
}

func TestIndexDateTimeMonth(t *testing.T) {
	runTestCaseFromString(t,
		"birthday: dateTime @index(month) .",
		birthdayRDFs,
	)
}

func TestIndexMultiple(t *testing.T) {
	runTestCaseFromString(t,
		"name: string @index(exact, hash, term, fulltext, trigram). ",
		`
		<ps> <name> "Peter Stace" .
		<pj> <name> "Peter Jackson" .
		`,
	)
}

// TODO: Inappropriate use of indexes (e.g. exact on an int field).

// TODO: Reverse edges.

// TODO: Language.

// TODO: Some really big files.

// TODO: XID edges.
