#!/bin/bash

set -euo pipefail

# TODO: Traps didn't work how I thought they did... they work per script rather
# than per function. So need to rethink the cleanup situation.

go install github.com/dgraph-io/dgraph/cmd/dgraph
go install github.com/dgraph-io/dgraph/cmd/dgraphloader
go install github.com/dgraph-io/dgraph/cmd/bulkloader
go install github.com/dgraph-io/dgraph/cmd/dgcmp

#red=$(tput setaf 1)
#green=$(tput setaf 2)
yellow=$(tput setaf 3)
magenta=$(tput setaf 5)
cyan=$(tput setaf 6)
default=$(tput sgr0)

function run_test {

	if [[ $# != 2 ]]; then
		echo "incorrect args"
		exit 1
	fi
	schemaFile=$1
	rdfFile=$2

	# Create temp dirs
	dgLoaderDir=$(mktemp -d --suffix="_bulk_loader_system_test")
	blLoaderDir=$(mktemp -d --suffix="_bulk_loader_system_test")
	h1 () { rm -r $dgLoaderDir $blLoaderDir; }
	trap h1 EXIT

	# Start dgraph
	dgraph -memory_mb=1024 -p $dgLoaderDir/p -w $dgLoaderDir/w &
	dgPid=$!
	h2 () { h1; kill $dgPid || true; } ## OK if this fails, we have probably already cleaned up the proc.
	trap h2 EXIT

	# Wait small amount of time for dgraph to start listening for gRPC.
	sleep 0.5

	# Run the dgraph loader
	t=$(date +%s.%N)
	dgraphloader -c 1 -s $schemaFile -r $rdfFile -cd $dgLoaderDir/c 2>&1 | sed "s/.*/$yellow&$default/"
	dgT=$(echo "$(date +%s.%n) - $t" | bc)

	# Stop dgraph. We'll wait for it to finish later.
	kill -s SIGINT $dgPid

	# Run the bulk loader.
	mkdir $blLoaderDir/p
	t=$(date +%s.%N)
	bulkloader -b $blLoaderDir/p -s $schemaFile -r $rdfFile 2>&1 | sed "s/.*/$cyan&$default/"
	blT=$(echo "$(date +%s.%n) - $t" | bc)

	# Wait for dgraph to finish.
	while ps -p $dgPid 1>/dev/null; do
		sleep 0.1 
	done

	# Compare the two badgers.
	dgcmp -a $dgLoaderDir/p -b $blLoaderDir/p 2>&1 | sed "s/.*/$magenta&$default/" || true # TODO: for now, ignore failure.

	# TODO: Timing
	echo "=== TIMING ==="
	echo "DgraphLoader: $dgT"
	echo "BulkLoader:   $blT"
}

function run_test_str {

	[[ $# == 2 ]] || { echo "incorrect args"; exit 1; }
	schema=$1
	rdfs=$2

	tmpDir=$(mktemp -d --suffix="_bulk_loader_system_test")
	trap "{ rm -r $tmpDir; }" EXIT

	echo "$schema" > $tmpDir/sch.schema
	echo "$rdfs"   > $tmpDir/data.rdf

	run_test $tmpDir/sch.schema $tmpDir/data.rdf
}

function run_test_schema_str {

	[[ $# == 2 ]] || { echo "incorrect args"; exit 1; }
	schema=$1
	rdfFile=$2

	tmpDir=$(mktemp -d --suffix="_bulk_loader_system_test")
	trap "{ rm -r $tmpDir; }" EXIT

	echo "$schema" > $tmpDir/sch.schema
	cat $tmpDir/sch.schema

	run_test $tmpDir/sch.schema $rdfFile
}

run_test_schema_str '
director.film:        uid @reverse @count .
genre:                uid @reverse .
initial_release_date: dateTime @index(year) .
name:                 string @index(term) .
starring:             uid @count .
' /home/petsta/1million.rdf.gz

exit 0 # Disable remaining tests for now while iterating on performance

# Reproduces a bug:
run_test_str '
	name: string @index(term) .
' '
	<foo> <name> "1" .
	<bar> <name> "11" .
	<17216961135462248174> <name> "1" .
'

run_test_str '' '
	<peter> <name> "Peter" .
'

run_test_str '' '
	<peter> <name> "Peter" .
	<peter> <age> "28"^^<xs:int> .
'
run_test_str '' '
	<peter> <name> "NotPeter" .
	<peter> <name> "Peter" .
'
run_test_str '' '
	<apple> <is> <fruit> .
'
run_test_str '' '
	<apple> <is> <fruit> .
	<banana> <is> <fruit> .
'
run_test_str '' '
	<apple> <is> <fruit> .
	<banana> <is> <fruit> .
	<apple> <name> "MrApple" .
	<banana> <name> "MrBanana" .
'
run_test_str '' '
	<abc> <pred> "hello"^^<xs:string> .
	<def> <pred> "42"^^<xs:int> .
'
run_test_str '' '
	<def> <pred> "42"^^<xs:int> .
	<abc> <pred> "hello"^^<xs:string> .
'
run_test_str '' '
	<a> <age> "15"^^<xs:int> .
	<b> <age> "13" .
'
run_test_str '' '
	<a> <age> "15"^^<xs:int> .
	<b> <age> "0xff" .
'
run_test_str '' '
	<a> <age> "15"^^<xs:int> .
	<b> <age> "13" .
	<c> <age> "14"^^<xs:string> .
	<d> <age> "14.5"^^<xs:string> .
	<e> <age> "14.5" .
'
run_test_str '' '
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
'
run_test_str '' '
	<subject> <predicate> <object> .
	<subject> <predicate> "object" .
'
run_test_str '' '
	<subject> <predicate> "object" .
	<subject> <predicate> <object> .
'
run_test_str '' '
	<subject> <predicate> <object> .
	<subject> <predicate> "object"^^<xs:string> .
'
run_test_str '' '
	<subject> <predicate> "object"^^<xs:string> .
	<subject> <predicate> <object> .
'
run_test_str '
	name: string .
' '
	<peter> <name> "Peter" .
'
run_test_str '
	friend: uid .
' '
	<alice> <friend> <bob> .
'
run_test_str '
	friend: uid @count .
' '
	<alice> <friend> <bob> .
	<alice> <friend> <carol> .
'
run_test_str '
	friend: uid @count .
' '
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
'
run_test_str '
	a: uid @count .
	b: uid @count .
	c: uid @count .
	b: uid @count .
' '
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
'
run_test_str '
	age: int @index(int) .
' '
	<peter> <age> "28"^^<xs:int> .
'
run_test_str '
	age: int @index(int) .
' '
	<peter> <age> "28"^^<xs:int> .
	<jim>   <age> "100"^^<xs:int> .
	<adam>  <age> "42"^^<xs:int> .
	<jess>  <age> "42"^^<xs:int> .
'
run_test_str '
	shoe_size: float @index(float) .
' '
	<peter> <shoe_size> "9.5"<xs:float> .
	<young_peter> <shoe_size> "2" .
'
run_test_str '
	exists: bool @index(bool) .
' '
	<santa>           <exists> "false" .
	<climate_change>  <exists> "true"  .
	<xenu>            <exists> "false" .
	<chuck_norris>    <exists> "true"  .
	<rumpelstiltskin> <exists> "false" .
'
#run_test_str '
#	location: geo @index(geo) .
#' '
#	<presidio_visitor_center> <location> "{'type':'Point','coordinates':[-122.4560447,37.8012321]}"^^<geo:geojson> .
#'
run_test_str '
	name: string @index(exact) .
' '
	<peter> <name> "Peter" .
	<doppelganger> <name> "Peter" .
	<john> <name> "John" .
'
run_test_str '
	name: string @index(hash) .
' '
	<peter> <name> "Peter" .
	<doppelganger> <name> "Peter" .
	<john> <name> "John" .
'
run_test_str '
	name: string @index(term) .
' '
	<ps> <name> "Peter Stace" .
	<pj> <name> "Peter Jackson" .
'
run_test_str '
	name: string @index(fulltext) .
' '
	<ps> <name> "Peter Stace" .
	<pj> <name> "Peter Jackson" .
'
run_test_str '
	text: string @index(trigram) .
' '
	<lorum> <text> "Sed ut perspiciatis, unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam eaque ipsa, quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt, explicabo. Nemo enim ipsam voluptatem, quia voluptas sit, aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos, qui ratione voluptatem sequi nesciunt, neque porro quisquam est, qui dolorem ipsum, quia dolor sit amet consectetur adipisci[ng] velit, sed quia non numquam [do] eius modi tempora inci[di]dunt, ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur? Quis autem vel eum iure reprehenderit, qui in ea voluptate velit esse, quam nihil molestiae consequatur, vel illum, qui dolorem eum fugiat, quo voluptas nulla pariatur?" .
	<ipsum> <text> "At vero eos et accusamus et iusto odio dignissimos ducimus, qui blanditiis praesentium voluptatum deleniti atque corrupti, quos dolores et quas molestias excepturi sint, obcaecati cupiditate non provident, similique sunt in culpa, qui officia deserunt mollitia animi, id est laborum et dolorum fuga. Et harum quidem rerum facilis est et expedita distinctio. Nam libero tempore, cum soluta nobis est eligendi optio, cumque nihil impedit, quo minus id, quod maxime placeat, facere possimus, omnis voluptas assumenda est, omnis dolor repellendus. Temporibus autem quibusdam et aut officiis debitis aut rerum necessitatibus saepe eveniet, ut et voluptates repudiandae sint et molestiae non recusandae. Itaque earum rerum hic tenetur a sapiente delectus, ut aut reiciendis voluptatibus maiores alias consequatur aut perferendis doloribus asperiores repellat" .
'
run_test_str '
	birthday: dateTime @index(year) .
' '
	<a> <birthday> "1997-08-24"^^<xs:dateTime> .
	<b> <birthday> "1997-04-17"^^<xs:dateTime> .
	<c> <birthday> "1998-04-17"^^<xs:dateTime> .
	<d> <birthday> "1998-04-04"^^<xs:dateTime> .
'
run_test_str '
	birthday: dateTime @index(month) .
' '
	<a> <birthday> "1997-08-24"^^<xs:dateTime> .
	<b> <birthday> "1997-04-17"^^<xs:dateTime> .
	<c> <birthday> "1998-04-17"^^<xs:dateTime> .
	<d> <birthday> "1998-04-04"^^<xs:dateTime> .
'
run_test_str '
	name: string @index(exact, hash, term, fulltext, trigram). 
' '
	<ps> <name> "Peter Stace" .
	<pj> <name> "Peter Jackson" .
'
run_test_str '
	friend: uid @reverse .
' '
	<a> <friend> <b> .
'
run_test_str '
	friend: uid @reverse .
' '
	<alice> <friend> <bob> .
	<alice> <friend> <carol> .
'
run_test_str '
	friend: uid @reverse @count .
' '
	<a> <friend> <b> .
'
run_test_str '
	friend: uid @reverse @count .
' '
	<alice> <friend> <bob> .
	<alice> <friend> <carol> .
'
run_test_str '' '
	<peter> <name> "Peter"@en .
'
run_test_str '' '
	<peter> <name> "Peter"     .
	<peter> <name> "Peter"@en  .
	<peter> <name> "Peder"@kw  .
	<peter> <name> "Pieter"@af .
	<peter> <name> "Pietru"@mt .
	<peter> <name> "Peddyr"@gv .
'
# "Ped" has peter1 and peter2, "tyr" has only peter1, "dyr" has only peter2.
run_test_str '
	name: string @index(trigram) .
' '
	<peter1> <name> "Pieter"@af .
	<peter1> <name> "Peder"@kw  .

	<peter2> <name> "Peddyr"@gv .
	<peter2> <name> "Pietru"@mt .
'

function fanout_rdfs {
	[[ $# == 1 ]] || { echo "incorrect args"; exit 1; }
	fanout=$1
	for (( i=1; i<$fanout; i++ )); do
		echo "<s> <p> <o_$i> ."
	done
}

run_test_str '' "$(fanout_rdfs 257)"
run_test_str '' "$(fanout_rdfs 9997)"
run_test_str '' "$(fanout_rdfs 9998)"
run_test_str '' "$(fanout_rdfs 9999)"
run_test_str '' "$(fanout_rdfs 10000)"
run_test_str '' "$(fanout_rdfs 10001)"
run_test_str '' "$(fanout_rdfs 19999)"
run_test_str '' "$(fanout_rdfs 20000)"
run_test_str '' "$(fanout_rdfs 30001)"
