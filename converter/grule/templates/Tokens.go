package templates

/*
	rule <RuleName> <RuleDescription> [salience <priority>] {
    when
        <boolean expression>
    then
        <assignment or operation expression>
}

*/

/*`rule r1 "r1" salience 4 {
    	when
        	in.InputBool["IN"] == true
	   	then
        	out.OutputString["OUT"] = "r1";
			Complete();
        	}`
*/

const RULE = `rule {{ template "RULENAME" }} {{template "SALIENCE" }}
              when {{ template "WHEN" }}
			  then {{ template "THEN" }}`

const WHEN = ``
const THEN = ``
const RULENAME = ` $ruleIndex = .
					`
const SALIENCE = ``
