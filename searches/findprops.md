SYNOPSIS
findprops [OPTION]... [FILE]...

DESCRIPTION
Find and match properties in XML files and output information about them. This command is typically used for selecting a set of properties in jobTracker configurations and reporting on the values pf various properties (e.g. User Name, Job ID, Hive Query String, etc). 

OPTIONS
-p property=<value>
Specifies a property to match, e.g. `user.name'. Optionally, append `=<value>' to specify only a property with a matching value, e.g. `user.name=jsmith'. Multiple -p options may be specified at the command line. Properties are matched in order of precedence, which has implications for performance; see the EXAMPLES, below. 
The -p option is not required and has no default value. 

-s <selection>
Like the -p option, -s matches a specified property and/or value, but this option prints the results of the match. 
The -s option is not required and defaults to the values of any parameters to -p option. 

-d <delimiter>
Specify a field delimiter for selected output. Whitespace characters may be specified with single quotes, e.g. ' ' for space, '\t' for tab, etc. 
The -d options is not required and defaults to comma (`,').

-c
Prepend the output with column headers, e.g. the names of the properties selected by any parameters to -s option. 
The -c option is not required, takes no arguments, and defaults to off. 

-o file
Specify a file to which selected matches (from -s option) will be written. 
The -o option is not required and defaults to STDOUT. 

FILE...
One or more files. These can be listed singly or as expressions compatible with e.g. the ls command. 

EXAMPLES
findprops -p user.name=jsmith -p job.id -s allfiles*.xml
Search all files matching allfiles*.xml in the current working directory. Of those, find all files with the property `user.name' with value `jsmith' and that also contain the property `job.id' (with any or no value). The -s option selects and prints the values of the `user.name' and `job.id' properties in comma-delimited columns to STDOUT. 
Note that properties are matched in order of precedence. In the above example, if there are many files that contain the property `job.id' but only a few that contain the property `user.name' with value `jsmith', then defining the properties in this order will start with the smallest subset of files and complete sooner. 
findprops -p work.dir -s user.name -s job.id -s query.string -d : -c -o /tmp/queries.csv *.xml
Search all *.xml files in the current working directory for those containing the property `work.dir' and select (print) the values of the `user.name', `job.id', and `query.string' properties. Separate the output values with a colon. Prepend the output with the delimited property names, e.g.: 
user.name:job.id:query.string
...
Write the output to the file /tmp/queries.csv.
